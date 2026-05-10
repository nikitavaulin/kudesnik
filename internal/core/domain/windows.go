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

	if w.Width <= 0 {
		return fmt.Errorf("width must be greater than 0: %w", core_errors.ErrInvalidArgument)
	}

	if w.Height <= 0 {
		return fmt.Errorf("height must be greater than 0: %w", core_errors.ErrInvalidArgument)
	}

	if w.Purpose == "" {
		return fmt.Errorf("purpose is required: %w", core_errors.ErrInvalidArgument)
	}

	return nil
}

func (w *Window) GetBase() *ProductBase                { return &w.ProductBase }
func (w *Window) GetCategoryName() ProductCategoryName { return WindowsCategory }
