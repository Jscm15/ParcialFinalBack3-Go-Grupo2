package database

import (
	"github.com/Jscm15/ParcialFinalBack3-Go-Grupo2/internal/patients"
)

func (s *SqlStore) GetPatientByID(id int) (patients.PatientModel, error) {
	var patient patients.PatientModel
	query := "SELECT * FROM patients WHERE ID = ?;"
	row := s.DB.QueryRow(query, id)
	err := row.Scan(&patient.ID, &patient.FirstName, &patient.LastName, &patient.Address, &patient.DNI, &patient.DischargeDate)
	if err != nil {
		return patients.PatientModel{}, err
	}
	return patient, nil
}

func (s *SqlStore) AddPatient(patient patients.PatientModel) (patients.PatientModel, error) {
	query := "INSERT INTO patients (FirstName, LastName, Address, DNI, DischargeDate) VALUES (?, ?, ?, ?, ?)"
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
	insertedId, _ := res.LastInsertId() 
	patient.ID = int(insertedId)


	return patient, nil
}

func (s *SqlStore) DeletePatientByID(id int) error {
	query := "DELETE FROM patients WHERE ID = ?"
	_, err := s.DB.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *SqlStore) ModifyPatientByID(id int, patient patients.PatientModel) (patients.PatientModel, error) {
	updateQuery := "UPDATE patients SET FirstName = ?, LastName = ?, Address = ?, DNI = ?, DischargeDate = ? WHERE ID = ?"
	stmt, err := s.DB.Prepare(updateQuery)
	defer stmt.Close()
	_, err = stmt.Exec(patient.FirstName, patient.LastName, patient.Address, patient.DNI, patient.DischargeDate, id)
	if err != nil {
		return patients.PatientModel{}, err
	}
	return patient, nil
}