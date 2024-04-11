package service

import (
	"testing"

	entity "github.com/TantowiPutra/Go-Unit-Test/category"
	"github.com/TantowiPutra/Go-Unit-Test/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var categoryRepository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}
var categoryService = CategoryService{repository: categoryRepository}

func TestCategoryService_Get(t *testing.T) {
	categoryRepository.Mock.On("FindById", "1").Return(nil)

	// ===== Compare with Traditional Method =================================
	category, err := categoryService.Get("1")

	assert.Nil(t, category)
	assert.NotNil(t, err)
}

func TestCategoryService_GetFound(t *testing.T) {
	category := entity.Category{
		Id:   "2",
		Name: "Handphone",
	}

	// When method FindById is invoked, with "certain value", set the return value
	// To store the mock data into mock object, utilize <objectName>.Mock.On("<FunctionName>", <param>).Return<ExpectedValue>, since it's stored in pointer address, it serve as a global variable
	// To invoke the saved mock data, utilize <objectName>.Mock.Called(param), if it's stored it will give the stored mock data as return value
	categoryRepository.Mock.On("FindById", "1").Return(category)
	categoryRepository.Mock.On("FindById", "2").Return(category)
	categoryRepository.Mock.On("FindById2", "2").Return(category)

	result, err := categoryService.Get("1")
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, category.Id, result.Id)
	assert.Equal(t, category.Name, result.Name)
}
