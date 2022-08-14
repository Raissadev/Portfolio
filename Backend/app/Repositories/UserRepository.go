package repositories

import (
	. "api/app/Models"
	"fmt"
)

type UserRepository struct {
	Model User
	RepositoryInterface
}

func (r *UserRepository) All() string {
	fmt.Println(r.Model.GetAll())
	return "asas"
}
