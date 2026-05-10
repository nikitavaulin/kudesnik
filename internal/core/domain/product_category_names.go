package domain

type ProductCategoryName string

const (
	WindowsCategory        ProductCategoryName = "Окна"
	BalconiesCategory      ProductCategoryName = "Балконы"
	DoorsCategory          ProductCategoryName = "Двери"
	EntranceDoorsCategory  ProductCategoryName = "Двери входные"
	InteriourDoorsCategory ProductCategoryName = "Двери межкомнатные"
	OthersCategory         ProductCategoryName = "Другое"
)

var categoriesEnum map[ProductCategoryName]any = map[ProductCategoryName]any{
	WindowsCategory:        struct{}{},
	DoorsCategory:          struct{}{},
	EntranceDoorsCategory:  struct{}{},
	InteriourDoorsCategory: struct{}{},
	OthersCategory:         struct{}{},
}

func IsCategoryName(name string) bool {
	_, ok := categoriesEnum[ProductCategoryName(name)]
	return ok
}

func GetCategoryName(name string) ProductCategoryName {
	if IsCategoryName(name) {
		return ProductCategoryName(name)
	}
	return OthersCategory
}
