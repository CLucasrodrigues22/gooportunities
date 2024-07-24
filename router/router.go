package router

import "github.com/gin-gonic/gin"

func InitializeRouter() {
	// Inicializando o Router utilizando as configurações Default do gin
	router := gin.Default()
	m := "Olá mundo"
	// Definindo uma Rota
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": m,
		})
	})
	// Estamos rodando a nossa api
	router.Run(":8000") // listen and serve on 0.0.0.0:8080
}
