package patients

import "time"

type PatientModel struct {
	ID     		  int       `json:"id"`
	FirstName     string    `json:"first_name"`
	LastName      string    `json:"last_name"`
	Address       string    `json:"address" `
	DNI  		  string    `json:"dni"`
	DischargeDate time.Time `json:"discharge_date"`
}