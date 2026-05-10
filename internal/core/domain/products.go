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

type ProductBase struct {
	ID           uuid.UUID  `json:"id"`
	Version      int        `json:"version"`
	ProductName  string     `json:"product_name"`
	Price        float64    `json:"price"`
	Description  *string    `json:"description"`
	IsVisible    bool       `json:"is_visible"`
	CategoryCode string     `json:"category_code"`
	ProducerID   *uuid.UUID `json:"producer_id,omitempty"`
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
func (p *ProductBase) GetCategoryName() ProductCategoryCode { return OthersCategory }

func NewProduct(
	id uuid.UUID,
	version int,
	name string,
	price float64,
	description *string,
	isVisible bool,
	categoryCode string,
	producerID *uuid.UUID,
) *ProductBase {
	return &ProductBase{
		ID:           id,
		Version:      version,
		ProductName:  name,
		Price:        price,
		Description:  description,
		IsVisible:    isVisible,
		CategoryCode: categoryCode,
		ProducerID:   producerID,
	}
}

func NewProductUninitialized(
	name string,
	price float64,
	description *string,
	categoryCode string,
	producerID *uuid.UUID,
) *ProductBase {
	return NewProduct(
		UninitializedID, UninitializedVersion,
		name, price, description,
		false,
		categoryCode, producerID,
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

	if err := ValidateCategoryCode(p.CategoryCode); err != nil {
		return err
	}

	return nil
}

type ProductBasePatch struct {
	ProductName  Nullable[string]    `json:"product_name"`
	Price        Nullable[float64]   `json:"price"`
	Description  Nullable[string]    `json:"description"`
	IsVisible    Nullable[bool]      `json:"is_visible"`
	CategoryCode Nullable[string]    `json:"category_code"`
	ProducerID   Nullable[uuid.UUID] `json:"producer_id"`
}

func NewProductPatch(
	productName Nullable[string],
	price Nullable[float64],
	description Nullable[string],
	isVisible Nullable[bool],
	categoryCode Nullable[string],
	producerID Nullable[uuid.UUID],
) *ProductBasePatch {
	return &ProductBasePatch{
		ProductName:  productName,
		Price:        price,
		Description:  description,
		IsVisible:    isVisible,
		CategoryCode: categoryCode,
		ProducerID:   producerID,
	}
}

func (p *ProductBasePatch) Validate() error {
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

	if p.CategoryCode.Set {
		if err := ValidateCategoryCode(*p.CategoryCode.Value); err != nil {
			return err
		}
	}

	return nil
}

func (p *ProductBase) ApplyPatch(basePatch *ProductBasePatch) error {
	if err := basePatch.Validate(); err != nil {
		return fmt.Errorf("invalid product patch: %w", err)
	}

	tmp := *p

	if basePatch.ProductName.Set {
		tmp.ProductName = *basePatch.ProductName.Value
	}

	if basePatch.Description.Set {
		tmp.Description = basePatch.Description.Value
	}

	if basePatch.Price.Set {
		tmp.Price = *basePatch.Price.Value
	}

	if basePatch.IsVisible.Set {
		tmp.IsVisible = *basePatch.IsVisible.Value
	}

	if basePatch.IsVisible.Set {
		tmp.CategoryCode = *basePatch.CategoryCode.Value
	}

	if basePatch.ProducerID.Set {
		tmp.ProducerID = basePatch.ProducerID.Value
	}

	if err := tmp.Validate(); err != nil {
		return fmt.Errorf("invalid product after patch: %w", err)
	}

	*p = tmp

	return nil
}
