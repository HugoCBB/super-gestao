package controllers

import (
	"net/http"

	"github.com/HugoCBB/super-gestao/internal/database"
	"github.com/HugoCBB/super-gestao/internal/models"
	"github.com/gin-gonic/gin"
)

func AdicionarMensagem(c *gin.Context) {
	var t models.Task

	if err := c.ShouldBindBodyWithJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Erro ao adicionar mensagem usuario",
		})
		return
	}

	database.DB.Create(&t)
	c.JSON(http.StatusOK, gin.H{
		"status": "Mensagem adicionada com sucesso",
	})
}
