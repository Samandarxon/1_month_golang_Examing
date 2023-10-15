package memory

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"playground/newProject/models"

	"github.com/google/uuid"
)

type productRepo struct {
	fileName string
}

func NewProductRepo(fn string) *productRepo {
	return &productRepo{fileName: fn}
}

// CreateBranch method creates new branch with given name and address and returns its id
func (b *productRepo) CreateProduct(req models.CreateProduct) (string, error) {

	products, err := b.read()
	fmt.Println(products)

	if err != nil {
		return "", err
	}

	id := uuid.NewString()
	products = append(products, models.Product{
		Id:         id,
		CategoryId: req.CategoryId,
		Name:       req.Name,
		Price:      req.Price,
	})

	err = b.write(products)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (b *productRepo) UpdateProduct(req models.Product) (msg string, err error) {

	products, err := b.read()
	if err != nil {
		return "", err
	}

	for i, v := range products {
		if v.Id == req.Id {
			products[i] = req
			msg = "updated successfully"
			err = b.write(products)
			if err != nil {
				return "", err
			}
			return
		}
	}
	return "", errors.New("not found")
}

func (b *productRepo) GetProduct(req models.PrimeryKeyProduct) (resp models.Product, err error) {
	products, err := b.read()
	if err != nil {
		return models.Product{}, err
	}
	for _, v := range products {
		if v.Id == req.Id {
			return v, nil
		}
	}
	return models.Product{}, errors.New("not found")
}
func (b *productRepo) GetAllProduct(req models.GetAllProductRequest) (resp models.GetAllProduct, err error) {
	products, err := b.read()
	if err != nil {
		return resp, err
	}
	start := req.Limit * (req.Page - 1)
	end := start + req.Limit
	if start > len(products) {
		resp.Products = []models.Product{}
		resp.Count = len(products)
		return resp, nil
	} else if end > len(products) {
		return models.GetAllProduct{
			Products: products[start:],
			Count:    len(products),
		}, nil
	}

	return models.GetAllProduct{
		Products: products[start:end],
		Count:    len(products)}, nil
}
func (b *productRepo) DeleteProduct(req models.PrimeryKeyProduct) (string, error) {
	products, err := b.read()
	if err != nil {
		return "", err
	}
	for i, v := range products {
		if v.Id == req.Id {
			products = append(products[:i], products[i+1:]...)
			err = b.write(products)
			if err != nil {
				return "", err
			}
			return "deleted successfully", nil
		}
	}
	return "", errors.New("not found")
}

func (u *productRepo) read() ([]models.Product, error) {
	var (
		products []models.Product
	)

	data, err := os.ReadFile(u.fileName)
	if err != nil {
		log.Printf("Error while Read data: %+v", err)
		return nil, err
	}
	err = json.Unmarshal(data, &products)
	if err != nil {
		log.Printf("Error while Unmarshal data: %+v", err)
		return nil, err
	}

	return products, nil
}

func (u *productRepo) write(products []models.Product) error {

	body, err := json.Marshal(products)
	if err != nil {
		return err
	}

	err = os.WriteFile(u.fileName, body, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}
