package memory

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"playground/newProject/models"

	"github.com/google/uuid"
)

type userRepo struct {
	fileName string
}

func NewUserRepo(fn string) *userRepo {
	return &userRepo{fileName: fn}
}

//CreateBranch method creates new branch with given name and address and returns its id
func (b *userRepo) CreateUser(req models.CreateUser) (string, error) {

	users, err := b.read()
	if err != nil {
		return "", err
	}

	id := uuid.NewString()
	users = append(users, models.User{
		Id:   id,
		Name: req.Name,
	})

	err = b.write(users)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (b *userRepo) UpdateUser(req models.User) (msg string, err error) {

	users, err := b.read()
	if err != nil {
		return "", err
	}

	for i, v := range users {
		if v.Id == req.Id {
			users[i] = req
			msg = "updated successfully"
			err = b.write(users)
			if err != nil {
				return "", err
			}
			return
		}
	}
	return "", errors.New("not found")
}

func (b *userRepo) GetUser(req models.IdRequest) (resp models.User, err error) {
	users, err := b.read()
	if err != nil {
		return models.User{}, err
	}
	for _, v := range users {
		if v.Id == req.Id {
			return v, nil
		}
	}
	return models.User{}, errors.New("not found")
}
func (b *userRepo) GetAllUser(req models.GetAllUserRequest) (resp models.GetAllUser, err error) {
	users, err := b.read()
	if err != nil {
		return resp, err
	}
	start := req.Limit * (req.Page - 1)
	end := start + req.Limit
	if start > len(users) {
		resp.Users = []models.User{}
		resp.Count = len(users)
		return resp, nil
	} else if end > len(users) {
		return models.GetAllUser{
			Users: users[start:],
			Count: len(users),
		}, nil
	}

	return models.GetAllUser{
		Users: users[start:end],
		Count: len(users)}, nil
}
func (b *userRepo) DeleteUser(req models.IdRequest) (string, error) {
	users, err := b.read()
	if err != nil {
		return "", err
	}
	for i, v := range users {
		if v.Id == req.Id {
			if i == (len(users) - 1) {
				users = users[:i]
			} else {
				users = append(users[:i], users[i+1:]...)
			}
			err = b.write(users)
			if err != nil {
				return "", err
			}
			return "deleted successfully", nil
		}
	}
	return "", errors.New("not found")
}

func (u *userRepo) read() ([]models.User, error) {
	var (
		users []models.User
	)

	data, err := os.ReadFile(u.fileName)
	if err != nil {
		log.Printf("Error while Read data: %+v", err)
		return nil, err
	}

	err = json.Unmarshal(data, &users)
	if err != nil {
		log.Printf("Error while Unmarshal data: %+v", err)
		return nil, err
	}

	return users, nil
}

func (u *userRepo) write(users []models.User) error {

	body, err := json.Marshal(users)
	if err != nil {
		return err
	}

	err = os.WriteFile(u.fileName, body, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}
