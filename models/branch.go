package models

type CreateBranch struct {
	Name    string
	Address string
}

type Branch struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

type IdRequest struct {
	Id string `json:"id"`
}

type GetAllBranchRequest struct {
	Page  int
	Limit int
	Name  string
}
type GetAllBranch struct {
	Branches []Branch
	Count    int
}

type FileNames struct {
	BranchFile                   string
	BranchProductFile            string
	BranchProductTransactionFile string
	UserFile                     string
	CategoryFile                 string
	ProductFile                  string
}
