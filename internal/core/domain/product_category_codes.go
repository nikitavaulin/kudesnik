package domain

import (
	"fmt"
)

type ProductCategoryCode string

const (
	WindowsCategory       ProductCategoryCode = "windows"
	BalconiesCategory     ProductCategoryCode = "balconies"
	DoorsCategory         ProductCategoryCode = "doors"
	EntranceDoorsCategory ProductCategoryCode = "entrance-doors"
	InteriorDoorsCategory ProductCategoryCode = "interior-doors"
	OthersCategory        ProductCategoryCode = "others"
)

const (
	MinCategoryCodeLength = 3
	MaxCategoryCodeLength = 30
)

var categoriesEnum map[ProductCategoryCode]any = map[ProductCategoryCode]any{
	WindowsCategory:       struct{}{},
	DoorsCategory:         struct{}{},
	EntranceDoorsCategory: struct{}{},
	InteriorDoorsCategory: struct{}{},
	OthersCategory:        struct{}{},
}

func ValidateCategoryCode(code string) error {
	_, ok := categoriesEnum[ProductCategoryCode(code)]
	if !ok {
		return fmt.Errorf("unknown category code")
	}
	return nil
}

func GetCategoryCode(code string) ProductCategoryCode {
	if err := ValidateCategoryCode(code); err != nil {
		return OthersCategory
	}
	return ProductCategoryCode(code)
}
