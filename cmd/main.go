package main

import (
	"fmt"
	"net/http"
	"github.com/Jscm15/ParcialFinalBack3-Go-Grupo2/cmd/server/external/database"
	"github.com/gin-gonic/gin"
)

func main()  {
	router := gin.New()

	router.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"ok": "ok"})
	})

	mysqlDatabase,err:=database.NewMySQLDatabase("localhost","3306","root","root","mydb")
	
	if err != nil {
		panic(err)
	}

	myDatabase:=database.NewDatabase(mysqlDatabase)
	
	/*
	appoimentsService:=appoiments.NewService(myDatabase)

	appoimentsHandler:= handler.NewAppoimentsHandler(appoimentsService,appoimentsService,appoimentsService)

	appoimentsGroup:=router.Group("/appoiments")

	appoimentsGroup.GET("/:id",appoimentsHandler.GetAppoimentByID)
	appoimentsGroup.GET("/dni/:dni",appoimentsHandler.GetAppoimentByPatient)
	appoimentsGroup.POST("",appoimentsHandler.CreateAppoiment)
	appoimentsGroup.PUT("/:id",appoimentsHandler.PutAppoiment)
	appoimentsGroup.DELETE("/:id",appoimentsHandler.DeleteAppoiment)
	*/
	
	err = router.Run()

	if err != nil {
		panic(err)
	}
}