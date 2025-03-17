package handlers

import (
	"net/http"
	"user-service/http/requests"
	"user-service/models"
	"user-service/repositories"
	"user-service/utils"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {

	var registerRequest requests.CreateUserRequest
	if err := c.ShouldBind(&registerRequest); err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	name := registerRequest.Name
	email := registerRequest.Email

	if _, err := repositories.GetUserByEmail(email); err == nil {
		utils.SendErrorResponse(c, http.StatusConflict, "User already exists")
		return
	}

	user := models.User{
		Name:  name,
		Email: email,
	}
	repositories.CreateUser(&user)

	utils.SendSuccessResponse(c, http.StatusCreated, gin.H{"message": "User created", "user": user})
}

func GetUsers(c *gin.Context) {
	users, err := repositories.GetUsers()
	if err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.SendSuccessResponse(c, http.StatusOK, gin.H{"users": users})
}
