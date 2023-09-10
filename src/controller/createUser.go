package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/wsldev/meu-primeiro-crud-go/src/configuration/rest_err"
	"github.com/wsldev/meu-primeiro-crud-go/src/controller/model/request"
)

func CreateUser(c *gin.Context) {

	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := rest_err.NewBadRequestError(
			fmt.Sprintf("There are some incorrect fields, error: %s", err.Error()))
		c.JSON(restErr.Code, restErr)
	}

	fmt.Println(userRequest)
}
