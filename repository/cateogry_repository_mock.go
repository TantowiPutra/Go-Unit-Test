package repository

import (
	entity "github.com/TantowiPutra/Go-Unit-Test/category"
	"github.com/stretchr/testify/mock"
)

type CategoryRepositoryMock struct {
	Mock mock.Mock
}

func (repository *CategoryRepositoryMock) FindById(id string) *entity.Category {
	arguments := repository.Mock.Called(id)

	if arguments.Get(0) == nil {
		return nil
	} else {
		category := arguments.Get(0).(entity.Category)

		return &category
	}
}

// Just testing wether a function could accept slice as parameter
// func TestParameterSlice(name []string) {

// }
