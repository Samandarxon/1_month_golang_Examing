package handler

import (
	"fmt"
	"playground/newProject/models"
)

func (h *handler) CreateBranchProductTransaction(branchId, productId, userId, types string, quantity int) {
	resp, err := h.strg.BranchProductTransaction().CreateBranchProductTransaction(models.CreateBranchProductTransaction{
		BranchID:  branchId,
		ProductID: productId,
		Quantity:  quantity,
		UserID:    userId,
		Type:      types,
	})
	if err != nil {
		fmt.Println("error from CreateBranch:", err.Error())
		return
	}
	fmt.Println("created new branch with id:", resp)
}
func (h *handler) UpdateBranchProductTransaction(id, branchId, productId string, quantity int) {
	resp, err := h.strg.BranchProductTransaction().UpdateBranchProductTransaction(models.BranchProductTransaction{
		ID:        id,
		BranchID:  branchId,
		ProductID: productId,
		Quantity:  quantity,
	})
	if err != nil {
		fmt.Println("error from UpdateBranch:", err.Error())
		return
	}
	fmt.Println(resp)
}

func (h *handler) GetBranchProductTransaction(id string) {
	resp, err := h.strg.BranchProductTransaction().GetBranchProductTransaction(models.IdRequest{Id: id})
	if err != nil {
		fmt.Println("error from GetBranch:", err.Error())
		return
	}
	fmt.Println(resp)
}

func (h *handler) GetAllBranchProductTransaction(page, limit int) models.GetAllBranchProductTransaction {
	if page < 1 {
		page = h.cfg.Page
	}
	if limit < 1 {
		limit = h.cfg.Limit
	}
	resp, err := h.strg.BranchProductTransaction().GetAllBranchProductTransaction(models.GetAllBranchProductTransactionRequest{
		Page:  page,
		Limit: limit,
	})
	if err != nil {
		fmt.Println("error from GetAllBranch:", err.Error())
		return models.GetAllBranchProductTransaction{}
	}
	return resp

}
func (h *handler) DeleteBranchProductTransaction(id string) {
	resp, err := h.strg.BranchProductTransaction().DeleteBranchProductTransaction(models.IdRequest{Id: id})
	if err != nil {
		fmt.Println("error from DeleteBranch:", err.Error())
		return
	}
	fmt.Println(resp)
}
