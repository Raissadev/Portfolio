package models

import (
	"api/config"
	"encoding/json"
	"errors"
	"log"
	"time"

	"gorm.io/gorm"
)

var dataSource config.DataSource

type User struct {
	ID         uint64    `json:"id" gorm:"primaryKey"`
	Name       string    `json:"name"`
	Email      string    `json:"email" gorm:"unique"`
	updated_at time.Time `json:"updated_at" gorm:"autoCreateTime"`
	created_at time.Time `json:"created_at" gorm:"autoCreateTime"`
}

func (us *User) Table() *gorm.DB {
	return config.Instance.Table("users")
}

func (us *User) Get(id uint64) (User, error) {
	var user User

	data := us.Table().Find(&user, id)

	if data.Error != nil || data.RowsAffected == 0 {
		errors.Is(data.Error, gorm.ErrRecordNotFound)
		return user, errors.New("not found")
	}

	return user, nil
}

func (us *User) GetAll() []User {
	var users []User

	all := us.Table().Find(&users)

	if all.RowsAffected == 0 {
		log.Default().Println("Read users returned with empty results")
	}

	return users
}

func (us *User) Create(params *json.Decoder) (User, error) {
	var user User
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

func (us *User) Update(id uint64, params *json.Decoder) (User, error) {
	var user User

	data := us.Table().Find(&user, id)

	if data.Error != nil || data.RowsAffected == 0 {
		log.Default().Println("Not found")
		return user, errors.New("not found")
	}

	err := params.Decode(&user)

	if err != nil {
		return user, errors.New("error")
	}

	up := us.Table().Save(&user)

	if up.Error != nil || up.RowsAffected == 0 {
		errors.Is(up.Error, gorm.ErrInvalidData)
		return user, errors.New("error")
	}

	return user, nil
}

func (us *User) Delete(id uint64) (bool, error) {
	var user User

	data := us.Table().Find(&user, id)

	if data.RowsAffected == 0 {
		log.Default().Println("not found")
		return false, errors.New("not found")
	}

	del := us.Table().Delete(&user)

	if del.Error != nil {
		errors.Is(del.Error, gorm.ErrInvalidData)
		return false, errors.New("error")
	}

	return true, nil
}
