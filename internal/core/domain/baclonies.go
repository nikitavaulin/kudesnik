package domain

import (
	"fmt"

	core_errors "github.com/nikitavaulin/kudesnik/internal/core/errors"
	core_validation "github.com/nikitavaulin/kudesnik/internal/core/tools/validation"
)

const (
	MinBalconyPurposeLength = 1
	MaxBalconyPurposeLength = 100

	MinBalconyMaterialLength = 1
	MaxBalconyMaterialLength = 100
)

// Balcony основная структура балкона
type Balcony struct {
	ProductBase
	Purpose  string `json:"purpose"`
	Material string `json:"material"`
}

// NewBalcony конструктор для создания нового балкона
func NewBalcony(base ProductBase, purpose, material string) *Balcony {
	return &Balcony{
		ProductBase: base,
		Purpose:     purpose,
		Material:    material,
	}
}

// Validate проверяет корректность данных балкона
func (b *Balcony) Validate() error {
	if err := b.ProductBase.Validate(); err != nil {
		return fmt.Errorf("product base validation failed: %w", err)
	}

	if err := core_validation.ValidateIntInBounds(len(b.Purpose), MinBalconyPurposeLength, MaxBalconyPurposeLength); err != nil {
		return fmt.Errorf("invalid purpose length: %v: %w", err, core_errors.ErrInvalidArgument)
	}

	if err := core_validation.ValidateIntInBounds(len(b.Material), MinBalconyMaterialLength, MaxBalconyMaterialLength); err != nil {
		return fmt.Errorf("invalid material length: %v: %w", err, core_errors.ErrInvalidArgument)
	}

	return nil
}

func (b *Balcony) GetBase() *ProductBase {
	return &b.ProductBase
}

func (b *Balcony) GetCategoryName() ProductCategoryCode {
	return BalconiesCategory
}

type BalconyPatch struct {
	ProductBasePatch
	Purpose  Nullable[string] `json:"purpose" swaggertype:"string"`
	Material Nullable[string] `json:"material" swaggertype:"string"`
}

// Validate проверяет корректность патча
func (p *BalconyPatch) Validate() error {
	if err := p.ProductBasePatch.Validate(); err != nil {
		return fmt.Errorf("product base validation failed: %w", err)
	}

	if p.Purpose.Set {
		if err := core_validation.ValidateIntInBounds(len(*p.Purpose.Value), MinBalconyPurposeLength, MaxBalconyPurposeLength); err != nil {
			return fmt.Errorf("invalid purpose length: %v: %w", err, core_errors.ErrInvalidArgument)
		}
	}

	if p.Material.Set {
		if err := core_validation.ValidateIntInBounds(len(*p.Material.Value), MinBalconyMaterialLength, MaxBalconyMaterialLength); err != nil {
			return fmt.Errorf("invalid material length: %v: %w", err, core_errors.ErrInvalidArgument)
		}
	}

	return nil
}

// ApplyPatch применяет патч к балкону
func (b *Balcony) ApplyPatch(patch *BalconyPatch) error {
	if err := patch.Validate(); err != nil {
		return fmt.Errorf("invalid balcony patch: %w", err)
	}

	tmp := *b

	if err := tmp.ProductBase.ApplyPatch(&patch.ProductBasePatch); err != nil {
		return fmt.Errorf("failed to patch product base: %w", err)
	}

	if patch.Purpose.Set {
		tmp.Purpose = *patch.Purpose.Value
	}

	if patch.Material.Set {
		tmp.Material = *patch.Material.Value
	}

	if err := tmp.Validate(); err != nil {
		return fmt.Errorf("invalid balcony after patch: %w", err)
	}

	*b = tmp

	return nil
}

func (p *BalconyPatch) isProductPatch() {}
