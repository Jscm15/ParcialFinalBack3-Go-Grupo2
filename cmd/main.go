// @title Ejemplo de API Swagger en Go
// @version 1.0
// @description Esta es una API Swagger de ejemplo en Go.
// @BasePath /api/v1
package main

import (
    "github.com/gin-gonic/gin"
    "github.com/swaggo/gin-swagger"
    "github.com/swaggo/gin-swagger/swaggerFiles"
)

func main() {
    r := gin.Default()

    // Rutas y controladores aquí

    // Ruta para la documentación Swagger
    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    r.Run(":8080")
}
