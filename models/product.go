package models

type Product struct {
	Id         string `json:"id"`
	CategoryId string `json:"category_id"`
	Name       string `json:"name"`
	Price      int    `json:"price"`
}

type CreateProduct struct {
	CategoryId string `json:"category_id"`
	Name       string `json:"name"`
	Price      int    `json:"price"`
}
type GetAllProductRequest struct {
	Page  int
	Limit int
}
type GetAllProduct struct {
	Products []Product
	Count    int
}

type PrimeryKeyProduct struct {
	Id string `json:"id"`
}

type UpadetProduct struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
