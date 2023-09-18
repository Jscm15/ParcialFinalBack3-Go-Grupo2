package appoiments

type Appoiment struct {
	ID                  int     `json:"id"`
	Patient          int  `json:"dni_patient" binding:"required"`
	Dentist 		string  `json:"registration_dentist" binding:"required"`
	DateAndHour         string `json:"date_and_hour" binding:"required"`
	Description         string  `json:"description"`
}