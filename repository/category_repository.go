package repository

import "belajar-golang-test-unit/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}
