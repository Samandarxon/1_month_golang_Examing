package models

type BranchProductTransaction struct {
	ID        string `json:"id"`
	BranchID  string `json:"branch_id"`
	UserID    string `json:"user_id"`
	ProductID string `json:"product_id"`
	Type      string `json:"type"`
	Quantity  int    `json:"quantity"`
	CreatedAt string `json:"created_at"`
}

type CreateBranchProductTransaction struct {
	BranchID  string `json:"branch_id"`
	UserID    string `json:"user_id"`
	ProductID string `json:"product_id"`
	Type      string `json:"type"`
	Quantity  int    `json:"quantity"`
	CreatedAt string `json:"created_at"`
}

type GetAllBranchProductTransactionRequest struct {
	Page      int
	Limit     int
	BranchID  string
	UserID    string
	ProductID string
	Type      string
}
type GetAllBranchProductTransaction struct {
	BranchProductTransactions []BranchProductTransaction
	Count                     int
}
