package dentists

type DentistRepository interface {
	GetByID(id int) (Dentist, error)
	Create(dentist Dentist) (Dentist, error)
	Modify(id int, dentist Dentist) (Dentist, error)
	UpdateField(id int, denstist Dentist) (Dentist, error)
	Delete(id int) error
}
type DentistaService struct {
	repository DentistRepository
}

// We implement the logic of obtaining, creating, modifying and deleting a dentist

func NewService(repository DentistRepository) *DentistaService {
	return &DentistaService{repository: repository}
}

func (s *DentistaService) GetByID(id int) (Dentist, error) {
	return s.repository.GetByID(id)
}

func (s *DentistaService) Create(dentista Dentist) (Dentist, error) {
	return s.repository.Create(dentista)
}

func (s *DentistaService) ModifyByID(id int, dentista Dentist) (Dentist, error) {
	return s.repository.Modify(id, dentista)
}

func (s *DentistaService) UpdateField(id int, fieldName string, fieldValue interface{}) (Dentist, error) {
	dentista, err := s.repository.GetByID(id)
	if err != nil {
		return Dentist{}, err
	}

	switch fieldName {
	case "Name":
		dentista.Name = fieldValue.(string)
	case "Lastname":
		dentista.Lastname = fieldValue.(string)
	case "matricula":
		dentista.Matricula = fieldValue.(string)
	}

	return s.repository.Modify(id, dentista)
}

func (s *DentistaService) Delete(id int) error {
	return s.repository.Delete(id)
}
