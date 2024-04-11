package service

import (
	"errors"

	entity "github.com/TantowiPutra/Go-Unit-Test/category"
	"github.com/TantowiPutra/Go-Unit-Test/repository"
)

type CategoryService struct {
	repository repository.CategoryRepository
}

func (service *CategoryService) Get(id string) (*entity.Category, error) {
	category := service.repository.FindById(id)

	if category == nil {
		return category, errors.New("category not found")
	} else {
		return category, nil
	}
}

func TestFunction() int {
	return 123
}
