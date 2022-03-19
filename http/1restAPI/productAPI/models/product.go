package models

type Product struct {
	Id    int     `json:"id"`
	Name  string  `json:"name"`
	Stock int     `json:"stock"`
	Price float64 `json:"price"`
}

// CRUD

func CreateReadProduct(products *[]Product, newProduc *Product) (err error) {
	*products = append(*products, *newProduc)
	return nil
}

func ReadProductByID(products *[]Product, id int) (product Product) {
	for _, item := range *products {
		if item.Id == id {
			return item
		}
	}
	return Product{}
}

func UpdatwProductByID(products *[]Product, updateProduct *Product) (product Product) {
	for index, item := range *products {
		if item.Id == (*updateProduct).Id {
			(*products)[index].Name = (*updateProduct).Name
			(*products)[index].Stock = (*updateProduct).Stock
			(*products)[index].Price = (*updateProduct).Price

			return (*products)[index]
		}
	}
	return Product{}
}

func DeleteProductById(product *[]Product, id int) (product Product) {
	var deleteProduct []Product
	for index, item := range *products {
		if item.Id == id {
			n := len(*products)
			deleteProduct = item
			(*products)[index] = (*products)[n-1]
			return deleteProduct
		}
	}
	return Product{}
}
