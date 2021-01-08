package services

import (
	"api/database"
	ModelUser "api/models"
	"crypto/md5"
	"errors"
)

// List Test MysQL
func List(limit int) []ModelUser.User {

	db := database.Connect()

	var user []ModelUser.User

	db.Limit(limit).Find(&user)

	return user
}

// Login Test MysQL
func Login(email string, password string) (ModelUser.User, error) {

	db := database.Connect()

	var user ModelUser.User

	hasher := md5.New()
	hasher.Write([]byte(password))

	result := db.Where("email = ?", email).Find(&user)

	if result.Error != nil {
		return user, errors.New("Ops, houve um erro")
	}

	if result.RowsAffected == 0 {
		return user, errors.New("Usuário não encontrado")
	}

	return user, nil
}

// Find Test MysQL
func Find(id int) ModelUser.User {

	db := database.Connect()

	var user ModelUser.User

	db.Find(&user, id)

	return user
}
