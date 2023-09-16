package handler

import (
	"net/http"
	"strconv"
	"github.com/Jscm15/ParcialFinalBack3-Go-Grupo2/internal/patients"
	"github.com/gin-gonic/gin"
)

type PatientGetter interface {
	GetByID(id int) (patients.PatientModel, error)
}

type PatientCreator interface {
	ModifyByID(id int, patient patients.PatientModel) (patients.PatientModel, error)
	AddPatient(patient patients.PatientModel) (patients.PatientModel, error)
}

type PatientDelete interface {
	DeleteByID(id int) error
}

type PatientHandler struct {
	patientGetter  PatientGetter
	patientCreator PatientCreator
	patientDelete  PatientDelete
}

func NewPatientHandler(getter PatientGetter, creator PatientCreator, delete PatientDelete) *PatientHandler {
	return &PatientHandler{
		patientGetter:  getter,
		patientCreator: creator,
		patientDelete:  delete,
	}
}

func (p *PatientHandler) GetPatientByID(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	patient, err := p.patientGetter.GetByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "patient not found"})
		return
	}
	ctx.JSON(http.StatusOK, patient)
}

func (p *PatientHandler) PutPatient(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "invalid id"})
		return
	}
	patientRequest := patients.PatientModel{}
	err = ctx.BindJSON(&patientRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	patient, err := p.patientCreator.ModifyByID(id, patientRequest)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "internal error"})
		return
	}
	ctx.JSON(200, patient)
}

func (p *PatientHandler) CreatePatient(ctx *gin.Context) {
	patientRequest := patients.PatientModel{}
	err := ctx.BindJSON(&patientRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	patient, err := p.patientCreator.AddPatient(patientRequest)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "internal error"})
		return
	}
	ctx.JSON(200, patient)
}

func (p *PatientHandler) DeletePatient(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "invalid id"})
		return
	}
	err = p.patientDelete.DeleteByID(id)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "internal error"})
		return
	}

	ctx.JSON(200, "")
}
