package services

import (
	"api/database"
	ModelUser "api/models"
	"api/util"
	"errors"
)

// List - Listar usuários
// Paginação
func List(limit int) []ModelUser.User {

	db := database.Connect()

	var user []ModelUser.User

	db.Limit(limit).Find(&user)

	return user
}

// Login - Faz Login e retorna token JWT
func Login(email string, password string) (string, error) {

	db := database.Connect()

	var user ModelUser.User

	hash := util.GetMD5Hash(password)

	result := db.Where(&ModelUser.User{Email: email, Password: hash}).Find(&user)

	if result.Error != nil {
		return "", errors.New("Ops, houve um erro")
	}

	if result.RowsAffected == 0 {
		return "", errors.New("Usuário não encontrado")
	}

	// Criando token com id do usuário
	token, _ := CreateToken(user.ID)

	return token, nil
}

// Create - Cria um novo usuário
func Create(name string, email string, password string) error {

	db := database.Connect()

	hash := util.GetMD5Hash(password)

	result := db.Create(&ModelUser.User{Name: name, Email: email, Password: hash, IsActive: 1})

	if result.Error != nil {
		return errors.New("Ops, houve um erro")
	}

	return nil
}

// Update - Atualiza nome e senha do usuário
func Update(id int, name string, password string) error {

	db := database.Connect()

	hash := util.GetMD5Hash(password)

	var user ModelUser.User

	result := db.Where("id_user = ?", id).Find(&user)

	if result.Error != nil {
		return errors.New("Ops, houve um erro")
	}

	if result.RowsAffected == 0 {
		return errors.New("Usuário não encontrado")
	}

	user.Name = name
	user.Password = hash

	db.Save(&user)

	return nil
}

// Find - Localiza usuário por ID
func Find(id int) (ModelUser.User, error) {

	db := database.Connect()

	var user ModelUser.User

	result := db.Find(&user, id)

	if result.Error != nil {
		return user, errors.New("Ops, houve um erro")
	}

	if result.RowsAffected == 0 {
		return user, errors.New("Usuário não encontrado")
	}

	return user, nil
}

// Delete - Exclui usuário logado
func Delete(id int) error {

	db := database.Connect()

	var user ModelUser.User

	result := db.Find(&user, id)

	if result.Error != nil {
		return errors.New("Ops, houve um erro")
	}

	if result.RowsAffected == 0 {
		return errors.New("Usuário não encontrado")
	}

	db.Delete(&user)

	return nil
}
