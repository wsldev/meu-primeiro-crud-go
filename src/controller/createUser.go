package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wsldev/meu-primeiro-crud-go/src/configuration/rest_err"
)

func CreateUser(c *gin.Context) {
	err := rest_err.NewBadRequestError("Você chamou a rota errada")

	c.JSON(err.Code, err)
}
