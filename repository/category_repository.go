package repository

import entity "github.com/TantowiPutra/Go-Unit-Test/category"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}
