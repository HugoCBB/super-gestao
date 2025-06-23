package controllers

import (
	"net/http"

	"github.com/HugoCBB/super-gestao/internal/database"
	"github.com/HugoCBB/super-gestao/internal/models"
	"github.com/gin-gonic/gin"
)

func Tasks(c *gin.Context) {
	// 	Obter todas as tasks
	var t []models.Task
	database.DB.Find(&t)
	c.JSON(http.StatusOK, gin.H{
		"tarefa": t,
	})
}

func AdicionarMensagem(c *gin.Context) {
	var t models.Task

	if err := c.ShouldBindBodyWithJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Erro ao adicionar mensagem",
		})
		return
	}

	database.DB.Create(&t)
	c.JSON(http.StatusOK, gin.H{
		"status": "Mensagem adicionada com sucesso",
	})
}

func ObterMensagemPorId(c *gin.Context) {
	// 	Obter task
	var t models.Task
	id := c.Param("id")

	database.DB.First(&t, id)
	if t.Id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Task nao encontrada",
		})
		return
	}

	// 	Enviar task
	c.JSON(http.StatusOK, gin.H{
		"task": t,
	})

}
