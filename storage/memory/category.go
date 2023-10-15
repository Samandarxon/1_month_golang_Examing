package memory

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"playground/newProject/models"

	"github.com/google/uuid"
)

type categoryRepo struct {
	fileName string
}

func NewCategoryRepo(fn string) *categoryRepo {
	return &categoryRepo{fileName: fn}
}

// CreateBranch method creates new branch with given name and address and returns its id
func (b *categoryRepo) CreateCategory(req models.CreateCategory) (string, error) {

	categorys, err := b.read()
	if err != nil {
		return "", err
	}

	id := uuid.NewString()
	categorys = append(categorys, models.Category{
		Id:   id,
		Name: req.Name,
	})

	err = b.write(categorys)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (b *categoryRepo) UpdateCategory(req models.Category) (msg string, err error) {

	categorys, err := b.read()
	if err != nil {
		return "", err
	}

	for i, v := range categorys {
		if v.Id == req.Id {
			categorys[i] = req
			msg = "updated successfully"
			err = b.write(categorys)
			if err != nil {
				return "", err
			}
			return
		}
	}
	return "", errors.New("not found")
}

func (b *categoryRepo) GetCategory(req models.PrimeryKeyCategory) (resp models.Category, err error) {
	categorys, err := b.read()
	if err != nil {
		return models.Category{}, err
	}
	for _, v := range categorys {
		if v.Id == req.Id {
			return v, nil
		}
	}
	return models.Category{}, errors.New("not found")
}
func (b *categoryRepo) GetAllCategory(req models.GetAllCategoryRequest) (resp models.GetAllCategory, err error) {
	categorys, err := b.read()
	if err != nil {
		return resp, err
	}
	start := req.Limit * (req.Page - 1)
	end := start + req.Limit
	if start > len(categorys) {
		resp.Categories = []models.Category{}
		resp.Count = len(categorys)
		return resp, nil
	} else if end > len(categorys) {
		return models.GetAllCategory{
			Categories: categorys[start:],
			Count:      len(categorys),
		}, nil
	}

	return models.GetAllCategory{
		Categories: categorys[start:end],
		Count:      len(categorys)}, nil
}
func (b *categoryRepo) DeleteCategory(req models.PrimeryKeyCategory) (string, error) {
	categorys, err := b.read()
	if err != nil {
		return "", err
	}
	for i, v := range categorys {
		if v.Id == req.Id {

			categorys = append(categorys[:i], categorys[i+1:]...)
			err = b.write(categorys)
			if err != nil {
				return "", err
			}
			err = b.write(categorys)
			if err != nil {
				return "", err
			}
			return "deleted successfully", nil
		}
	}
	return "", errors.New("not found")
}

func (u *categoryRepo) read() ([]models.Category, error) {
	var (
		categorys []models.Category
	)

	data, err := os.ReadFile(u.fileName)
	if err != nil {
		log.Printf("Error while Read data: %+v", err)
		return nil, err
	}

	err = json.Unmarshal(data, &categorys)
	if err != nil {
		log.Printf("Error while Unmarshal data: %+v", err)
		return nil, err
	}

	return categorys, nil
}

func (u *categoryRepo) write(categorys []models.Category) error {

	body, err := json.Marshal(categorys)
	if err != nil {
		return err
	}

	err = os.WriteFile(u.fileName, body, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}
