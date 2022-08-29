package repositories

import (
	. "api/app/Models"
)

type UserRepository struct {
	Model User
	RepositoryInterface
}

func (r *UserRepository) All() []User {
	return r.Model.GetAll()
}

func (r *UserRepository) Get(id uint64) (User, error) {
	return r.Model.Get(id)
}
