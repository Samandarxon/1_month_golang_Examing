package memory

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"playground/newProject/models"

	"github.com/google/uuid"
)

type branchProductRepo struct {
	fileName string
}

func NewBranchProductRepo(fn string) *branchProductRepo {
	return &branchProductRepo{fileName: fn}
}

//CreateBranch method creates new branch with given name and address and returns its id
func (b *branchProductRepo) CreateBranchProduct(req models.CreateBranchProduct) (string, error) {

	branchProducts, err := b.read()
	if err != nil {
		return "", err
	}

	id := uuid.NewString()
	branchProducts = append(branchProducts, models.BranchProduct{
		ID:        id,
		BranchID:  req.BranchID,
		ProductID: req.ProductID,
		Quantity:  req.Quantity,
	})

	err = b.write(branchProducts)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (b *branchProductRepo) UpdateBranchProduct(req models.BranchProduct) (msg string, err error) {

	branchProducts, err := b.read()
	if err != nil {
		return "", err
	}

	for i, v := range branchProducts {
		if v.ID == req.ID {
			branchProducts[i] = req
			msg = "updated successfully"
			err = b.write(branchProducts)
			if err != nil {
				return "", err
			}
			return
		}
	}
	return "", errors.New("not found")
}

func (b *branchProductRepo) GetBranchProduct(req models.IdRequest) (resp models.BranchProduct, err error) {
	branchProducts, err := b.read()
	if err != nil {
		return models.BranchProduct{}, err
	}
	for _, v := range branchProducts {
		if v.ID == req.Id {
			return v, nil
		}
	}
	return models.BranchProduct{}, errors.New("not found")
}
func (b *branchProductRepo) GetAllBranchProduct(req models.GetAllBranchProductRequest) (resp models.GetAllBranchProduct, err error) {
	branchProducts, err := b.read()
	if err != nil {
		return resp, err
	}
	start := req.Limit * (req.Page - 1)
	end := start + req.Limit
	if start > len(branchProducts) {
		resp.BranchProducts = []models.BranchProduct{}
		resp.Count = len(branchProducts)
		return resp, nil
	} else if end > len(branchProducts) {
		return models.GetAllBranchProduct{
			BranchProducts: branchProducts[start:],
			Count:          len(branchProducts),
		}, nil
	}

	return models.GetAllBranchProduct{
		BranchProducts: branchProducts[start:end],
		Count:          len(branchProducts)}, nil
}
func (b *branchProductRepo) DeleteBranchProduct(req models.IdRequest) (string, error) {
	branchProducts, err := b.read()
	if err != nil {
		return "", err
	}
	for i, v := range branchProducts {
		if v.ID == req.Id {
			if i == (len(branchProducts) - 1) {
				branchProducts = branchProducts[:i]
			} else {
				branchProducts = append(branchProducts[:i], branchProducts[i+1:]...)
			}
			err = b.write(branchProducts)
			if err != nil {
				return "", err
			}
			return "deleted successfully", nil
		}
	}

	return "", errors.New("not found")
}

func (u *branchProductRepo) read() ([]models.BranchProduct, error) {
	var (
		branchProducts []models.BranchProduct
	)

	data, err := os.ReadFile(u.fileName)
	if err != nil {
		log.Printf("Error while Read data: %+v", err)
		return nil, err
	}

	err = json.Unmarshal(data, &branchProducts)
	if err != nil {
		log.Printf("Error while Unmarshal data: %+v", err)
		return nil, err
	}

	return branchProducts, nil
}

func (u *branchProductRepo) write(branchProducts []models.BranchProduct) error {

	body, err := json.Marshal(branchProducts)
	if err != nil {
		return err
	}

	err = os.WriteFile(u.fileName, body, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}
