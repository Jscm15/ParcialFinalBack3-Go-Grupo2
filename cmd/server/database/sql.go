package database

import (
	"database/sql"
	"github.com/Jscm15/ParcialFinalBack3-Go-Grupo2/internal/patients"
)

type SqlStore struct {
	DB *sql.DB
}
func NewDatabase(db *sql.DB) *SqlStore {
	return &SqlStore{db}
}

func (s *SqlStore) GetByID(id int) (patients.PatientModel, error) {
	var patient patients.PatientModel
	query := "SELECT * FROM patient WHERE ID = ?;"
	row := s.DB.QueryRow(query, id)
	err := row.Scan(&patient.ID, &patient.FirstName, &patient.LastName, &patient.Address, &patient.DNI, &patient.DischargeDate)
	if err != nil {
		return patients.PatientModel{}, err
	}
	return patient, nil
}

func (s *SqlStore) Add(patient patients.PatientModel) (patients.PatientModel, error) {
	query := "INSERT INTO patient (Name, LastName, Address, DNI, DischargeDate) VALUES (?, ?, ?, ?, ?)"
	stmt, err := s.DB.Prepare(query)
	if err != nil {
		return patients.PatientModel{}, err
	}

	res, err := stmt.Exec(patient.FirstName, patient.LastName, patient.Address, patient.DNI, patient.DischargeDate)
	if err != nil {
		return patients.PatientModel{}, err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return patients.PatientModel{}, err
	}

	return patient, nil
}

func (s *SqlStore) Delete(id int) error {
	query := "DELETE FROM patient WHERE ID = ?"
	_, err := s.DB.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *SqlStore) Update(id int, patient patients.PatientModel) (patients.PatientModel, error) {
	updateQuery := "UPDATE patient SET Name = ?, LastName = ?, Address = ?, DNI = ?, DischargeDate = ? WHERE ID = ?"
	_, err := s.DB.Exec(updateQuery, patient.FirstName, patient.LastName, patient.Address, patient.DNI, patient.DischargeDate, id)
	if err != nil {
		return patients.PatientModel{}, err
	}
	return patient, nil
}