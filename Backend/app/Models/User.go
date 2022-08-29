package models

import (
	"api/config"
	"log"

	"gorm.io/gorm"
)

type User struct {
	ID         uint64 `json:"id" gorm:"primaryKey"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Updated_at string `json:"updated_at"`
	Created_at string `json:"created_at"`
}

func (us *User) Table() *gorm.DB {
	return config.Instance.Table("users")
}

func (us *User) GetAll() []User {
	var users []User

	all := us.Table().Find(&users)

	if all.RowsAffected == 0 {
		log.Default().Println("Read users returned with empty results")
	}

	return users
}
