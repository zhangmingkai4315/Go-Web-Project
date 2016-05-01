package controllers

import (
	"encoding/json"
	"net/http"
	"restful-api/common"
	"restful-api/data"
	"restful-api/models"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var dataResource UserResource
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(w, err, "Invaild user data", 500)
		return
	}
	user := &dataResource.Data
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("users")
	repo := &data.UserRepository{c}
	repo.CreateUser(user)
	user.HashPassword = string("")
	if j, err := json.Marshal(UserResource{Data: *user}); err != nil {
		common.DisplayAppError(w, err, "An unexpected has occurred", 500)
		return
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(j)
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	var dataResource LoginResource
	var token string
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(w, err, "Invaild login info", 500)
		return
	}

	loginModel := dataResource.Data
	loginUser := models.User{
		Email:    loginModel.Email,
		Password: loginModel.Password,
	}
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("users")
	repo := &data.UserRepository{c}

	user, err := repo.Login(loginUser)
	if err != nil {
		common.DisplayAppError(w, err, "Invaild login info credentials", 401)
		return
	} else {
		token, err = common.GenerateJWT(user.Email, "member")
		if err != nil {
			common.DisplayAppError(w, err, "Error while generagter token", 500)
			return
		}
	}
	w.Header().Set("Content-Type", "application/json")
	user.HashPassword = string("")
	authUser := AuthUserModel{
		User:  user,
		Token: token,
	}
	j, err := json.Marshal(AuthUserResource{Data: authUser})
	if err != nil {
		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(j)

}
