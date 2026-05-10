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
	GetBase() *ProductBase
	GetCategoryName() ProductCategoryName
	Validate() error
}

func GetProductEmptyInstance(categoryName string) Product {
	category := GetCategoryName(categoryName)
	switch category {
	case WindowsCategory:
		return &Window{}
	default:
		return &ProductBase{}
	}
}

type ProductBase struct {
	ID          uuid.UUID  `json:"id"`
	Version     int        `json:"version"`
	ProductName string     `json:"product_name"`
	Price       float64    `json:"price"`
	Description *string    `json:"description"`
	IsVisible   bool       `json:"is_visible"`
	CategoryID  uuid.UUID  `json:"category_id"`
	ProducerID  *uuid.UUID `json:"producer_id,omitempty"`
}

type ProductBaseDetailed struct {
	ProductBase
	ProductDetails
}

type ProductDetails struct {
	CategoryName        string `json:"category_name"`
	ProducerCompanyName string `json:"producer_company_name"`
}

func (p *ProductBase) GetBase() *ProductBase                { return p }
func (p *ProductBase) GetCategoryName() ProductCategoryName { return OthersCategory }

func NewProduct(
	id uuid.UUID,
	version int,
	name string,
	price float64,
	description *string,
	isVisible bool,
	categoryID uuid.UUID,
	producerID *uuid.UUID,
) *ProductBase {
	return &ProductBase{
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
) *ProductBase {
	return NewProduct(
		UninitializedID, UninitializedVersion,
		name, price, description,
		false,
		categoryID, producerID,
	)
}

func (p *ProductBase) Validate() error {
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

type ProductPatch struct {
	ProductName Nullable[string]
	Price       Nullable[float64]
	Description Nullable[string]
	IsVisible   Nullable[bool]
	CategoryID  Nullable[uuid.UUID]
	ProducerID  Nullable[uuid.UUID]
}

func NewProductPatch(
	productName Nullable[string],
	price Nullable[float64],
	description Nullable[string],
	isVisible Nullable[bool],
	categoryID Nullable[uuid.UUID],
	producerID Nullable[uuid.UUID],
) *ProductPatch {
	return &ProductPatch{
		ProductName: productName,
		Price:       price,
		Description: description,
		IsVisible:   isVisible,
		CategoryID:  categoryID,
		ProducerID:  producerID,
	}
}

func (p *ProductPatch) Validate() error {
	if p.ProductName.Set {
		if *p.ProductName.Value == "" {
			return fmt.Errorf("product name cannot be empty: %w", core_errors.ErrInvalidArgument)
		}

		if err := core_validation.ValidateIntInBounds(len(*p.ProductName.Value), MinProductNameLength, MaxProductNameLength); err != nil {
			return fmt.Errorf("wrong product name length: %v: %w", err, core_errors.ErrInvalidArgument)
		}
	}

	if p.Price.Set && *p.Price.Value < 0 {
		return fmt.Errorf("product price must be non-negative: %w", core_errors.ErrInvalidArgument)
	}

	return nil
}

func (p *ProductBase) ApplyPatch(patch ProductPatch) error {
	if err := patch.Validate(); err != nil {
		return fmt.Errorf("invalid product patch: %w", err)
	}

	tmp := *p

	if patch.ProductName.Set {
		tmp.ProductName = *patch.ProductName.Value
	}

	if patch.Description.Set {
		tmp.Description = patch.Description.Value
	}

	if patch.Price.Set {
		tmp.Price = *patch.Price.Value
	}

	if patch.IsVisible.Set {
		tmp.IsVisible = *patch.IsVisible.Value
	}

	if patch.IsVisible.Set {
		tmp.CategoryID = *patch.CategoryID.Value
	}

	if patch.ProducerID.Set {
		tmp.ProducerID = patch.ProducerID.Value
	}

	if err := tmp.Validate(); err != nil {
		return fmt.Errorf("invalid product after patch: %w", err)
	}

	*p = tmp

	return nil
}
