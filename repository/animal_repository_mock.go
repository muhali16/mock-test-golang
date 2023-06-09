package repository

import (
	"mock-golang/entity"

	"github.com/stretchr/testify/mock"
)

type AnimalRepositoryMock struct {
	Mock mock.Mock
}

func (repository *AnimalRepositoryMock) FindById(id int) *entity.Animal {
	arguments := repository.Mock.Called(id)
	if arguments.Get(0) == nil {
		return nil
	} else {
		animal := arguments.Get(0).(entity.Animal)
		return &animal
	}
}
