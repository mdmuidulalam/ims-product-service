package main

import (
	"log"
	"os"
	"product-service/config"
	"runtime/debug"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag/example/basic/docs"
)

func main() {
	config := &config.Config{}

	configInstance := *config.GetConfig()

	r := gin.Default()

	if configInstance.Environment.Type != "local" {
		r.Use(logger())
	}

	setupApiDoc(r, configInstance.Environment.Port)

	r.Run(":" + configInstance.Environment.Port)
}

func setupApiDoc(r *gin.Engine, port string) {
	docs.SwaggerInfo.Title = "Musk Daily API Documentation"
	docs.SwaggerInfo.Description = "Simple API descriptions for musk daily API"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:" + port
	docs.SwaggerInfo.BasePath = ""
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		file, err := os.OpenFile("panic.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			log.Fatal(err)
		}
		log.SetOutput(file)

		defer func() {
			if err := recover(); err != nil {
				log.Println(err)
				log.Println("stacktrace from panic: \n" + string(debug.Stack()))
				c.Writer.WriteHeader(500)
			}
		}()

		c.Next()
	}
}
