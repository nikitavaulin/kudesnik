package domain

type Product interface {
	GetBase() *ProductBase
	GetCategoryName() ProductCategoryCode
	Validate() error
}

type ProductPatch interface {
	Validate() error
	isProductPatch()
}

func GetProductEmptyInstance(categoryCode string) Product {
	category := GetCategoryCode(categoryCode)
	switch category {
	case WindowsCategory:
		return &Window{}
	case EntranceDoorsCategory:
		return &EntranceDoor{}
	case InteriorDoorsCategory:
		return &InteriorDoor{}
	default:
		return &ProductBase{}
	}
}

func GetProductPatchEmptyInstance(categoryCode string) ProductPatch {
	category := GetCategoryCode(categoryCode)
	switch category {
	case WindowsCategory:
		return &WindowPatch{}
	case EntranceDoorsCategory:
		return &EntranceDoorPatch{}
	case InteriorDoorsCategory:
		return &InteriorDoorPatch{}
	default:
		return &ProductBasePatch{}
	}
}
