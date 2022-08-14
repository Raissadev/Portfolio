package models

import (
	"database/sql"
	"fmt"
)

type User struct {
	ID         uint64 `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Updated_at string `json:"updated_at"`
	Created_at string `json:"created_at"`
	Db         *sql.DB
}

var user User

func (us *User) Get(id uint64) (User, error) {
	db := us.Db
	query := `select title, content from users where id = $1`
	row, err := db.Query(query, id)

	if err != nil {
		return user, err
	}

	defer row.Close()

	if row.Next() {
		var name, email string

		err := row.Scan(&name, &email)
		if err != nil {
			return user, err
		}

		user = User{
			ID:    id,
			Name:  name,
			Email: email,
		}
	}

	return user, nil
}

func (us *User) GetAll() string {
	db := us.Db

	// var users []User

	rows, _ := db.Query(`select * from users;`)
	fmt.Println(rows, db)
	// if err != nil {
	// return users, err
	// }

	// defer rows.Close()

	// for rows.Next() {
	// var id uint64
	// var name, email string

	// err := rows.Scan(&id, &name, &email)

	// if err != nil {
	// return users, err
	// }

	// user := User{
	// ID:    id,
	// Name:  name,
	// Email: email,
	// }

	// users = append(users, user)
	// }

	return "helo"
	// return users, nil
}
