package controllers

import (
	"net/http"

	"github.com/HugoCBB/super-gestao/internal/database"
	"github.com/HugoCBB/super-gestao/internal/models"
	"github.com/gin-gonic/gin"
)

func CadastrarUsuario(c *gin.Context) {
	var u models.User

	if err := c.ShouldBindBodyWithJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Erro ao cadastrar usuario",
		})
		return
	}

	database.DB.Create(&u)
	c.JSON(http.StatusOK, gin.H{
		"status": "Usuario cadastrado com sucesso",
	})

}

func Usuario(c *gin.Context) {
	var u models.User
	database.DB.Find(&u)
	c.JSON(http.StatusOK, &u)

}
