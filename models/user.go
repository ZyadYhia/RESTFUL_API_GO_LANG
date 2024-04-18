package models

import (
	"errors"
	"example.com/rest-api/db"
	"example.com/rest-api/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) Save() error {
	query := "INSERT INTO users(email, password) VALUES (?,?)"
	// to make sure no SQL Injection
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close() // to close whenever function ends
	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	//Exec() used to change or add row in DB
	result, err := stmt.Exec(u.Email, hashedPassword)
	if err != nil {
		return err
	}
	userId, err := result.LastInsertId()
	if err != nil {
		return err
	}
	u.ID = userId
	return nil
}
func (u *User) ValidateCredentials() error {
	query := "SELECT id, password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, u.Email)
	var retrievePassword string
	err := row.Scan(&u.ID, &retrievePassword)
	if err != nil {
		return err
	}
	passwordIsValid := utils.CheckPasswordHash(u.Password, retrievePassword)

	if !passwordIsValid {
		return errors.New("Credentials invalid")
	}
	return nil
}
