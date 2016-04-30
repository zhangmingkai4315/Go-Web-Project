package controllers

import "restful-api/models"

type (
	UserResource struct {
		Data models.User `json:"data"`
	}

	LoginResource struct {
		Data LoginModel `json:"data"`
	}

	AuthUserResource struct {
		Data AuthUserModel `json:"data"`
	}

	LoginModel struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	AuthUserModel struct {
		User  models.User `json:"user"`
		Token string      `json:"token"`
	}
)

type (
	TaskResource struct {
		Data models.Task `json:"data"`
	}
	TaskResource struct {
		Data []models.Task `json:"data"`
	}
)
