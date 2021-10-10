package controllers

import (
	"api/src/authentication"
	"api/src/database"
	"api/src/models"
	"api/src/repository"
	"api/src/utils"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// CreateUser - create a new user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	bodyReq, error := ioutil.ReadAll(r.Body)
	if error != nil {
		utils.Error(w, http.StatusUnprocessableEntity, error)
		return
	}

	var user models.User
	if error = json.Unmarshal(bodyReq, &user); error != nil {
		utils.Error(w, http.StatusBadRequest, error)
		return
	}
	if error = user.Prepare("cadastro"); error != nil {
		utils.Error(w, http.StatusBadRequest, error)
		return
	}
	db, error := database.Connect()
	if error != nil {
		utils.Error(w, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()

	userRepo := repository.NewUserRepo(db)
	user.ID, error = userRepo.Create(user)
	if error != nil {
		utils.Error(w, http.StatusInternalServerError, error)
		return
	}
	utils.JSON(w, http.StatusCreated, user)
}

// GetUsers - Get all users from database
func GetUsers(w http.ResponseWriter, r *http.Request) {
	nameOrNick := strings.ToLower(r.URL.Query().Get("user"))

	db, error := database.Connect()
	if error != nil {
		utils.Error(w, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()

	userRepo := repository.NewUserRepo(db)
	users, error := userRepo.Find(nameOrNick)
	if error != nil {
		utils.Error(w, http.StatusInternalServerError, error)
		return
	}
	utils.JSON(w, http.StatusOK, users)
}

// GetUser - get an user from database by id
func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userID, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		utils.Error(w, http.StatusBadRequest, err)
		return
	}

	// Verify userID params with userID from token
	userIDToken, err := authentication.GetUserID(r)
	if err != nil {
		utils.Error(w, http.StatusUnauthorized, err)
		return
	}
	if userIDToken != userID {
		utils.Error(w, http.StatusForbidden, errors.New("User unauthorized"))
		return
	}
	db, err := database.Connect()
	if err != nil {
		utils.Error(w, http.StatusInternalServerError, err)
	}
	defer db.Close()

	userRepo := repository.NewUserRepo(db)
	user, err := userRepo.FindById(userID)
	if err != nil {
		utils.Error(w, http.StatusNotFound, err)
		return
	}
	utils.JSON(w, http.StatusOK, user)
}

// DeleteUser - remove an user from database by id
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		utils.Error(w, http.StatusBadRequest, err)
		return
	}
	// Verify userID params with userID from token
	userIDToken, err := authentication.GetUserID(r)
	if err != nil {
		utils.Error(w, http.StatusUnauthorized, err)
		return
	}
	if userIDToken != userID {
		utils.Error(w, http.StatusForbidden, errors.New("User unauthorized"))
		return
	}
	db, err := database.Connect()
	if err != nil {
		utils.Error(w, http.StatusInternalServerError, err)
	}
	defer db.Close()

	userRepo := repository.NewUserRepo(db)
	if err := userRepo.Delete(userID); err != nil {
		utils.Error(w, http.StatusInternalServerError, err)
	}
	utils.JSON(w, http.StatusNoContent, nil)
}

// UpdateUser - Update an user from database by id
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		utils.Error(w, http.StatusBadRequest, err)
		return
	}
	// Verify userID params with userID from token
	userIDToken, err := authentication.GetUserID(r)
	if err != nil {
		utils.Error(w, http.StatusUnauthorized, err)
		return
	}
	if userIDToken != userID {
		utils.Error(w, http.StatusForbidden, errors.New("User unauthorized"))
		return
	}

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

	if err = user.Prepare("update"); err != nil {
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
	if err = userRepo.Update(userID, user); err != nil {
		utils.Error(w, http.StatusInternalServerError, error)
		return
	}

	utils.JSON(w, http.StatusNoContent, nil)
}

// FollowUser - user follow another user
func FollowUser(w http.ResponseWriter, r *http.Request) {
	// get id from user's token
	follower_id, err := authentication.GetUserID(r)
	if err != nil {
		utils.Error(w, http.StatusUnauthorized, err)
		return
	}

	params := mux.Vars(r)
	user_id, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		utils.Error(w, http.StatusBadRequest, err)
		return
	}
	if user_id == follower_id {
		utils.Error(w, http.StatusForbidden, errors.New("Not possible follow yourself"))
		return
	}
	db, err := database.Connect()
	if err != nil {
		utils.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	userRepo := repository.NewUserRepo(db)
	if err := userRepo.FollowUser(follower_id, user_id); err != nil {
		utils.Error(w, http.StatusInternalServerError, err)
		return
	}
	utils.JSON(w, http.StatusNoContent, nil)
}

// UnFollowUser - Unfollow user by id
func UnFollowUser(w http.ResponseWriter, r *http.Request) {
	// get id from user's token
	follower_id, err := authentication.GetUserID(r)
	if err != nil {
		utils.Error(w, http.StatusUnauthorized, err)
		return
	}

	params := mux.Vars(r)
	user_id, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		utils.Error(w, http.StatusBadRequest, err)
		return
	}
	if user_id == follower_id {
		utils.Error(w, http.StatusForbidden, errors.New("Not possible unfollow yourself"))
		return
	}
	db, err := database.Connect()
	if err != nil {
		utils.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	userRepo := repository.NewUserRepo(db)
	if err := userRepo.UnFollowUser(follower_id, user_id); err != nil {
		utils.Error(w, http.StatusInternalServerError, err)
		return
	}
	utils.JSON(w, http.StatusNoContent, nil)
}

// GetFollowers - get all followers from an user
func GetFollowers(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userID, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		utils.Error(w, http.StatusBadRequest, err)
		return
	}
	db, err := database.Connect()
	if err != nil {
		utils.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	userRepo := repository.NewUserRepo(db)
	followers, err := userRepo.GetFollowers(userID)
	if err != nil {
		utils.Error(w, http.StatusInternalServerError, err)
		return
	}
	utils.JSON(w, http.StatusOK, followers)
}

// GetFollowing - get all users that user is following
func GetFollowing(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userID, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		utils.Error(w, http.StatusBadRequest, err)
		return
	}
	db, err := database.Connect()
	if err != nil {
		utils.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	userRepo := repository.NewUserRepo(db)
	followers, err := userRepo.GetFollowing(userID)
	if err != nil {
		utils.Error(w, http.StatusInternalServerError, err)
		return
	}
	utils.JSON(w, http.StatusOK, followers)
}
