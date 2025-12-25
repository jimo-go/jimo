package models

import (
	"strings"

	"github.com/jimo-go/framework/auth"
	"github.com/jimo-go/framework/database"
)

type User struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	PasswordHash string `json:"-"`
}

func (User) TableName() string { return "users" }

func Users() *database.Record[User] {
	return database.Model[User]()
}

func SeedUsers() error {
	pass, err := auth.HashPassword("password")
	if err != nil {
		return err
	}

	items := []User{
		{ID: 1, Name: "Amina", Email: "amina@example.com", PasswordHash: pass},
		{ID: 2, Name: "Marko", Email: "marko@example.com", PasswordHash: pass},
		{ID: 3, Name: "Lejla", Email: "lejla@example.com", PasswordHash: pass},
	}

	m := Users()
	for i := range items {
		u := items[i]
		if err := m.Create(&u); err != nil {
			return err
		}
	}
	return nil
}

func FindUserByEmail(email string) (User, bool, error) {
	email = strings.TrimSpace(strings.ToLower(email))
	if email == "" {
		return User{}, false, nil
	}

	items, err := Users().All()
	if err != nil {
		return User{}, false, err
	}
	for _, u := range items {
		if strings.ToLower(u.Email) == email {
			return u, true, nil
		}
	}
	return User{}, false, nil
}
