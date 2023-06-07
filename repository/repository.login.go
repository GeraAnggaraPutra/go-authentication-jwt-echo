package repository

import (
	"database/sql"
	"fmt"
	"go-auth-jwt/entity"
	"go-auth-jwt/helpers"
)

type RepositoryLogin interface {
	CheckLogin(username, password string) (bool, error)
}

type repositoryLogin struct {
	db *sql.DB
}

func NewRepositoryLogin(db *sql.DB) *repositoryLogin {
	return &repositoryLogin{db}
}

func (r *repositoryLogin) CheckLogin(username, password string) (bool, error) {
	var pwd string
	users := &entity.Users{}
	sqlQuery := "SELECT * FROM users WHERE username = ?"
	err := r.db.QueryRow(sqlQuery, username).Scan(&users.ID, &users.USERNAME, &pwd)
	if err == sql.ErrNoRows {
		fmt.Println("Username not found")
		return false, err
	}

	if err != nil {
		fmt.Println("Query error")
		return false, err
	}

	match, err := helpers.CheckPasswordHash(password, pwd)
	if !match {
		fmt.Println("Hash and password doesn't match.")
		return false, err
	}

	return true, nil
}
