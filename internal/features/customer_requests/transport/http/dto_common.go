package customer_requests_transport_http

import (
	"time"

	"github.com/google/uuid"
	"github.com/nikitavaulin/kudesnik/internal/core/domain"
)

type CustomerRequestDTO struct {
	CustomerPhoneNumber string     `json:"customer_phone_number"`
	CustomerFullname    *string    `json:"customer_fullname"`
	DesiredDate         *time.Time `json:"desired_date,omitempty"`
	DesiredTime         *time.Time `json:"desired_time,omitempty"`
	ExtraComment        *string    `json:"extra_comment,omitempty"`
	ProductID           *uuid.UUID `json:"product_id,omitempty"`
}

type CustomerRequestDetailedDTO struct {
	CustomerRequestDTO
	ChosenProduct *domain.ProductBaseDetailed `json:"chosen_product"`
}

type CustomerRequestIDResponseDTO struct {
	CustomerRequestID string `json:"customer_request_id"`
}

func ToCustomerRequestDetailedDTO(detailed domain.CustomerRequestDetailed) CustomerRequestDetailedDTO {
	return CustomerRequestDetailedDTO{
		CustomerRequestDTO: CustomerRequestDTO{
			CustomerPhoneNumber: detailed.CustomerPhoneNumber,
			CustomerFullname:    detailed.Fullname,
			DesiredDate:         detailed.DesiredDate,
			DesiredTime:         detailed.DesiredTime,
			ExtraComment:        detailed.ExtraComment,
			ProductID:           detailed.ProductID,
		},
		ChosenProduct: detailed.ChosenProduct,
	}
}

func ToCustomerRequestDetailedDTOs(detailedList []domain.CustomerRequestDetailed) []CustomerRequestDetailedDTO {
	dtos := make([]CustomerRequestDetailedDTO, len(detailedList))
	for i, detailed := range detailedList {
		dtos[i] = ToCustomerRequestDetailedDTO(detailed)
	}
	return dtos
}
