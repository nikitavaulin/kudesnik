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
	ProductName  core_http_types.Nullable[string]
	Price        core_http_types.Nullable[float64]
	Description  core_http_types.Nullable[string]
	IsVisible    core_http_types.Nullable[bool]
	CategoryCode core_http_types.Nullable[string]
	ProducerID   core_http_types.Nullable[uuid.UUID]
}

type PatchProductResponseDTO ProductDTOResponse

// PatchProduct godoc
// @Summary Обновить товар
// @Description Обновить товар (требуется указать конкретную категорию)
// @Security BearerAuth
// @Tags products
// @Accept json
// @Produce json
// @Param id path string true "ID товара" Format(uuid)
// @Param category_code path string true "код категории товар которой создается (windows, entrance-doors, interior-doors, balconies, others)"
// @Param request body domain.ProductBasePatch true "PatchProduct тело запроса"
// @Success 200 {object} domain.ProductBase "Обновленный товар"
// @Failure 400 {object} core_http_response.ErrorResponse "Bad request"
// @Failure 401 {object} core_http_response.ErrorResponse "Unauthorized"
// @Failure 403 {object} core_http_response.ErrorResponse "Forbidden"
// @Failure 409 {object} core_http_response.ErrorResponse "Conflict"
// @Failure 500 {object} core_http_response.ErrorResponse "Internal server error"
// @Router /products/{category_code}/{id} [patch]
func (h *ProductsHTTPHandler) PatchProduct(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPReponseHandler(log, rw)

	log.Debug("invoke patch product handler")

	productID, err := core_http_request.GetUUIDFromPath(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get productID")
		return
	}

	category, err := core_http_request.GetCategoryCodeFromPath(r)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get category_code")
		return
	}

	patch := domain.GetProductPatchEmptyInstance(string(category))

	if err := core_http_request.DecodeRequest(r, &patch); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode and validate patch request")
		return
	}

	patchedProduct, err := h.productsService.PatchProduct(ctx, productID, patch)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to patch product")
		return
	}

	responseHandler.JSONResponse(patchedProduct, http.StatusOK)
}

func patchDomainFromDTO(productPatchDTO PatchProductRequestDTO) domain.ProductBasePatch {
	return *domain.NewProductPatch(
		productPatchDTO.ProductName.ToDomain(),
		productPatchDTO.Price.ToDomain(),
		productPatchDTO.Description.ToDomain(),
		productPatchDTO.IsVisible.ToDomain(),
		productPatchDTO.CategoryCode.ToDomain(),
		productPatchDTO.ProducerID.ToDomain(),
	)
}
