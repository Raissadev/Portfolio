package models

import (
	"api/config"
	"encoding/json"
	"errors"
	"log"
	"time"

	"gorm.io/gorm"
)

var users []User
var user User
var dataSource config.DataSource

type User struct {
	ID         uint64    `json:"id" gorm:"primaryKey"`
	Name       string    `json:"name"`
	Email      string    `json:"email" gorm:"uniqueIndex"`
	updated_at time.Time `json:"updated_at" gorm:"autoCreateTime"`
	created_at time.Time `json:"created_at" gorm:"autoCreateTime"`
}

func (us *User) Table() *gorm.DB {
	dataSource.Open()
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

func (us *User) Create(params *json.Decoder) (User, error) {
	err := params.Decode(&user)

	if err != nil {
		return user, errors.New("error")
	}

	data := us.Table().Save(&user)

	if data.Error != nil {
		errors.Is(data.Error, gorm.ErrInvalidData)
		return user, errors.New("error")
	}

	return user, nil
}
