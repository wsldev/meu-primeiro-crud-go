package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wsldev/meu-primeiro-crud-go/src/configuration/rest_err"
	"github.com/wsldev/meu-primeiro-crud-go/src/controller/model/request"
)

func CreateUser(c *gin.Context) {

	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := rest_err.NewBadRequestError(
			"There are some incorrect fields")
		c.JSON(restErr.Code, restErr)
		fmt.Println(err.Error())
		return
	}

	c.JSON(http.StatusCreated, userRequest)
}
