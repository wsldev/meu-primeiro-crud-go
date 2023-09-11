package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wsldev/meu-primeiro-crud-go/src/configuration/validation"
	"github.com/wsldev/meu-primeiro-crud-go/src/controller/model/request"
	"github.com/wsldev/meu-primeiro-crud-go/src/controller/model/response"
)

func CreateUser(c *gin.Context) {

	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		log.Printf("Error trying to marshal object, error=%s", err.Error())
		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	fmt.Println(userRequest)
	response := response.UserResponse{
		Email: userRequest.Email,
		ID:    "test",
		Name:  userRequest.Name,
		Age:   userRequest.Age,
	}

	c.JSON(http.StatusCreated, response)
}
