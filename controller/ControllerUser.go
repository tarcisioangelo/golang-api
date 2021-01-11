package testes

import (
	ServiceUser "api/services"
	"api/util"
	"io/ioutil"

	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
)

// Create - Criando usuário
func Create(w http.ResponseWriter, r *http.Request) {

	body, _ := ioutil.ReadAll(r.Body)

	type NewUser struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var user NewUser

	json.Unmarshal(body, &user)

	// Salvando dados
	err := ServiceUser.Create(user.Name, user.Email, user.Password)

	// Caso retorne um erro
	if err != nil {
		w.Write(util.MessageInfo("message", err.Error()))
		return
	}

	// Envia o token como retorno
	w.Write(util.MessageInfo("message", "Cadastro realizado com sucesso"))
}

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
	token, err := ServiceUser.Login(login.Email, login.Password)

	// Caso retorne um erro
	if err != nil {
		w.Write(util.MessageInfo("message", err.Error()))
		return
	}

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

	id, err := strconv.Atoi(params["id"])

	user, err := ServiceUser.Find(id)

	if err != nil {
		w.Write(util.MessageInfo("message", err.Error()))
		return
	}

	json.NewEncoder(w).Encode(&user)
}

// Update - Atualizando usuário
func Update(w http.ResponseWriter, r *http.Request) {

	body, _ := ioutil.ReadAll(r.Body)

	userID := context.Get(r, "userID").(int)

	type UpdateUser struct {
		Name     string `json:"name"`
		Password string `json:"password"`
	}

	var user UpdateUser

	json.Unmarshal(body, &user)

	// Salvando dados
	err := ServiceUser.Update(userID, user.Name, user.Password)

	// Caso retorne um erro
	if err != nil {
		w.Write(util.MessageInfo("message", err.Error()))
		return
	}

	// Envia o token como retorno
	w.Write(util.MessageInfo("message", "Cadastro atualizado com sucesso"))
}

// Delete deleta um contato
func Delete(w http.ResponseWriter, r *http.Request) {
	userID := context.Get(r, "userID").(int)

	// Excluindo usuário logado
	err := ServiceUser.Delete(userID)

	if err != nil {
		w.Write(util.MessageInfo("message", err.Error()))
		return
	}

	w.Write(util.MessageInfo("message", "Excluído com sucesso"))
}
