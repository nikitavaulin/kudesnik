package domain

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	core_errors "github.com/nikitavaulin/kudesnik/internal/core/errors"
	core_validation "github.com/nikitavaulin/kudesnik/internal/core/tools/validation"
)

type CustomerRequestStatus string

const (
	NewRequestStatus        CustomerRequestStatus = "new"
	InProgressRequestStatus CustomerRequestStatus = "in_progress"
	CompletedRequestStatus  CustomerRequestStatus = "completed"
	CancelledRequestStatus  CustomerRequestStatus = "cancelled"
)

func IsCustomerRequestStatus(status CustomerRequestStatus) bool {
	validStatus := map[CustomerRequestStatus]bool{
		NewRequestStatus:        true,
		InProgressRequestStatus: true,
		CompletedRequestStatus:  true,
		CancelledRequestStatus:  true,
	}
	return validStatus[status]
}

type CustomerRequest struct {
	ID                  uuid.UUID
	Version             int
	DesiredDate         *time.Time
	DesiredTime         *time.Time
	ExtraComment        *string
	Status              CustomerRequestStatus
	CustomerPhoneNumber string
	ProductID           *uuid.UUID
	HandledAt           *time.Time
	CreatedAt           time.Time
	HandlerAdminID      *uuid.UUID
}

type CustomerRequestDetailed struct {
	CustomerRequest
	Fullname      *string
	ChosenProduct *ProductBaseDetailed
}

type CustomerRequestForList struct {
	ID                  uuid.UUID             `json:"id"`
	DesiredDate         *time.Time            `json:"desired_date,omitempty"`
	DesiredTime         *time.Time            `json:"desired_time,omitempty"`
	Status              CustomerRequestStatus `json:"status"`
	CustomerPhoneNumber string                `json:"customer_phone_number"`
	CustomerFullname    *string               `json:"customer_fullname,omitempty"`
	CreatedAt           time.Time             `json:"created_at"`
	HandlerAdminID      *uuid.UUID            `json:"handler_admin_id,omitempty"`
	HandlerAdminName    *string               `json:"handler_admin_name,omitempty"`
	HandledAt           *time.Time            `json:"handled_at,omitempty"`
}

func NewCustomerRequest(
	customerPhoneNumber string,
	desiredDate *time.Time,
	desiredTime *time.Time,
	extraComment *string,
	productID *uuid.UUID,
) *CustomerRequest {
	now := time.Now()

	return &CustomerRequest{
		ID:                  uuid.New(),
		Version:             1,
		DesiredDate:         desiredDate,
		DesiredTime:         desiredTime,
		ExtraComment:        extraComment,
		Status:              NewRequestStatus,
		CustomerPhoneNumber: customerPhoneNumber,
		ProductID:           productID,
		CreatedAt:           now,
	}
}

func (c *CustomerRequest) Validate() error {
	if !IsCustomerRequestStatus(c.Status) {
		return fmt.Errorf("unknown status: %w", core_errors.ErrInvalidArgument)
	}

	if err := core_validation.ValidatePhoneNumber(c.CustomerPhoneNumber); err != nil {
		return fmt.Errorf("invalid phone number: %w", err)
	}

	if c.CreatedAt.IsZero() {
		return fmt.Errorf("created_at is required: %w", core_errors.ErrInvalidArgument)
	}

	if err := c.ValidateStatus(); err != nil {
		return fmt.Errorf("invalid status: %v: %w", err, core_errors.ErrInvalidArgument)
	}

	if c.DesiredDate != nil && c.DesiredDate.IsZero() {
		return fmt.Errorf("desired_date cannot be zero if provided")
	}

	if c.DesiredTime != nil && c.DesiredTime.IsZero() {
		return fmt.Errorf("desired_time cannot be zero if provided")
	}

	if c.ProductID != nil && *c.ProductID == uuid.Nil {
		return fmt.Errorf("product_id cannot be nil UUID if provided")
	}

	return nil
}

func (c *CustomerRequest) ValidateStatus() error {
	switch c.Status {
	case CompletedRequestStatus, CancelledRequestStatus:
		// Для completed и cancelled: handled_at и handler_admin_id обязательны,
		// и handled_at должен быть >= created_at
		if c.HandledAt == nil {
			return fmt.Errorf("handled_at is required for status: %s", c.Status)
		}
		if c.HandlerAdminID == nil || *c.HandlerAdminID == uuid.Nil {
			return fmt.Errorf("handler_admin_id is required for status: %s", c.Status)
		}
		if c.HandledAt.Before(c.CreatedAt) {
			return fmt.Errorf("handled_at must be >= created_at for status: %s", c.Status)
		}

	case InProgressRequestStatus:
		// Для in_progress: handled_at обязателен
		if c.HandledAt == nil {
			return fmt.Errorf("handled_at is required for in_progress status")
		}

	case NewRequestStatus:
		// Для new: handled_at и handler_admin_id должны быть NULL
		if c.HandledAt != nil {
			return fmt.Errorf("handled_at must be nil for new status")
		}
		if c.HandlerAdminID != nil {
			return fmt.Errorf("handler_admin_id must be nil for new status")
		}
	}
	return nil
}
