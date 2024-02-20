package service

import (
	"errors"
	"go_unit_test/entity"
	"go_unit_test/repository"
)

type CategoryService struct {
	Repository repository.CategoryRepository
}

func (service CategoryService) Get(id string) (*entity.Category, error) {
	category := service.Repository.FindById(id)
	if category == nil {
		return category, errors.New("Category Tidak Ditemukan")
	} else {
		return category, nil
	}
}
