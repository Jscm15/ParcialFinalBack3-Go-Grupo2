package main

import (
	"net/http"
	"os"

	"github.com/Jscm15/ParcialFinalBack3-Go-Grupo2/cmd/server/config"
	"github.com/Jscm15/ParcialFinalBack3-Go-Grupo2/cmd/server/database"
	"github.com/Jscm15/ParcialFinalBack3-Go-Grupo2/cmd/server/handler"
	"github.com/Jscm15/ParcialFinalBack3-Go-Grupo2/cmd/server/middlewares"
	"github.com/Jscm15/ParcialFinalBack3-Go-Grupo2/internal/appoiments"
	"github.com/Jscm15/ParcialFinalBack3-Go-Grupo2/internal/dentists"
	"github.com/Jscm15/ParcialFinalBack3-Go-Grupo2/internal/patients"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)


func main() {

	env := os.Getenv("ENV")
	if env == "" {
		env = "local"
	}

	cfg, err := config.NewConfig(env)

	if err != nil {
		panic(err)
	}

	authMidd := middlewares.NewAuth(cfg.PublicConfig.PublicKey, cfg.PrivateConfig.SecretKey)


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

	appoimentsGroup.GET("/:id", authMidd.AuthHeader, appoimentsHandler.GetAppointmentByID)
	appoimentsGroup.GET("/dni/:dni", authMidd.AuthHeader, appoimentsHandler.GetAppointmentByDni)
	appoimentsGroup.POST("", authMidd.AuthHeader, appoimentsHandler.CreateAppoiment)
	appoimentsGroup.PUT("/:id", authMidd.AuthHeader, appoimentsHandler.ModifyAppointment)
	appoimentsGroup.DELETE("/:id", authMidd.AuthHeader, appoimentsHandler.DeleteAppoiment)

	patientsService := patients.NewService(myDatabase)

	patientsHandler := handler.NewPatientHandler(patientsService, patientsService, patientsService)

	patientsGroup := router.Group("/patients")

	patientsGroup.GET("/:id", authMidd.AuthHeader, patientsHandler.GetPatientByID)
	patientsGroup.POST("", authMidd.AuthHeader, patientsHandler.AddPatient)
	patientsGroup.PUT("/:id", authMidd.AuthHeader, patientsHandler.ModifyPatientByID)
	patientsGroup.DELETE("/:id", authMidd.AuthHeader, patientsHandler.DeletePatientByID)
	
	dentistsService := dentists.NewService(myDatabase)
	
	dentistsHandler := handler.NewDentistsHandler(dentistsService, dentistsService, dentistsService)
	
	dentistGroup := router.Group("/dentists")
	dentistGroup.POST("", authMidd.AuthHeader, dentistsHandler.CreateDentist)
	dentistGroup.GET("/:id", authMidd.AuthHeader, dentistsHandler.GetDentistByID)
	dentistGroup.PUT("/:id", authMidd.AuthHeader, dentistsHandler.UpdateDentistByID)
	dentistGroup.GET("/matricula/:matricula", authMidd.AuthHeader, dentistsHandler.GetDentistByMatricula)
	dentistGroup.DELETE("/:id", authMidd.AuthHeader, dentistsHandler.DeleteDentistByID)

	

	err = router.Run()

	if err != nil {
		panic(err)
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    router.Run(":8080")

}
