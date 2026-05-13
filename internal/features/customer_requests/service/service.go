package customer_requests_service

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/nikitavaulin/kudesnik/internal/core/domain"
)

type CustomerRequestsService struct {
	requestsRepo CustomerRequestsRepository
	productsRepo ProductsRepository
}

type CustomerRequestsRepository interface {
	CreateCustomerRequest(ctx context.Context, request domain.CustomerRequest, customer domain.Customer) (uuid.UUID, error)
	DeleteCustomerRequest(ctx context.Context, requestID uuid.UUID) error

	UpdateCustomerRequestStatus(ctx context.Context, reqID, adminID uuid.UUID, status domain.CustomerRequestStatus, handledAt time.Time) error
	PatchCustomerRequest(ctx context.Context, requestID uuid.UUID, request domain.CustomerRequest) error

	GetCustomerRequest(ctx context.Context, id uuid.UUID) (domain.CustomerRequestDetailed, error)
	GetCustomerRequests(ctx context.Context, status *domain.CustomerRequestStatus, phone *string, adminID *uuid.UUID, limit, offset *int) ([]domain.CustomerRequestForList, error)
}

type ProductsRepository interface {
	GetProductDetailed(ctx context.Context, id uuid.UUID) (domain.ProductBaseDetailed, error)
}

func NewCustomerRequestsService(requestsRepo CustomerRequestsRepository, productsRepo ProductsRepository) *CustomerRequestsService {
	return &CustomerRequestsService{
		requestsRepo: requestsRepo,
		productsRepo: productsRepo,
	}
}
