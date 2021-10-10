package models

import (
	"api/src/hash"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

// User - user model
type User struct {
	ID       uint64    `json:"id,omitempty"`
	Name     string    `json:"name,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Password string    `json:"password,omitempty"`
	CreateAt time.Time `json:"CreateAt,omitempty"`
}

func (user *User) Prepare(step string) error {
	if err := user.validate(step); err != nil {
		return err
	}
	if err := user.format(step); err != nil {
		return err
	}
	return nil
}

func (user *User) validate(step string) error {
	if user.Name == "" {
		return errors.New("Name: invalid arguments")
	}
	if user.Nick == "" {
		return errors.New("Nick: invalid arguments")
	}
	if user.Email == "" {
		return errors.New("Email: invalid arguments")
	}
	if err := checkmail.ValidateFormat(user.Email); err != nil {
		return errors.New("Email invalid")
	}
	if user.Password == "" && step == "cadastro" {
		return errors.New("Password: invalid arguments")
	}
	return nil
}

func (user *User) format(step string) error {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)
	user.Password = strings.TrimSpace(user.Password)

	if step == "cadastro" {
		passwordHash, err := hash.Hash(user.Password)
		if err != nil {
			return err
		}
		user.Password = string(passwordHash)
	}
	return nil
}
