package handler

import (
	"fmt"
	"playground/newProject/models"
)

func (h *handler) CreateBranch(name, address string) {
	resp, err := h.strg.Branch().CreateBranch(models.CreateBranch{
		Name:    name,
		Address: address,
	})
	if err != nil {
		fmt.Println("error from CreateBranch:", err.Error())
		return
	}
	fmt.Println("created new branch with id:", resp)
}
func (h *handler) UpdateBranch(id string, name, address string) {
	resp, err := h.strg.Branch().UpdateBranch(models.Branch{
		Id:      id,
		Name:    name,
		Address: address,
	})
	if err != nil {
		fmt.Println("error from UpdateBranch:", err.Error())
		return
	}
	fmt.Println(resp)
}

func (h *handler) GetBranch(id string) models.Branch {
	resp, err := h.strg.Branch().GetBranch(models.IdRequest{Id: id})
	if err != nil {
		fmt.Println("error from GetBranch:", err.Error())
		return models.Branch{}
	}
	// fmt.Println(resp)
	return resp
}

func (h *handler) GetAllBranch(page, limit int, search string) {
	if page < 1 {
		page = h.cfg.Page
	}
	if limit < 1 {
		limit = h.cfg.Limit
	}
	resp, err := h.strg.Branch().GetAllBranch(models.GetAllBranchRequest{
		Page:  page,
		Limit: limit,
		Name:  search,
	})
	if err != nil {
		fmt.Println("error from GetAllBranch:", err.Error())
		return
	}
	fmt.Println(resp)
}
func (h *handler) DeleteBranch(id string) {
	resp, err := h.strg.Branch().DeleteBranch(models.IdRequest{Id: id})
	if err != nil {
		fmt.Println("error from DeleteBranch:", err.Error())
		return
	}
	fmt.Println(resp)
}
