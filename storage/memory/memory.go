package memory

import (
	"playground/newProject/models"
	"playground/newProject/storage"
)

type store struct {
	branches                  *branchRepo
	users                     *userRepo
	branchProducts            *branchProductRepo
	branchProductTransactions *branchProductTransactionRepo
	categories                *categoryRepo
	products                  *productRepo
}

func NewStorage(files models.FileNames) storage.StorageI {
	return &store{
		branches:                  NewBranchRepo(files.BranchFile),
		users:                     NewUserRepo(files.UserFile),
		branchProducts:            NewBranchProductRepo(files.BranchProductFile),
		branchProductTransactions: NewBranchProductTransactionRepo(files.BranchProductTransactionFile),
		categories:                NewCategoryRepo(files.CategoryFile),
		products:                  NewProductRepo(files.ProductFile),
	}
}

func (s *store) Branch() storage.BranchesI {
	return s.branches
}
func (s *store) BranchProduct() storage.BranchesProductsI {
	return s.branchProducts
}
func (s *store) BranchProductTransaction() storage.BranchProductTransactionI {
	return s.branchProductTransactions
}
func (s *store) User() storage.UsersI {
	return s.users
}

func (s *store) Category() storage.CategoryI {
	return s.categories
}
func (s *store) Product() storage.ProductI {
	return s.products
}
