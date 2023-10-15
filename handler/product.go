package handler

import (
	"fmt"
	"playground/newProject/models"
)

func (h *handler) CreateProduct(categoryId, name string, price int) {
	fmt.Println(categoryId)
	resp, err := h.strg.Product().CreateProduct(models.CreateProduct{
		CategoryId: categoryId,
		Name:       name,
		Price:      price,
	})
	if err != nil {
		fmt.Println("error from CreateProduct:", err.Error())
		return
	}
	fmt.Println("created new Product with id:", resp)
}
func (h *handler) UpdateProduct(id string, categoryId, name string, price int) {
	resp, err := h.strg.Product().UpdateProduct(models.Product{
		Id:         id,
		CategoryId: categoryId,
		Name:       name,
		Price:      price,
	})
	if err != nil {
		fmt.Println("error from UpdateProduct:", err.Error())
		return
	}
	fmt.Println(resp)
}

func (h *handler) GetProduct(id string) models.Product {
	resp, err := h.strg.Product().GetProduct(models.PrimeryKeyProduct{Id: id})
	if err != nil {
		fmt.Println("error from GetProduct:", err.Error())
		return models.Product{}
	}
	return resp
}

func (h *handler) GetAllProduct(page, limit int) {
	if page < 1 {
		page = h.cfg.Page
	}
	if limit < 1 {
		limit = h.cfg.Limit
	}
	resp, err := h.strg.Product().GetAllProduct(models.GetAllProductRequest{
		Page:  page,
		Limit: limit,
	})
	if err != nil {
		fmt.Println("error from GetAllProduct:", err.Error())
		return
	}
	for index, el := range resp.Products {
		fmt.Println(index+1, el)
	}
}
func (h *handler) DeleteProduct(id string) {
	resp, err := h.strg.Product().DeleteProduct(models.PrimeryKeyProduct{Id: id})
	if err != nil {
		fmt.Println("error from DeleteProduct:", err.Error())
		return
	}
	fmt.Println(resp)
}
