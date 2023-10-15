package memory

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"playground/newProject/models"
	"time"

	"github.com/google/uuid"
)

type branchProductTransactionRepo struct {
	fileName string
}

func NewBranchProductTransactionRepo(fn string) *branchProductTransactionRepo {
	return &branchProductTransactionRepo{fileName: fn}
}

// CreateBranch method creates new branch with given name and address and returns its id
func (b *branchProductTransactionRepo) CreateBranchProductTransaction(req models.CreateBranchProductTransaction) (string, error) {

	branchProducts, err := b.read()
	if err != nil {
		return "", err
	}

	id := uuid.NewString()
	branchProducts = append(branchProducts, models.BranchProductTransaction{
		ID:        id,
		BranchID:  req.BranchID,
		ProductID: req.ProductID,
		Quantity:  req.Quantity,
		UserID:    req.UserID,
		Type:      req.Type,
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
	})

	err = b.write(branchProducts)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (b *branchProductTransactionRepo) UpdateBranchProductTransaction(req models.BranchProductTransaction) (msg string, err error) {

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

func (b *branchProductTransactionRepo) GetBranchProductTransaction(req models.IdRequest) (resp models.BranchProductTransaction, err error) {
	brancheProductTransactions, err := b.read()
	if err != nil {
		return models.BranchProductTransaction{}, err
	}
	for _, v := range brancheProductTransactions {
		if v.ID == req.Id {
			return v, nil
		}
	}
	return models.BranchProductTransaction{}, errors.New("not found")
}
func (b *branchProductTransactionRepo) GetAllBranchProductTransaction(req models.GetAllBranchProductTransactionRequest) (resp models.GetAllBranchProductTransaction, err error) {
	branchProductTransactions, err := b.read()
	if err != nil {
		return resp, err
	}
	start := req.Limit * (req.Page - 1)
	end := start + req.Limit
	if start > len(branchProductTransactions) {
		resp.BranchProductTransactions = []models.BranchProductTransaction{}
		resp.Count = len(branchProductTransactions)
		return resp, nil
	} else if end > len(branchProductTransactions) {
		return models.GetAllBranchProductTransaction{
			BranchProductTransactions: branchProductTransactions[start:],
			Count:                     len(branchProductTransactions),
		}, nil
	}

	return models.GetAllBranchProductTransaction{
		BranchProductTransactions: branchProductTransactions[start:end],
		Count:                     len(branchProductTransactions)}, nil
}
func (b *branchProductTransactionRepo) DeleteBranchProductTransaction(req models.IdRequest) (string, error) {
	branchProductTransactions, err := b.read()
	if err != nil {
		return "", err
	}
	for i, v := range branchProductTransactions {
		if v.ID == req.Id {
			if i == (len(branchProductTransactions) - 1) {
				branchProductTransactions = branchProductTransactions[:i]
			} else {
				branchProductTransactions = append(branchProductTransactions[:i], branchProductTransactions[i+1:]...)
			}
			err = b.write(branchProductTransactions)
			if err != nil {
				return "", err
			}
			return "deleted successfully", nil
		}
	}

	return "", errors.New("not found")
}

func (u *branchProductTransactionRepo) read() ([]models.BranchProductTransaction, error) {
	var (
		branchProductTransactions []models.BranchProductTransaction
	)

	data, err := os.ReadFile(u.fileName)
	if err != nil {
		log.Printf("Error while Read data: %+v", err)
		return nil, err
	}

	err = json.Unmarshal(data, &branchProductTransactions)
	if err != nil {
		log.Printf("Error while Unmarshal data: %+v", err)
		return nil, err
	}

	return branchProductTransactions, nil
}

func (u *branchProductTransactionRepo) write(branchProductTransactions []models.BranchProductTransaction) error {

	body, err := json.Marshal(branchProductTransactions)
	if err != nil {
		return err
	}

	err = os.WriteFile(u.fileName, body, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}
