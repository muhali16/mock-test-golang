package service

import (
	"mock-golang/entity"
	"mock-golang/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var AnimalRepository = &repository.AnimalRepositoryMock{Mock: mock.Mock{}}
var animalService = AnimalService{Repository: AnimalRepository}

func TestAnimalService_GetNotFound(t *testing.T) {
	// melakukan mock
	AnimalRepository.Mock.On("FindById", 1).Return(nil)

	animal1, err := animalService.Get(1)
	assert.Nil(t, animal1)
	assert.NotNil(t, err)
}

func TestAnimalService_GetSuccess(t *testing.T) {
	animal := entity.Animal{
		Id:   103,
		Name: "Kuda",
	}
	AnimalRepository.Mock.On("FindById", 103).Return(animal)

	animal2, err := animalService.Get(103)
	assert.Nil(t, err)
	assert.Equal(t, animal.Id, animal2.Id)
	assert.Equal(t, animal.Name, animal2.Name)
}
