package service

import (
	"go_unit_test/entity"
	"go_unit_test/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var categoryRepository = repository.CategoryRepositoryMock{Mock: mock.Mock{}}
var categoryService = CategoryService{Repository: &categoryRepository}

func TestCategoryService_GetNotFound(t *testing.T) {

	// Program Mock
	categoryRepository.Mock.On("FindById", "1").Return(nil)

	category, err := categoryService.Get("1")
	assert.Nil(t, category)
	assert.NotNil(t, err)
}

func TestCategoryService_GetSuccess(t *testing.T) {

	category := entity.Category{
		Id:   "3",
		Name: "Laptop",
	}

	categoryRepository.Mock.On("FindById", "3").Return(category)

	result, err := categoryService.Get("3")
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, category.Id, result.Id)
	assert.Equal(t, category.Name, result.Name)
}
