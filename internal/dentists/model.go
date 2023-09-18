package dentists

type Dentist struct {
	ID        int    `json:"id"`
	Lastname  string `json:"LastName"`
	Name      string `json:"name"`
	Matricula string `json:"matricula"`
}
