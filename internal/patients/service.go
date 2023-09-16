package patients

type IService interface {
	GetByID(id int) (*PatientModel, error)
	AddPatient(patient PatientModel) (*PatientModel, error)
	ModifyByID(id int, patient PatientModel) (*PatientModel, error)
	DeleteByID(id int) error
}

type Service struct {
	repository IRepository
}

func NewService(repository IRepository) *Service {
	return &Service{repository: repository}
}

func (s *Service) GetByID(id int) (PatientModel, error) {
	return s.repository.GetByID(id)
}

func (s *Service) ModifyByID(id int, patient PatientModel) (PatientModel, error) {
	return s.repository.Update(id, patient)
}

func (s *Service) AddPatient(patient PatientModel) (PatientModel, error) {
	return s.repository.Add(patient)
}

func (s *Service) DeleteByID(id int) error {
	return s.repository.Delete(id)
}
