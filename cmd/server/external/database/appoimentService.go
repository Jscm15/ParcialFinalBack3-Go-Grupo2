package database

import(
	"github.com/Jscm15/ParcialFinalBack3-Go-Grupo2/internal/appoiments"
	"database/sql"
	"fmt"
)

type SqlStore struct {
	*sql.DB
}

func NewDatabase(db *sql.DB) *SqlStore {
	return &SqlStore{db}
}

func (s *SqlStore) GetByID(id int) (appoiments.Appoiment, error) {
	var appoimentReturn appoiments.Appoiment

	query := fmt.Sprintf("SELECT * FROM appoiments WHERE id = %d;", id)
	row := s.DB.QueryRow(query)
	err := row.Scan(&appoimentReturn.ID, &appoimentReturn.Patient, &appoimentReturn.Dentist, &appoimentReturn.DateAndHour,
		&appoimentReturn.Description )
	if err != nil {
		return appoiments.Appoiment{}, err
	}
	return appoimentReturn, nil
}

func (s *SqlStore) GetByDni(dni int) (appoiments.Appoiment, error) {
	var appoimentReturn appoiments.Appoiment

	query := fmt.Sprintf("SELECT * FROM appoiments WHERE patient = '%d';", dni)
	row := s.DB.QueryRow(query)
	err := row.Scan(&appoimentReturn.ID, &appoimentReturn.Patient, &appoimentReturn.Dentist, &appoimentReturn.DateAndHour,
		&appoimentReturn.Description )
	if err != nil {
		return appoiments.Appoiment{}, err
	}
	return appoimentReturn, nil
}

func (s *SqlStore) Create(appoiment appoiments.Appoiment) (appoiments.Appoiment, error) {
	query:= fmt.Sprintf("INSERT INTO appoiments(id,patient, dentist, dateAndHour, description) VALUES('%d','%d', '%s', '%s', '%s');", appoiment.ID,appoiment.Patient, appoiment.Dentist,appoiment.DateAndHour, appoiment.Description) 
	stmt, err := s.DB.Prepare(query)
	if err != nil {
		return appoiments.Appoiment{}, err
	}
	if err != nil {
		return appoiments.Appoiment{}, err
	}
	defer s.DB.Close() 
	var result sql.Result
	result, err = stmt.Exec(appoiment.Patient, appoiment.Dentist,appoiment.DateAndHour,appoiment.Description)
	if err != nil {
	return appoiments.Appoiment{}, err
	}
	insertedId, _ := result.LastInsertId() 
	appoiment.ID = int(insertedId)
	return appoiment, nil
}

func (s *SqlStore) Modify(id int, appoiment appoiments.Appoiment) (appoiments.Appoiment, error) {
	query := fmt.Sprintf("UPDATE appoiments SET patient = '%d', dentist = %s, dateAndHour = '%s',"+
		" description = '%s' WHERE id = %v;", appoiment.Patient, appoiment.Dentist,
		appoiment.DateAndHour, appoiment.Description, appoiment.ID)
	stmt, err := s.DB.Prepare(query)
	if err != nil {
		return appoiments.Appoiment{}, err
	}

	_, err = stmt.Exec()
	if err != nil {
		return appoiments.Appoiment{}, err
	}

	return appoiment, nil
}

func (s *SqlStore) UpdateDate(id int, appoiment appoiments.Appoiment) (appoiments.Appoiment, error) {
	query := fmt.Sprintf("UPDATE appoiments SET dateAndHour = '%s', WHERE id = %v;",
		appoiment.DateAndHour, appoiment.ID)
	stmt, err := s.DB.Prepare(query)
	if err != nil {
		return appoiments.Appoiment{}, err
	}
	_, err = stmt.Exec()
	if err != nil {
		return appoiments.Appoiment{}, err
	}

	return appoiment, nil
}

func (s *SqlStore) Delete(id int)  error {
	query := fmt.Sprintf("DELETE FROM appoiments WHERE id = %d;", id)
	_, err := s.DB.Exec(query,id)
	if err!=nil{
		return err
	}
	return nil
}
