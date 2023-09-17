package appoiments

import "time"

type Appoiment struct {
	ID                  int     `json:"id"`
	DniPatient          string  `json:"dni_patient" binding:"required"`
	RegistrationDentist string  `json:"registration_dentist" binding:"required"`
	DateAndHour         string `json:"date_and_hour" binding:"required"`
	Description         string  `json:"description"`
}