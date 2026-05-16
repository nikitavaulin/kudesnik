package domain

import (
	"fmt"

	core_errors "github.com/nikitavaulin/kudesnik/internal/core/errors"
)

type Window struct {
	ProductBase
	Purpose  string `json:"purpose"`
	Width    int    `json:"width"`
	Height   int    `json:"height"`
	Material string `json:"material"`
}

type WindowDetailed struct {
	Window
	ProductDetails
}

func NewWindow(base ProductBase, width, height int, purpose, material string) *Window {
	return &Window{
		ProductBase: base,
		Width:       width,
		Height:      height,
		Purpose:     purpose,
		Material:    material,
	}
}

func (w *Window) Validate() error {
	if err := w.ProductBase.Validate(); err != nil {
		return fmt.Errorf("product base validation failed: %w", err)
	}

	if w.Width < 0 {
		return fmt.Errorf("width cannot be non-negative: %w", core_errors.ErrInvalidArgument)
	}

	if w.Height < 0 {
		return fmt.Errorf("height cannot be non-negative: %w", core_errors.ErrInvalidArgument)
	}

	if w.Purpose == "" {
		return fmt.Errorf("purpose is required: %w", core_errors.ErrInvalidArgument)
	}

	return nil
}

func (w *Window) GetBase() *ProductBase                { return &w.ProductBase }
func (w *Window) GetCategoryName() ProductCategoryCode { return WindowsCategory }

type WindowPatch struct {
	ProductBasePatch
	Purpose  Nullable[string] `json:"purpose" swaggertype:"string"`
	Width    Nullable[int]    `json:"width" swaggertype:"integer"`
	Height   Nullable[int]    `json:"height" swaggertype:"integer"`
	Material Nullable[string] `json:"material" swaggertype:"string"`
}

func (w *WindowPatch) Validate() error {
	if err := w.ProductBasePatch.Validate(); err != nil {
		return fmt.Errorf("product base validation failed: %w", err)
	}

	if w.Width.Set && *w.Width.Value < 0 {
		return fmt.Errorf("width cannot be non-negative: %w", core_errors.ErrInvalidArgument)
	}

	if w.Height.Set && *w.Height.Value < 0 {
		return fmt.Errorf("height cannot be non-negative: %w", core_errors.ErrInvalidArgument)
	}

	if w.Purpose.Set && *w.Purpose.Value == "" {
		return fmt.Errorf("purpose is required: %w", core_errors.ErrInvalidArgument)
	}

	return nil
}

func (w *Window) ApplyPatch(windowPatch *WindowPatch) error {
	if err := windowPatch.Validate(); err != nil {
		return fmt.Errorf("invalid window patch: %w", err)
	}

	tmp := *w

	if err := tmp.ProductBase.ApplyPatch(&windowPatch.ProductBasePatch); err != nil {
		return fmt.Errorf("failed to patch product base: %w", err)
	}

	if windowPatch.Purpose.Set {
		tmp.Purpose = *windowPatch.Purpose.Value
	}

	if windowPatch.Width.Set {
		tmp.Width = *windowPatch.Width.Value
	}

	if windowPatch.Height.Set {
		tmp.Height = *windowPatch.Height.Value
	}

	if windowPatch.Material.Set {
		tmp.Material = *windowPatch.Material.Value
	}

	if err := tmp.Validate(); err != nil {
		return fmt.Errorf("invalid window after patch: %w", err)
	}

	*w = tmp

	return nil
}
