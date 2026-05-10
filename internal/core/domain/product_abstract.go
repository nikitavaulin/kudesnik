package domain

type Product interface {
	GetBase() *ProductBase
	GetCategoryName() ProductCategoryCode
	Validate() error
}

type ProductPatch interface {
	Validate() error
}

func GetProductEmptyInstance(categoryCode string) Product {
	category := GetCategoryCode(categoryCode)
	switch category {
	case WindowsCategory:
		return &Window{}
	default:
		return &ProductBase{}
	}
}

func GetProductPatchEmptyInstance(categoryCode string) ProductPatch {
	category := GetCategoryCode(categoryCode)
	switch category {
	case WindowsCategory:
		return &WindowPatch{}
	default:
		return &ProductBasePatch{}
	}
}
