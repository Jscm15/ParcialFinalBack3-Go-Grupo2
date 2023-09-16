package patients

import (
	"errors"
)

type IStore interface {
	GetPatientById(id int) (*PatientModel, error)
	CreatePatient(patient PatientModel) (*PatientModel, error)
	UpdatePatient(id int, patient PatientModel) (*PatientModel, error)
	DeletePatient(id int) error
}

type Repository struct {
	Store IStore
}

type IRepository interface {
	GetByID(id int) (PatientModel, error)
	Update(id int, patient PatientModel) (PatientModel, error)
	Add(patient PatientModel) (PatientModel, error)
	Delete(id int) error
}

func (r *Repository) GetByID(id int) (*PatientModel, error) {
	patient, err := r.Store.GetPatientById(id)
	if err != nil {
		return nil, errors.New("el paciente no existe")
	}
	return patient, nil
}

func (r *Repository) Add(patient PatientModel) (*PatientModel, error) {
	existingPatient, _ := r.Store.GetPatientById(patient.ID)
	if existingPatient != nil {
		return nil, errors.New("ya existe un paciente con ese DNI")
	}
	p, err := r.Store.CreatePatient(patient)
	if err != nil {
		return nil, errors.New("error al crear el paciente")
	}
	return p, nil
}

func (r *Repository) Delete(id int) error {
	_, err := r.GetByID(id)
	if err != nil {
		return err
	}
	err = r.Store.DeletePatient(id)
	if err != nil {
		return errors.New("error al eliminar el paciente")
	}
	return nil
}

func (r *Repository) Update(id int, patient PatientModel) (*PatientModel, error) {
	_, err := r.Store.GetPatientById(id)
	if err != nil {
		return nil, errors.New("el paciente no existe")
	}
	p, err := r.Store.UpdatePatient(id, patient)
	if err != nil {
		return nil, errors.New("error al actualizar el paciente")
	}
	return p, nil
}
