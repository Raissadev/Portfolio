package repositories

import (
	. "api/app/Models"
	"encoding/json"
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

func (r *UserRepository) Create(params *json.Decoder) (User, error) {
	return r.Model.Create(params)
}

func (r *UserRepository) Update(id uint64, params *json.Decoder) (User, error) {
	return r.Model.Update(id, params)
}

func (r *UserRepository) Delete(id uint64) (bool, error) {
	return r.Model.Delete(id)
}
