package database

import (
	"database/sql"
	"github.com/Jscm15/ParcialFinalBack3-Go-Grupo2/internal/patients"
)

type SqlStore struct {
	DB *sql.DB
}

func (s *SqlStore) GetPatientById(id int) (*patients.PatientModel, error) {
	var patient patients.PatientModel
	query := "SELECT * FROM patient WHERE ID = ?;"
	row := s.DB.QueryRow(query, id)
	err := row.Scan(&patient.ID, &patient.FirstName, &patient.LastName, &patient.Address, &patient.DNI, &patient.DischargeDate)
	if err != nil {
		return nil, err
	}
	return &patient, nil
}

func (s *SqlStore) CreatePatient(patient patients.PatientModel) (*patients.PatientModel, error) {
	query := "INSERT INTO patient (Name, LastName, Address, DNI, DischargeDate) VALUES (?, ?, ?, ?, ?)"
	stmt, err := s.DB.Prepare(query)
	if err != nil {
		return nil, err
	}

	res, err := stmt.Exec(patient.FirstName, patient.LastName, patient.Address, patient.DNI, patient.DischargeDate)
	if err != nil {
		return nil, err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return nil, err
	}

	return &patient, nil
}

func (s *SqlStore) DeletePatient(id int) error {
	query := "DELETE FROM patient WHERE ID = ?"
	_, err := s.DB.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *SqlStore) UpdatePatient(id int, patient patients.PatientModel) (*patients.PatientModel, error) {
	updateQuery := "UPDATE patient SET Name = ?, LastName = ?, Address = ?, DNI = ?, DischargeDate = ? WHERE ID = ?"
	_, err := s.DB.Exec(updateQuery, patient.FirstName, patient.LastName, patient.Address, patient.DNI, patient.DischargeDate, id)
	if err != nil {
		return nil, err
	}
	return &patient, nil
}