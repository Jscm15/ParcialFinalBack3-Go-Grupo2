package database

import (
	"fmt"
	"github.com/Jscm15/ParcialFinalBack3-Go-Grupo2/internal/dentists"
	"log"
)
func (s *SqlStore) GetDentistByID(id int) (dentists.Dentist, error) {
	var dentistReturn dentists.Dentist

	query := fmt.Sprintf("SELECT * FROM dentists WHERE id = %d;", id)
	row := s.DB.QueryRow(query)
	err := row.Scan(&dentistReturn.ID, &dentistReturn.Name, &dentistReturn.Lastname, &dentistReturn.Matricula)
	if err != nil {
		return dentists.Dentist{}, err
	}
	return dentistReturn, nil
}

func (s *SqlStore) GetDentistByMatricula(matricula string) (dentists.Dentist, error) {
	var dentistReturn dentists.Dentist

	query := fmt.Sprintf("SELECT * FROM dentists WHERE matricula = '%s';", matricula)
	row := s.DB.QueryRow(query)
	err := row.Scan(&dentistReturn.ID, &dentistReturn.Name, &dentistReturn.Lastname, &dentistReturn.Matricula)
	if err != nil {
		return dentists.Dentist{}, err
	}
	return dentistReturn, nil
}

func (s *SqlStore) CreateDentist(d dentists.Dentist) (dentists.Dentist, error) {
	stmt, err := s.DB.Prepare("INSERT INTO dentists (name, lastname, matricula) VALUES ($1, $2, $3) RETURNING id")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	var insertedId int
	err = stmt.QueryRow(d.Name, d.Lastname, d.Matricula).Scan(&insertedId)

	if err != nil {
		return dentists.Dentist{}, err
	}

	d.ID = insertedId
	return d, nil
}

func (s *SqlStore) UpdateDentistByID(id int, d dentists.Dentist) (dentists.Dentist, error) {
	stmt, err := s.DB.Prepare("UPDATE dentists SET name = $1, lastname = $2, Matricula = $3 WHERE id = $4")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Query(d.Name, d.Lastname, d.Matricula, id)
	if err != nil {
		return dentists.Dentist{}, err
	}
	d.ID = id
	return d, nil
}

func (s *SqlStore) DeleteDentistByID(id int) error {
	stmt, err := s.DB.Prepare("DELETE FROM dentists WHERE id = $1")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Query(id)
	if err != nil {
		return err
	}
	return nil
}
