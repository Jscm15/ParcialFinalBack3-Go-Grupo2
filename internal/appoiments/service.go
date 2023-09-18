package appoiments

type IRepositoryAppointment interface{
GetAppointmentByID (id int)(Appoiment, error)
GetAppointmentByDni(dni int) (Appoiment,error)
CreateAppointment(appoiment Appoiment) (Appoiment, error)
ModifyAppointment(id int, appoiment Appoiment) (Appoiment, error)
UpdateDate(id int, appoiment Appoiment) (Appoiment, error)
DeleteAppointment(id int) error
}

type ServiceAppointment struct {
	repository IRepositoryAppointment
}
func NewService(repository IRepositoryAppointment) *ServiceAppointment  {
	return &ServiceAppointment{repository: repository}
}

func (s *ServiceAppointment) GetAppointmentByID(id int) (Appoiment, error){
	return s.repository.GetAppointmentByID(id)
}

func (s *ServiceAppointment) GetAppointmentByDni(dni int) (Appoiment, error){
	return s.repository.GetAppointmentByDni(dni)
}

func (s *ServiceAppointment) CreateAppointment(appoiment Appoiment) (Appoiment, error){
	return s.repository.CreateAppointment(appoiment)
}

func (s *ServiceAppointment) ModifyAppointment(id int, appoiment Appoiment) (Appoiment, error) {
	return s.repository.ModifyAppointment(id, appoiment)
}

func (s *ServiceAppointment) UpdateDate(id int, appoiment Appoiment) (Appoiment, error) {
	return s.repository.UpdateDate(id, appoiment)
}

func (s *ServiceAppointment) DeleteAppointment(id int)  error {
	return s.repository.DeleteAppointment(id)
}
