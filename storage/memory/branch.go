package memory

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"playground/newProject/models"

	"github.com/google/uuid"
)

type branchRepo struct {
	fileName string
}

func NewBranchRepo(fn string) *branchRepo {
	return &branchRepo{fileName: fn}
}

//CreateBranch method creates new branch with given name and address and returns its id
func (b *branchRepo) CreateBranch(req models.CreateBranch) (string, error) {

	branches, err := b.read()
	if err != nil {
		return "", err
	}

	id := uuid.NewString()
	branches = append(branches, models.Branch{
		Id:      id,
		Name:    req.Name,
		Address: req.Address,
	})

	err = b.write(branches)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (b *branchRepo) UpdateBranch(req models.Branch) (msg string, err error) {

	branches, err := b.read()
	if err != nil {
		return "", err
	}

	for i, v := range branches {
		if v.Id == req.Id {
			branches[i] = req
			msg = "updated successfully"
			err = b.write(branches)
			if err != nil {
				return "", err
			}
			return
		}
	}
	return "", errors.New("not found")
}

func (b *branchRepo) GetBranch(req models.IdRequest) (resp models.Branch, err error) {
	branches, err := b.read()
	if err != nil {
		return models.Branch{}, err
	}
	for _, v := range branches {
		if v.Id == req.Id {
			return v, nil
		}
	}
	return models.Branch{}, errors.New("not found")
}
func (b *branchRepo) GetAllBranch(req models.GetAllBranchRequest) (resp models.GetAllBranch, err error) {
	branches, err := b.read()
	if err != nil {
		return resp, err
	}
	start := req.Limit * (req.Page - 1)
	end := start + req.Limit
	if start > len(branches) {
		resp.Branches = []models.Branch{}
		resp.Count = len(branches)
		return resp, nil
	} else if end > len(branches) {
		return models.GetAllBranch{
			Branches: branches[start:],
			Count:    len(branches),
		}, nil
	}

	return models.GetAllBranch{
		Branches: branches[start:end],
		Count:    len(branches)}, nil
}
func (b *branchRepo) DeleteBranch(req models.IdRequest) (string, error) {
	branches, err := b.read()
	if err != nil {
		return "", err
	}
	for i, v := range branches {
		if v.Id == req.Id {
			if i == (len(branches) - 1) {
				branches = branches[:i]
			} else {
				branches = append(branches[:i], branches[i+1:]...)
			}
			err = b.write(branches)
			if err != nil {
				return "", err
			}
			return "deleted successfully", nil
		}
	}
	return "", errors.New("not found")

}

func (u *branchRepo) read() ([]models.Branch, error) {
	var (
		branches []models.Branch
	)

	data, err := os.ReadFile(u.fileName)
	if err != nil {
		log.Printf("Error while Read data: %+v", err)
		return nil, err
	}

	err = json.Unmarshal(data, &branches)
	if err != nil {
		log.Printf("Error while Unmarshal data: %+v", err)
		return nil, err
	}

	return branches, nil
}

func (u *branchRepo) write(branches []models.Branch) error {

	body, err := json.Marshal(branches)
	if err != nil {
		return err
	}

	err = os.WriteFile(u.fileName, body, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}
