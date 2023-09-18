package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"github.com/Jscm15/ParcialFinalBack3-Go-Grupo2/internal/appoiments"
	"github.com/gin-gonic/gin"
)

type AppoimentsGetter interface {
	GetAppointmentByID(id int) (appoiments.Appoiment,error)
	GetAppointmentByDni(dni int) (appoiments.Appoiment,error)
}

type AppoimentsCreator interface {
	ModifyAppointment(id int, appoiment appoiments.Appoiment) (appoiments.Appoiment, error)
	UpdateDate(id int, appoiment appoiments.Appoiment) (appoiments.Appoiment, error)
	CreateAppointment(appoiment appoiments.Appoiment) (appoiments.Appoiment, error)
}

type AppoimentsDelete interface {
	DeleteAppointment(id int) error
}

type AppoimentsHandler struct {
	appoimentsGetter AppoimentsGetter
	appoimentsCreator AppoimentsCreator
	appoimentsDelete AppoimentsDelete
}

func NewAppoimentsHandler(getter AppoimentsGetter,creator AppoimentsCreator, deleter AppoimentsDelete) *AppoimentsHandler  {
	return &AppoimentsHandler{appoimentsGetter: getter,
		appoimentsCreator: creator,
		appoimentsDelete:  deleter}
}

func (ah *AppoimentsHandler) GetAppointmentByID( ctx *gin.Context)  {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid appoiment id"})
		return
	}
	appoiment, err := ah.appoimentsGetter.GetAppointmentByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
		return
	}
	ctx.JSON(http.StatusOK, appoiment)
}

func (ah *AppoimentsHandler) GetAppointmentByDni( ctx *gin.Context)  {
	dniParam := ctx.Param("dni")
	dni, err := strconv.Atoi(dniParam)
	if err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid appoiment id"})
		return
	}
	appoiment, err := ah.appoimentsGetter.GetAppointmentByDni(dni)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
		return
	}
	ctx.JSON(http.StatusOK, appoiment)
}

func (ah *AppoimentsHandler) CreateAppoiment(ctx *gin.Context) {
	appoimentRequest := appoiments.Appoiment{}
	err := ctx.BindJSON(&appoimentRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	appoiment, err := ah.appoimentsCreator.CreateAppointment(appoimentRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}
	
	ctx.JSON(http.StatusOK, appoiment)
}

func (ah *AppoimentsHandler) ModifyAppointment(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid appoiment id"})
		return
	}
	appoimentRequest := appoiments.Appoiment{}
	err = ctx.BindJSON(&appoimentRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	appoiment, err := ah.appoimentsCreator.ModifyAppointment(id, appoimentRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}
	ctx.JSON(http.StatusOK, appoiment)
}

func (ah *AppoimentsHandler) UpdateDate(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid appoiment id"})
		return
	}
	appoimentRequest := appoiments.Appoiment{}
	err = ctx.BindJSON(&appoimentRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if appoimentRequest.DateAndHour == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
		}
	appoiment, err := ah.appoimentsCreator.UpdateDate(id,appoimentRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}
	ctx.JSON(http.StatusOK, appoiment)
}

func (ah *AppoimentsHandler) DeleteAppoiment(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"),10, 64)
	if err != nil {
	ctx.JSON(http.StatusBadRequest, gin.H{ "error": "invalid ID"})
	return
	}
	err = ah.appoimentsDelete.DeleteAppointment(int(id))
	if err != nil {
	ctx.JSON(http.StatusBadRequest, gin.H{ "error": err.Error() })
	return
	}
	ctx.JSON(http.StatusOK, gin.H{ "data": fmt.Sprintf("Appoiment %d deleted", id) })
	}