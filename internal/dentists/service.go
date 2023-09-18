package dentists

type DentistRepository interface {
	GetDentistByID(id int) (Dentist, error)
	CreateDentist(dentist Dentist) (Dentist, error)
	UpdateDentistByID(id int, dentist Dentist) (Dentist, error)
	GetDentistByMatricula(matricula string) (Dentist, error)
	DeleteDentistByID(id int) error
}
type DentistaService struct {
	repository DentistRepository
}

// We implement the logic of obtaining, creating, modifying and deleting a dentist

func NewService(repository DentistRepository) *DentistaService {
	return &DentistaService{repository: repository}
}

func (s *DentistaService) GetDentistByID(id int) (Dentist, error) {
	return s.repository.GetDentistByID(id)
}

func (s *DentistaService) CreateDentist(dentista Dentist) (Dentist, error) {
	return s.repository.CreateDentist(dentista)
}

func (s *DentistaService) UpdateDentistByID(id int, dentista Dentist) (Dentist, error) {
	return s.repository.UpdateDentistByID(id, dentista)
}

func (s *DentistaService) GetDentistByMatricula(matricula string) (Dentist, error) {
	return s.repository.GetDentistByMatricula(matricula)
}

func (s *DentistaService) DeleteDentistByID(id int) error {
	return s.repository.DeleteDentistByID(id)
}
