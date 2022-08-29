package models

import (
	"api/config"
	"errors"
	"log"

	"gorm.io/gorm"
)

var users []User
var user User

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

func (us *User) Get(id uint64) (User, error) {
	data := us.Table().First(&user, "id = ?", id)

	if data.Error != nil {
		errors.Is(data.Error, gorm.ErrRecordNotFound)
		return user, errors.New("not found")
	}

	return user, nil
}

func (us *User) GetAll() []User {
	all := us.Table().Find(&users)

	if all.RowsAffected == 0 {
		log.Default().Println("Read users returned with empty results")
	}

	return users
}
