package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Jscm15/ParcialFinalBack3-Go-Grupo2/internal/dentists"
	"github.com/gin-gonic/gin"
)

type Getter interface {
	GetByID(id int) (dentists.Dentist, error)
}

type Creator interface {
	Create(dentista dentists.Dentist) (dentists.Dentist, error)
	UpdateDate(id int, dentista dentists.Dentist) (dentists.Dentist, error)
}

type Deleter interface {
	DeleteByID(id int) error
}
type DentistsHandler struct {
	Getter  Getter
	Creator Creator
	Deleter Deleter
}

type DentistasHandler struct {
	service dentists.DentistaService
}

func NewDentistasHandler(service dentists.DentistaService) *DentistasHandler {
	return &DentistasHandler{service: service}
}

func (dh *DentistasHandler) GetDentistaByID(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	dentista, err := dh.service.GetByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Dentista no encontrado"})
		return
	}

	ctx.JSON(http.StatusOK, dentista)
}

func (dh *DentistasHandler) CreateDentista(ctx *gin.Context) {
	var dentista dentists.Dentist
	if err := ctx.ShouldBindJSON(&dentista); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdDentista, err := dh.service.Create(dentista)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error interno"})
		return
	}

	ctx.JSON(http.StatusOK, createdDentista)
}

func (dh *DentistasHandler) UpdateDentista(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var dentista dentists.Dentist
	if err := ctx.ShouldBindJSON(&dentista); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedDentista, err := dh.service.ModifyByID(id, dentista)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error, dentist doesn't exist"})
		return
	}

	ctx.JSON(http.StatusOK, updatedDentista)
}

func (dh *DentistasHandler) DeleteDentista(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	err = dh.service.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error interno"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": fmt.Sprintf("Dentista %d eliminado", id)})
}
