package patients

type IRepositoryPatient interface {
	GetPatientByID(id int) (PatientModel, error)
	Update(id int, patient PatientModel) (PatientModel, error)
	Add(patient PatientModel) (PatientModel, error)
	Delete(id int) error
}

type ServicePatient struct {
	repository IRepositoryPatient
}

func NewService(repository IRepositoryPatient) *ServicePatient {
	return &ServicePatient{repository: repository}
}

func (s *ServicePatient) GetPatientByID(id int) (PatientModel, error) {
	return s.repository.GetPatientByID(id)
}

func (s *ServicePatient) ModifyByID(id int, patient PatientModel) (PatientModel, error) {
	return s.repository.Update(id, patient)
}

func (s *ServicePatient) AddPatient(patient PatientModel) (PatientModel, error) {
	return s.repository.Add(patient)
}

func (s *ServicePatient) DeleteByID(id int) error {
	return s.repository.Delete(id)
}
