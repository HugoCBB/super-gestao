package routers

import (
	"github.com/HugoCBB/super-gestao/internal/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequest() {

	r := gin.Default()

	api := r.Group("api/v1")
	{
		user := api.Group("/user")
		{
			user.POST("/cadastrar", controllers.CadastrarUsuario)
			user.GET("/", controllers.Usuario)
		}

		task := api.Group("/task")
		{
			task.GET("/", controllers.Tasks)
			task.GET("/obter/:id", controllers.ObterMensagemPorId)
			task.POST("/adicionar", controllers.AdicionarMensagem)

		}
	}

	r.Run(":8080")

}
