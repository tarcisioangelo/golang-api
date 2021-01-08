package testes

import (
	ModelUser "api/models"
	ServiceAuth "api/services"
	ServiceUser "api/services"
	"api/util"
	"io/ioutil"

	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Login - Fazendo login
func Login(w http.ResponseWriter, r *http.Request) {

	body, _ := ioutil.ReadAll(r.Body)

	type Login struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var login Login

	json.Unmarshal(body, &login)

	// Buscando dados de login
	user, err := ServiceUser.Login(login.Email, login.Password)

	// Caso retorne um erro
	if err != nil {
		w.Write(util.MessageInfo("message", err.Error()))
		return
	}

	// Criando token com id do usuário
	token, _ := ServiceAuth.CreateToken(user.ID)

	// Envia o token como retorno
	w.Write(util.MessageInfo("token", token))
}

// List - Listando Usuários
func List(w http.ResponseWriter, r *http.Request) {
	users := ServiceUser.List(5)
	json.NewEncoder(w).Encode(&users)
}

// Find - Busca apenas um contato
func Find(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, _ := strconv.Atoi(params["id"])

	user := ServiceUser.Find(id)

	json.NewEncoder(w).Encode(&user)
}

// Update - atualiza
func Update(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)

	var user ModelUser.User

	json.Unmarshal(body, &user)

	user.ID = 1

	json.NewEncoder(w).Encode(user)
}

// Delete deleta um contato
func Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	json.NewEncoder(w).Encode(params)
}
