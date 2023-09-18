package patients

type IRepositoryPatient interface {
	GetPatientByID(id int) (PatientModel, error)
	ModifyPatientByID(id int, patient PatientModel) (PatientModel, error)
	AddPatient(patient PatientModel) (PatientModel, error)
	DeletePatientByID(id int) error
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

func (s *ServicePatient) ModifyPatientByID(id int, patient PatientModel) (PatientModel, error) {
	return s.repository.ModifyPatientByID(id, patient)
}

func (s *ServicePatient) AddPatient(patient PatientModel) (PatientModel, error) {
	return s.repository.AddPatient(patient)
}

func (s *ServicePatient) DeletePatientByID(id int) error {
	return s.repository.DeletePatientByID(id)
}
