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
	ID                  uuid.UUID             `json:"id"`
	Version             int                   `json:"-"`
	DesiredDate         *time.Time            `json:"desired_date,omitempty"`
	DesiredTime         *time.Time            `json:"desired_time,omitempty"`
	ExtraComment        *string               `json:"extra_comment,omitempty"`
	Status              CustomerRequestStatus `json:"status"`
	CustomerPhoneNumber string                `json:"customer_phone_number"`
	ProductID           *uuid.UUID            `json:"product_id,omitempty"`
	HandledAt           *time.Time            `json:"handled_at,omitempty"`
	CreatedAt           time.Time             `json:"created_at"`
	HandlerAdminID      *uuid.UUID            `json:"handler_admin_id,omitempty"`
}

type CustomerRequestDetailed struct {
	CustomerRequest
	CustomerFullname *string              `json:"customer_fullname,omitempty"`
	HandlerAdminName *string              `json:"handler_admin_id,omitempty"`
	ChosenProduct    *ProductBaseDetailed `json:"chosen_product,omitempty"`
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

type CustomerRequestPatch struct {
	DesiredDate  Nullable[time.Time] `json:"desired_date"`
	DesiredTime  Nullable[time.Time] `json:"desired_time"`
	ExtraComment Nullable[string]    `json:"extra_comment"`
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

// Validate проверяет корректность патча для заявки клиента
func (p *CustomerRequestPatch) Validate() error {
	if p.DesiredDate.Set && p.DesiredDate.Value == nil {
		return fmt.Errorf("desired_date can't be patched to NULL: %w", core_errors.ErrInvalidArgument)
	}

	if p.DesiredTime.Set && p.DesiredTime.Value == nil {
		return fmt.Errorf("desired_time can't be patched to NULL: %w", core_errors.ErrInvalidArgument)
	}

	if p.ExtraComment.Set && p.ExtraComment.Value == nil {
		return fmt.Errorf("extra_comment can't be patched to NULL: %w", core_errors.ErrInvalidArgument)
	}

	if p.DesiredDate.Set && p.DesiredDate.Value != nil {
		if (*p.DesiredDate.Value).IsZero() {
			return fmt.Errorf("desired_date cannot be zero if provided: %w", core_errors.ErrInvalidArgument)
		}
	}

	if p.DesiredTime.Set && p.DesiredTime.Value != nil {
		if (*p.DesiredTime.Value).IsZero() {
			return fmt.Errorf("desired_time cannot be zero if provided: %w", core_errors.ErrInvalidArgument)
		}
	}

	if p.DesiredDate.Set && p.DesiredDate.Value != nil && p.DesiredTime.Set && p.DesiredTime.Value != nil {
		desiredDateTime := time.Date(
			(*p.DesiredDate.Value).Year(),
			(*p.DesiredDate.Value).Month(),
			(*p.DesiredDate.Value).Day(),
			(*p.DesiredTime.Value).Hour(),
			(*p.DesiredTime.Value).Minute(),
			(*p.DesiredTime.Value).Second(),
			(*p.DesiredTime.Value).Nanosecond(),
			(*p.DesiredTime.Value).Location(),
		)

		if desiredDateTime.Before(time.Now()) {
			return fmt.Errorf("desired datetime cannot be in the past: %w", core_errors.ErrInvalidArgument)
		}
	}

	return nil
}

// ApplyPatch применяет патч к заявке клиента
func (c *CustomerRequest) ApplyPatch(patch CustomerRequestPatch) error {
	if err := patch.Validate(); err != nil {
		return fmt.Errorf("invalid customer request patch: %w", err)
	}

	// Создаем копию для проверки
	tmp := *c

	if patch.DesiredDate.Set {
		tmp.DesiredDate = patch.DesiredDate.Value
	}

	if patch.DesiredTime.Set {
		tmp.DesiredTime = patch.DesiredTime.Value
	}

	if patch.ExtraComment.Set {
		tmp.ExtraComment = patch.ExtraComment.Value
	}

	// Валидируем полученный объект
	if err := tmp.Validate(); err != nil {
		return fmt.Errorf("invalid patched customer request: %w", err)
	}

	// Применяем изменения
	*c = tmp

	return nil
}

// NewCustomerRequestPatch создает новый патч для заявки клиента
func NewCustomerRequestPatch(
	desiredDate Nullable[time.Time],
	desiredTime Nullable[time.Time],
	extraComment Nullable[string],
) CustomerRequestPatch {
	return CustomerRequestPatch{
		DesiredDate:  desiredDate,
		DesiredTime:  desiredTime,
		ExtraComment: extraComment,
	}
}
