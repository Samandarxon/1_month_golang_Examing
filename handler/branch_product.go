package handler

import (
	"fmt"
	"playground/newProject/models"
)

func (h *handler) CreateBranchProduct(branchId, productId string, quantity int) {
	resp, err := h.strg.BranchProduct().CreateBranchProduct(models.CreateBranchProduct{
		BranchID:  branchId,
		ProductID: productId,
		Quantity:  quantity,
	})
	if err != nil {
		fmt.Println("error from CreateBranch:", err.Error())
		return
	}
	fmt.Println("created new branch with id:", resp)
}
func (h *handler) UpdateBranchProduct(id, branchId, productId string, quantity int) {
	resp, err := h.strg.BranchProduct().UpdateBranchProduct(models.BranchProduct{
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

func (h *handler) GetBranchProduct(id string) {
	resp, err := h.strg.BranchProduct().GetBranchProduct(models.IdRequest{Id: id})
	if err != nil {
		fmt.Println("error from GetBranch:", err.Error())
		return
	}
	fmt.Println(resp)
}

func (h *handler) GetAllBranchProduct(page, limit int, branchId, productId string) {
	if page < 1 {
		page = h.cfg.Page
	}
	if limit < 1 {
		limit = h.cfg.Limit
	}
	resp, err := h.strg.BranchProduct().GetAllBranchProduct(models.GetAllBranchProductRequest{
		Page:      page,
		Limit:     limit,
		BranchID:  branchId,
		ProductID: productId,
	})
	if err != nil {
		fmt.Println("error from GetAllBranch:", err.Error())
		return
	}
	fmt.Println(resp)
}
func (h *handler) DeleteBranchProduct(id string) {
	resp, err := h.strg.BranchProduct().DeleteBranchProduct(models.IdRequest{Id: id})
	if err != nil {
		fmt.Println("error from DeleteBranch:", err.Error())
		return
	}
	fmt.Println(resp)
}
