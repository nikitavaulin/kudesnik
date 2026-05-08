package products_transport_http

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/nikitavaulin/kudesnik/internal/core/domain"
	core_logger "github.com/nikitavaulin/kudesnik/internal/core/logger"
	core_http_request "github.com/nikitavaulin/kudesnik/internal/core/transport/http/request"
	core_http_response "github.com/nikitavaulin/kudesnik/internal/core/transport/http/response"
	core_http_types "github.com/nikitavaulin/kudesnik/internal/core/transport/http/types"
)

type PatchProductRequestDTO struct {
	ProductName core_http_types.Nullable[string]
	Price       core_http_types.Nullable[float64]
	Description core_http_types.Nullable[string]
	IsVisible   core_http_types.Nullable[bool]
	CategoryID  core_http_types.Nullable[uuid.UUID]
	ProducerID  core_http_types.Nullable[uuid.UUID]
}

type PatchProductResponseDTO ProductDTOResponse

func (h *ProductsHTTPHandler) PatchProduct(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPReponseHandler(log, rw)

	log.Debug("invoke patch product handler")

	categoryID, err := core_http_request.GetUUIDFromPath(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get productID")
		return
	}

	var productPatchDTO PatchProductRequestDTO

	if err := core_http_request.DecodeAndValidateRequest(r, &productPatchDTO); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode and validate patch request")
		return
	}

	patch := patchDomainFromDTO(productPatchDTO)

	patchedProduct, err := h.productsService.PatchProduct(ctx, categoryID, patch)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to patch product")
		return
	}

	productDTO := PatchProductResponseDTO(productDtoFromDomain(patchedProduct))
	responseHandler.JSONResponse(productDTO, http.StatusOK)
}

func patchDomainFromDTO(productPatchDTO PatchProductRequestDTO) domain.ProductPatch {
	return *domain.NewProductPatch(
		productPatchDTO.ProductName.ToDomain(),
		productPatchDTO.Price.ToDomain(),
		productPatchDTO.Description.ToDomain(),
		productPatchDTO.IsVisible.ToDomain(),
		productPatchDTO.CategoryID.ToDomain(),
		productPatchDTO.ProducerID.ToDomain(),
	)
}
