package models

type CreateBranchProduct struct {
	BranchID  string `json:"branch_id"`
	ProductID string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}

type BranchProduct struct {
	ID        string `json:"id"`
	BranchID  string `json:"branch_id"`
	ProductID string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}

type GetAllBranchProductRequest struct {
	Page      int
	Limit     int
	BranchID  string
	ProductID string
}
type GetAllBranchProduct struct {
	BranchProducts []BranchProduct
	Count          int
}
