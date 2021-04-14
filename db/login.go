package db

import (
	"errors"
	"log"

	"github.com/natajonasdacoliveira/credisis-unidades-federativas/model"
	"golang.org/x/crypto/bcrypt"
)

func Login(loginInput model.User) (uint64, error) {
	db := InitDb()
	users:= []model.User{}
	user := model.User{}
	password := []byte(loginInput.Password)

	err := db.Select(&users, "SELECT * FROM user WHERE email = ?", loginInput.Email)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	if(len(users) == 0) {
		err = errors.New("credenciais inválidas")
		return 0, err
	}

	user = users[0]
	
	err = bcrypt.CompareHashAndPassword( []byte(user.Password), password)
	if err != nil {
		err = errors.New("credenciais inválidas")
		return 0, err
	}

	return user.ID, nil
}
