package models

import (
	"log"
	"time"
)

type User struct {
	ID        int
	UUID      string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
}

func (u *User) CreateUser() (err error) {
	cmdU := `INSERT INTO users (
		uuid,
		name,
		email,
		password,
		created_at) values (?, ?, ?, ?, ?)`

	_, err = Db.Exec(cmdU, createUUID(), u.Name, u.Email, Encrypt(u.Password), time.Now())

	if err != nil {
		log.Fatalln(err)
	}

	return err
}
