package domain

import (
	"fmt"

	"github.com/google/uuid"
	core_errors "github.com/nikitavaulin/kudesnik/internal/core/errors"
	core_validation "github.com/nikitavaulin/kudesnik/internal/core/tools/validation"
)

const (
	MinProductNameLength = 3
	MaxProductNameLength = 100
)

type Product interface {
	GetID() uuid.UUID
	GetVersion() int
	GetProductName() string
	GetPrice() float64
	GetDescription() *string
	GetIsVisible() bool
	GetCategoryID() uuid.UUID
	GetProducerID() *uuid.UUID
	GetDetails() any
}

type BaseProduct struct {
	ID          uuid.UUID
	Version     int
	ProductName string
	Price       float64
	Description *string
	IsVisible   bool
	CategoryID  uuid.UUID
	ProducerID  *uuid.UUID
}

func NewProduct(
	id uuid.UUID,
	version int,
	name string,
	price float64,
	description *string,
	isVisible bool,
	categoryID uuid.UUID,
	producerID *uuid.UUID,
) *BaseProduct {
	return &BaseProduct{
		ID:          id,
		Version:     version,
		ProductName: name,
		Price:       price,
		Description: description,
		IsVisible:   isVisible,
		CategoryID:  categoryID,
		ProducerID:  producerID,
	}
}

func NewProductUninitialized(
	name string,
	price float64,
	description *string,
	categoryID uuid.UUID,
	producerID *uuid.UUID,
) *BaseProduct {
	return NewProduct(
		UninitializedID, UninitializedVersion,
		name, price, description,
		false,
		categoryID, producerID,
	)
}

func (p *BaseProduct) Validate() error {
	if p.ProductName == "" {
		return fmt.Errorf("product name cannot be empty: %w", core_errors.ErrInvalidArgument)
	}

	if err := core_validation.ValidateIntInBounds(len(p.ProductName), MinProductNameLength, MaxProductNameLength); err != nil {
		return fmt.Errorf("wrong product name length: %v: %w", err, core_errors.ErrInvalidArgument)
	}

	if p.Price < 0 {
		return fmt.Errorf("product price must be non-negative: %w", core_errors.ErrInvalidArgument)
	}

	return nil
}

func (p BaseProduct) GetID() uuid.UUID          { return p.ID }
func (p BaseProduct) GetVersion() int           { return p.Version }
func (p BaseProduct) GetProductName() string    { return p.ProductName }
func (p BaseProduct) GetPrice() float64         { return p.Price }
func (p BaseProduct) GetDescription() *string   { return p.Description }
func (p BaseProduct) GetIsVisible() bool        { return p.IsVisible }
func (p BaseProduct) GetCategoryID() uuid.UUID  { return p.CategoryID }
func (p BaseProduct) GetProducerID() *uuid.UUID { return p.ProducerID }

func (p BaseProduct) GetDetails() any {
	return nil
}
