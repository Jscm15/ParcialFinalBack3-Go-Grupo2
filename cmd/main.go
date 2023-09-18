package main

import (
<<<<<<< HEAD
	"github.com/Jscm15/ParcialFinalBack3-Go-Grupo2/cmd/server/external/database"
	"github.com/Jscm15/ParcialFinalBack3-Go-Grupo2/cmd/server/handler"
	"github.com/Jscm15/ParcialFinalBack3-Go-Grupo2/internal/appoiments"
=======
	"net/http"
	"github.com/Jscm15/ParcialFinalBack3-Go-Grupo2/cmd/server/database"
	"github.com/Jscm15/ParcialFinalBack3-Go-Grupo2/cmd/server/handler"
	"github.com/Jscm15/ParcialFinalBack3-Go-Grupo2/internal/patients"
>>>>>>> main
	"github.com/gin-gonic/gin"
	"net/http"

)

func main()  {
	router := gin.New()

	router.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"ok": "ok"})
	})


	mysqlDatabase, err := database.NewMySQLDatabase("localhost", "3306", "root", "root", "mydb")

	if err != nil {
		panic(err)
	}
	myDatabase := database.NewDatabase(mysqlDatabase)

<<<<<<< HEAD
	myDatabase:=database.NewDatabase(mysqlDatabase)

=======
	
>>>>>>> main
	appoimentsService:=appoiments.NewService(myDatabase)

	appoimentsHandler:= handler.NewAppoimentsHandler(appoimentsService,appoimentsService,appoimentsService)

	appoimentsGroup:=router.Group("/appoiments")

	appoimentsGroup.GET("/:id",appoimentsHandler.GetAppoimentByID)
	appoimentsGroup.GET("/dni/:dni",appoimentsHandler.GetAppoimentByPatient)
	appoimentsGroup.POST("",appoimentsHandler.CreateAppoiment)
	appoimentsGroup.PUT("/:id",appoimentsHandler.PutAppoiment)
	appoimentsGroup.DELETE("/:id",appoimentsHandler.DeleteAppoiment)
	

  patientsService := patients.NewService(myDatabase)
  patientsHandler := handler.NewPatientHandler(patientsService, patientsService, patientsService)
  patientsGroup := router.Group("/patients")
  patientsGroup.GET("/:id", patientsHandler.GetPatientByID)
  patientsGroup.POST("", patientsHandler.CreatePatient)
  patientsGroup.PUT("/:id", patientsHandler.PutPatient)
  patientsGroup.DELETE("/:id", patientsHandler.DeletePatient)

	err = router.Run()

	if err != nil {
		panic(err)
	}
}

