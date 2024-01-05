package models

import (
	"errors"
	"fmt"

	"github.com/gy8534/go-event-booking/db"
	"github.com/gy8534/go-event-booking/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) Save() error {
	query := "INSERT INTO users(email, password) VALUES (?, ?)"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	hashedPass, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	exec, err := stmt.Exec(u.Email, hashedPass)
	if err != nil {
		return err
	}
	id, err := exec.LastInsertId()
	u.ID = id
	return err
}

func (u *User) ValidateCredentials() error {
	query := "SELECT id, password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, u.Email)

	var retrievedPass string
	err := row.Scan(&u.ID, &retrievedPass)
	if err != nil {
		fmt.Println(err)
		return errors.New("credentials invalid")
	}

	if ok := utils.ValidatePass(retrievedPass, u.Password); !ok {
		return errors.New("credentials invalid")
	}

	return nil
}
