package models

type Category struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
type CreateCategory struct {
	Name string
}
type GetAllCategoryRequest struct {
	Page  int
	Limit int
}
type GetAllCategory struct {
	Categories []Category
	Count      int
}

type PrimeryKeyCategory struct {
	Id string `json:"id"`
}

type UpadetCategory struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
