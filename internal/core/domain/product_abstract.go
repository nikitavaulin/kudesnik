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

type ProductDetailed struct {
	Product        `json:"product"`
	ProductDetails `json:"details"`
}

type ProductDetails struct {
	CategoryName        string  `json:"category_name"`
	InstallationPrice   float64 `json:"installation_price"`
	ProducerCompanyName *string `json:"producer_company_name,omitempty"`
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
	case BalconiesCategory:
		return &Balcony{}
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
	case BalconiesCategory:
		return &BalconyPatch{}
	default:
		return &ProductBasePatch{}
	}
}
