package handler

import (
	"fmt"
	"playground/newProject/models"
)

func (h *handler) CreateUser(name, address string) {
	resp, err := h.strg.User().CreateUser(models.CreateUser{
		Name: name,
	})
	if err != nil {
		fmt.Println("error from CreateUser:", err.Error())
		return
	}
	fmt.Println("created new User with id:", resp)
}
func (h *handler) UpdateUser(id string, name, address string) {
	resp, err := h.strg.User().UpdateUser(models.User{
		Id:   id,
		Name: name,
	})
	if err != nil {
		fmt.Println("error from UpdateUser:", err.Error())
		return
	}
	fmt.Println(resp)
}

func (h *handler) GetUser(id string) models.User {
	resp, err := h.strg.User().GetUser(models.IdRequest{Id: id})
	if err != nil {
		fmt.Println("error from GetUser:", err.Error())
		return models.User{}
	}
	// fmt.Println(resp)
	return resp
}

func (h *handler) GetAllUser(page, limit int, search string) {
	if page < 1 {
		page = h.cfg.Page
	}
	if limit < 1 {
		limit = h.cfg.Limit
	}
	resp, err := h.strg.User().GetAllUser(models.GetAllUserRequest{
		Page:  page,
		Limit: limit,
		Name:  search,
	})
	if err != nil {
		fmt.Println("error from GetAllUser:", err.Error())
		return
	}
	fmt.Println(resp)
}
func (h *handler) DeleteUser(id string) {
	resp, err := h.strg.User().DeleteUser(models.IdRequest{Id: id})
	if err != nil {
		fmt.Println("error from DeleteUser:", err.Error())
		return
	}
	fmt.Println(resp)
}
