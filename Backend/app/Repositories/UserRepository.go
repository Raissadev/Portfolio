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
