package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"github.com/Jscm15/ParcialFinalBack3-Go-Grupo2/internal/dentists"
	"github.com/gin-gonic/gin"
)

type Getter interface {
	GetDentistByID(id int) (dentists.Dentist, error)
	GetDentistByMatricula(matricula string) (dentists.Dentist, error)
}

type Creator interface {
	CreateDentist(dentista dentists.Dentist) (dentists.Dentist, error)
	UpdateDentistByID(id int, dentista dentists.Dentist) (dentists.Dentist, error)
}

type Deleter interface {
	DeleteDentistByID(id int) error
}
type DentistsHandler struct {
	dentistsGetter  Getter
	dentistsCreator Creator
	dentistsDeleter Deleter
}


func NewDentistsHandler(getter Getter, creator Creator, delete Deleter) *DentistsHandler {
	return &DentistsHandler{
		dentistsGetter:  getter,
		dentistsCreator: creator,
		dentistsDeleter:  delete,
	}
}

func (dh *DentistsHandler) GetDentistByID(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	dentista, err := dh.dentistsGetter.GetDentistByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "internal error"})
		return
	}

	ctx.JSON(http.StatusOK, dentista)
}

func (dh *DentistsHandler) GetDentistByMatricula(ctx *gin.Context) {
	idParam := ctx.Param("matricula")
	
	dentista, err := dh.dentistsGetter.GetDentistByMatricula(idParam)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "invalid id"})
		return
	}

	ctx.JSON(http.StatusOK, dentista)
}

func (dh *DentistsHandler) CreateDentist(ctx *gin.Context) {
	var dentista dentists.Dentist
	if err := ctx.ShouldBindJSON(&dentista); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdDentista, err := dh.dentistsCreator.CreateDentist(dentista)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}

	ctx.JSON(http.StatusOK, createdDentista)
}

func (dh *DentistsHandler) UpdateDentistByID(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var dentista dentists.Dentist
	if err := ctx.ShouldBindJSON(&dentista); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedDentista, err := dh.dentistsCreator.UpdateDentistByID(id, dentista)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error, dentist doesn't exist"})
		return
	}

	ctx.JSON(http.StatusOK, updatedDentista)
}

func (dh *DentistsHandler) DeleteDentistByID(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	err = dh.dentistsDeleter.DeleteDentistByID(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": fmt.Sprintf("Dentist %d deleted", id)})
}

