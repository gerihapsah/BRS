package model

import (
	"db"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func UserAll() ([]User, error) {
	var err error
	result := []User{}
	sql := "SELECT * FROM tb_user"
	err = db.SQL.Select(&result, sql)
	return result, standardizeError(err)
}

func UserById(id string) (User, error) {
	var err error
	result := User{}
	sql := "SELECT * FROM tb_user  WHERE id = ?"
	err = db.SQL.Get(&result, sql, id)
	return result, standardizeError(err)
}

func UserByEmail(email string) (User, error) {
	var err error
	result := User{}
	sql := "SELECT * FROM tb_user  WHERE email = ?"
	err = db.SQL.Get(&result, sql, email)
	return result, standardizeError(err)
}

func UserCreate(user User) error {
	var err error
	sql := "INSERT INTO tb_user (name, email, password) VALUES (?,?,?)"
	_, err = db.SQL.Exec(sql, user.Name, user.Email, user.Password)
	return err
}

func UserUpdate(id string, user User) error {
	var err error
	sql := "UPDATE tb_user SET name=?, email=?, password=? WHERE id=?"
	_, err = db.SQL.Exec(sql, user.Name, user.Email, user.Password, id)
	return err
}

func UserDelete(id string) error {
	var err error
	sql := "DELETE FROM tb_user WHERE id=?"
	_, err = db.SQL.Exec(sql, id)
	return err
}
