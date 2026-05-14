package customer_requests_transport_http

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/nikitavaulin/kudesnik/internal/core/domain"
	core_http_server "github.com/nikitavaulin/kudesnik/internal/core/transport/http/server"
)

type CustomerRequestsTransportHTTP struct {
	requestsService CustomerRequestsService
}

type CustomerRequestsService interface {
	CreateCustomerRequest(ctx context.Context, request domain.CustomerRequest, customer domain.Customer) (uuid.UUID, error)
	GetCustomerRequest(ctx context.Context, id uuid.UUID) (domain.CustomerRequestDetailed, error)
	GetCustomerRequests(
		ctx context.Context,
		status *domain.CustomerRequestStatus,
		phone *string,
		adminID *uuid.UUID,
		limit, offset *int,
	) ([]domain.CustomerRequestForList, error)

	UpdateCustomerRequestStatus(ctx context.Context, reqID, adminID uuid.UUID, status domain.CustomerRequestStatus) error
	PatchCustomerRequest(ctx context.Context, requestID uuid.UUID, patch domain.CustomerRequestPatch) error
	DeleteCustomerRequest(ctx context.Context, requestID uuid.UUID) error
}

func NewCustomerRequestsTransportHTTP(requestsService CustomerRequestsService) *CustomerRequestsTransportHTTP {
	return &CustomerRequestsTransportHTTP{
		requestsService: requestsService,
	}
}

func (h *CustomerRequestsTransportHTTP) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:  http.MethodPost,
			Path:    "/requests",
			Handler: h.CreateCustomerRequest,
		},
		{
			Method:       http.MethodGet,
			Path:         "/requests/{id}",
			Handler:      h.GetCustomerRequest,
			RequiresAuth: true,
			AllowedRoles: []domain.Role{domain.AdminRole, domain.ManagerRole},
		},
		{
			Method:       http.MethodGet,
			Path:         "/requests",
			Handler:      h.GetCustomerRequests,
			RequiresAuth: true,
			AllowedRoles: []domain.Role{domain.AdminRole, domain.ManagerRole},
		},
		{
			Method:       http.MethodPatch,
			Path:         "/requests/status/{id}",
			Handler:      h.UpdateCustomerRequestStatus,
			RequiresAuth: true,
			AllowedRoles: []domain.Role{domain.AdminRole, domain.ManagerRole},
		},
		{
			Method:       http.MethodPatch,
			Path:         "/requests/{id}",
			Handler:      h.PatchCustomerRequest,
			RequiresAuth: true,
			AllowedRoles: []domain.Role{domain.AdminRole, domain.ManagerRole},
		},
		{
			Method:       http.MethodDelete,
			Path:         "/requests/{id}",
			Handler:      h.DeleteCustomerRequest,
			RequiresAuth: true,
			AllowedRoles: []domain.Role{domain.AdminRole},
		},
	}
}
