package product_categories_transport_http

import (
	"fmt"
	"net/http"

	"github.com/nikitavaulin/kudesnik/internal/core/domain"
	core_errors "github.com/nikitavaulin/kudesnik/internal/core/errors"
	core_logger "github.com/nikitavaulin/kudesnik/internal/core/logger"
	core_validation "github.com/nikitavaulin/kudesnik/internal/core/tools/validation"
	core_http_request "github.com/nikitavaulin/kudesnik/internal/core/transport/http/request"
	core_http_response "github.com/nikitavaulin/kudesnik/internal/core/transport/http/response"
	core_http_types "github.com/nikitavaulin/kudesnik/internal/core/transport/http/types"
	core_http_utils "github.com/nikitavaulin/kudesnik/internal/core/transport/http/utils"
)

type PatchProductCategoryRequest struct {
	CategoryName      core_http_types.Nullable[string]  `json:"category_name"`
	InstallationPrice core_http_types.Nullable[float64] `json:"installation_price"`
}

func (r *PatchProductCategoryRequest) Validate() error {
	if r.CategoryName.Set {
		if r.CategoryName.Value == nil {
			return fmt.Errorf("CategoryName can't be patched to NULL: %w", core_errors.ErrInvalidArgument)
		}
		categoryNameLength := len([]byte(*r.CategoryName.Value))
		err := core_validation.ValidateIntInBounds(categoryNameLength, domain.MinProductCategoryNameLength, domain.MaxProductCategoryNameLength)
		if err != nil {
			return fmt.Errorf("invalid CategoryName: %v: %w", err, core_errors.ErrInvalidArgument)
		}
	}
	if r.InstallationPrice.Set {
		if *r.InstallationPrice.Value < 0 {
			return fmt.Errorf("InstallationPrice should be non-negative: %w", core_errors.ErrInvalidArgument)
		}
	}
	return nil
}

type PatchProductCategoryResponse ProductCategoryDTOResponse

func (h *ProductCategoryHTTPHandler) PatchProductCategory(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPReponseHandler(log, rw)

	log.Debug("invoke patch product category handler")

	categoryID, err := core_http_utils.GetUUIDFromPath(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get categoryID")
		return
	}

	var categoryPatchDTO PatchProductCategoryRequest

	if err := core_http_request.DecodeAndValidateRequest(r, &categoryPatchDTO); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode and validate patch request")
		return
	}

	patch := patchDomainFromDTO(categoryPatchDTO)

	patchedCategory, err := h.categoriesService.PatchProductCategory(ctx, categoryID, patch)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to patch product category")
		return
	}

	categoryDTO := PatchProductCategoryResponse(dtoFromDomain(patchedCategory))
	responseHandler.JSONResponse(categoryDTO, http.StatusOK)
}

func patchDomainFromDTO(categoryPatchDTO PatchProductCategoryRequest) domain.ProductCategoryPatch {
	return domain.ProductCategoryPatch{
		CategoryName:      categoryPatchDTO.CategoryName.ToDomain(),
		InstallationPrice: categoryPatchDTO.InstallationPrice.ToDomain(),
	}
}
