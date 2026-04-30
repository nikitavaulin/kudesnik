package domain

type Product struct {
	ID          ID
	Version     int
	ProductName string
	Price       float64
	Description *string
	IsVisible   bool
	CategoryID  ID
	ProducerID  *ID
}

// func NewProduct(
// 	productName string,
// 	price float64,
// 	description *string,
// 	categoryId ID,
// 	producerId *ID,
// ) *Product {
// 	return &Product{

// 		ProductName: productName,
// 		Price:       price,
// 		Description: description,
// 		CategoryID:  categoryId,
// 		ProducerID:  producerId,
// 	}
// }
