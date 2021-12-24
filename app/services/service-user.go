package services

import (
	"errors"

	"github.com/Holilamd/module-pemeliharaan/app/models"
	"github.com/Holilamd/module-pemeliharaan/app/repositorys"
)

type Serviceuser interface {
	All() ([]models.User, error)
	Create(input models.Userinput) (models.Userinput, error)
}
type serviceuser struct {
	repository repositorys.Repositoryuser
}

func NewServiceUser(repository repositorys.Repositoryuser) *serviceuser {
	return &serviceuser{repository}
}
func (s *serviceuser) All() ([]models.User, error) {
	result, err := s.repository.All()
	if err != nil {
		return result, err
	}
	return result, nil
}

func (s *serviceuser) Create(input models.Userinput) (models.Userinput, error) {
	avalaibleMail, _ := s.repository.FindByEmail(input.Email)
	if avalaibleMail.Email != `` {
		return avalaibleMail, errors.New("Email allready")
	}

	result, err := s.repository.Create(input)
	if err != nil {
		return result, err
	}
	return result, nil
}
