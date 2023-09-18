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

	query := fmt.Sprintf("SELECT * FROM dentists WHERE Registry = '%s';", matricula)
	row := s.DB.QueryRow(query)
	err := row.Scan(&dentistReturn.ID, &dentistReturn.Name, &dentistReturn.Lastname, &dentistReturn.Matricula)
	if err != nil {
		return dentists.Dentist{}, err
	}
	return dentistReturn, nil
}

func (s *SqlStore) CreateDentist(d dentists.Dentist) (dentists.Dentist, error) {
	stmt, err := s.DB.Prepare("INSERT INTO dentists (FirstName, LastName, Registry) VALUES (?, ?, ?)")
	if err != nil {
		return dentists.Dentist{},err
	}
	defer stmt.Close()

	
	res,err := stmt.Exec(d.Name, d.Lastname, d.Matricula)
	
	_, err = res.RowsAffected()
	if err != nil {
		return dentists.Dentist{}, err
	}
	insertedId, _ := res.LastInsertId() 
	d.ID = int(insertedId)
	return d, nil
}

func (s *SqlStore) UpdateDentistByID(id int, d dentists.Dentist) (dentists.Dentist, error) {
	stmt, err := s.DB.Prepare("UPDATE dentists SET FirstName = ?, LastName = ?, Registry = ? WHERE ID = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(d.Name, d.Lastname, d.Matricula, id)
	if err != nil {
		return dentists.Dentist{}, err
	}
	return d, nil
}

func (s *SqlStore) DeleteDentistByID(id int) error {
	stmt, err := s.DB.Prepare("DELETE FROM dentists WHERE ID = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}
