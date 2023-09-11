package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wsldev/meu-primeiro-crud-go/src/configuration/logger"
	"github.com/wsldev/meu-primeiro-crud-go/src/configuration/validation"
	"github.com/wsldev/meu-primeiro-crud-go/src/controller/model/request"
	"github.com/wsldev/meu-primeiro-crud-go/src/controller/model/response"
	"go.uber.org/zap"
)

func CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser controller",
		zap.String("journey", "createUser"))

	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err,
			zap.String("journey", "createUser"))

		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	response := response.UserResponse{
		Email: userRequest.Email,
		ID:    "test",
		Name:  userRequest.Name,
		Age:   userRequest.Age,
	}

	logger.Info("User created successfully",
		zap.String("journey", "createUser"))

	c.JSON(http.StatusCreated, response)
}
