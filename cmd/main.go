package main

import (
	"net/http"
	"github.com/Jscm15/ParcialFinalBack3-Go-Grupo2/cmd/server/database"
	"github.com/Jscm15/ParcialFinalBack3-Go-Grupo2/cmd/server/handler"
	"github.com/Jscm15/ParcialFinalBack3-Go-Grupo2/internal/appoiments"
	"github.com/Jscm15/ParcialFinalBack3-Go-Grupo2/internal/patients"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()

	router.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"ok": "ok"})
	})

	mysqlDatabase, err := database.NewMySQLDatabase("localhost", "3306", "root", "root", "mydb")

	if err != nil {
		panic(err)
	}
	myDatabase := database.NewDatabase(mysqlDatabase)

	appoimentsService := appoiments.NewService(myDatabase)

	appoimentsHandler := handler.NewAppoimentsHandler(appoimentsService, appoimentsService, appoimentsService)

	appoimentsGroup := router.Group("/appoiments")

	appoimentsGroup.GET("/:id", appoimentsHandler.GetAppointmentByID)
	appoimentsGroup.GET("/dni/:dni", appoimentsHandler.GetAppointmentByDni)
	appoimentsGroup.POST("", appoimentsHandler.CreateAppoiment)
	appoimentsGroup.PUT("/:id", appoimentsHandler.ModifyAppointment)
	appoimentsGroup.DELETE("/:id", appoimentsHandler.DeleteAppoiment)

	patientsService := patients.NewService(myDatabase)

	patientsHandler := handler.NewPatientHandler(patientsService, patientsService, patientsService)

	patientsGroup := router.Group("/patients")

	patientsGroup.GET("/:id", patientsHandler.GetPatientByID)
	patientsGroup.POST("", patientsHandler.AddPatient)
	patientsGroup.PUT("/:id", patientsHandler.ModifyPatientByID)
	patientsGroup.DELETE("/:id", patientsHandler.DeletePatientByID)
	
        patientsGroup := router.Group("/patients")
	
	dentistRepository := database.NewDentistRepository(db)
	
	dentistsHandler := handler.NewDentistsHandler(dentistsService)
	dentistGroup := router.Group("/dentists")
	dentistGroup.POST("", dentistsHandler.PostDentist)
	dentistGroup.GET("/:id", dentistsHandler.GetDentistByID)
	dentistGroup.PUT("/:id", dentistsHandler.PutDentist)
	dentistGroup.PATCH("/:id", dentistsHandler.PatchDentist)
	dentistGroup.DELETE("/:id", dentistsHandler.DeleteDentist)

	

	err = router.Run()

	if err != nil {
		panic(err)
	}
}
