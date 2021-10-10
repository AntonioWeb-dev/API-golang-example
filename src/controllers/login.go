package controllers

import (
	"api/src/authentication"
	"api/src/database"
	"api/src/hash"
	"api/src/models"
	"api/src/repository"
	"api/src/utils"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Login - Make the users's authentication
func Login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.Error(w, http.StatusUnprocessableEntity, err)
		return
	}
	var user models.User
	if err = json.Unmarshal(body, &user); err != nil {
		utils.Error(w, http.StatusBadRequest, err)
		return
	}
	db, error := database.Connect()
	if error != nil {
		utils.Error(w, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()
	userRepo := repository.NewUserRepo(db)
	userFound, err := userRepo.FindByEmail(user.Email)
	if error != nil {
		utils.Error(w, http.StatusInternalServerError, error)
		return
	}
	if err = hash.Verify(user.Password, userFound.Password); err != nil {
		utils.Error(w, http.StatusUnauthorized, err)
		return
	}
	token, err := authentication.Token(userFound.ID)
	if err != nil {
		utils.Error(w, http.StatusInternalServerError, err)
		return
	}
	w.Write([]byte(token))
}
