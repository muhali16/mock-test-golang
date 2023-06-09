package service

import (
	"errors"
	"mock-golang/entity"
	"mock-golang/repository"
)

type AnimalService struct {
	Repository repository.AnimalRepository
}

func (service *AnimalService) Get(id int) (*entity.Animal, error) {
	animal := service.Repository.FindById(id)
	if animal == nil {
		return nil, errors.New("Data not found")
	}
	return animal, nil
}
