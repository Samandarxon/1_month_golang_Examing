package handler

import (
	"fmt"
	"playground/newProject/models"
)

func (h *handler) CreateCategory(name string) {
	resp, err := h.strg.Category().CreateCategory(models.CreateCategory{
		Name: name,
	})
	if err != nil {
		fmt.Println("error from CreateCategory:", err.Error())
		return
	}
	fmt.Println("created new Category with id:", resp)
}
func (h *handler) UpdateCategory(id string, name string) {
	resp, err := h.strg.Category().UpdateCategory(models.Category{
		Id:   id,
		Name: name,
	})
	if err != nil {
		fmt.Println("error from UpdateCategory:", err.Error())
		return
	}
	fmt.Println(resp)
}

func (h *handler) GetCategory(id string) models.Category {
	resp, err := h.strg.Category().GetCategory(models.PrimeryKeyCategory{Id: id})
	if err != nil {
		fmt.Println("error from GetCategory:", err.Error())
		return models.Category{}
	}
	// fmt.Println(resp)
	return resp
}

func (h *handler) GetAllCategory(page, limit int) {
	if page < 1 {
		page = h.cfg.Page
	}
	if limit < 1 {
		limit = h.cfg.Limit
	}
	resp, err := h.strg.Category().GetAllCategory(models.GetAllCategoryRequest{
		Page:  page,
		Limit: limit,
	})
	if err != nil {
		fmt.Println("error from GetAllCategory:", err.Error())
		return
	}
	fmt.Println(resp)
}
func (h *handler) DeleteCategory(id string) {
	resp, err := h.strg.Category().DeleteCategory(models.PrimeryKeyCategory{Id: id})
	if err != nil {
		fmt.Println("error from DeleteCategory:", err.Error())
		return
	}
	fmt.Println(resp)
}
