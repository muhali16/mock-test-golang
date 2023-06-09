package repository

import "mock-golang/entity"

type AnimalRepository interface {
	FindById(id int) *entity.Animal
}
