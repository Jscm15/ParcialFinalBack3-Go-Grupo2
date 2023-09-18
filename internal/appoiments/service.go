package appoiments

type Repository interface{
GetByID (id int)(Appoiment, error)
GetByDni(dni int) (Appoiment,error)
Create(appoiment Appoiment) (Appoiment, error)
Modify(id int, appoiment Appoiment) (Appoiment, error)
UpdateDate(id int, appoiment Appoiment) (Appoiment, error)
Delete(id int) error
}

type Service struct {
	repository Repository
}
func NewService(repository Repository) *Service  {
	return &Service{repository: repository}
}

func (s *Service) GetByID(id int) (Appoiment, error){
	return s.repository.GetByID(id)
}

func (s *Service) GetByDni(dni int) (Appoiment, error){
	return s.repository.GetByDni(dni)
}

func (s *Service) Create(appoiment Appoiment) (Appoiment, error){
	appoiment,err := s.repository.Create(appoiment)
	return Appoiment{}, err

}

func (s *Service) Modify(id int, appoiment Appoiment) (Appoiment, error) {
	return s.repository.Modify(id, appoiment)
}

func (s *Service) UpdateDate(id int, appoiment Appoiment) (Appoiment, error) {
	return s.repository.UpdateDate(id, appoiment)
}

func (s *Service) Delete(id int)  error {
	return s.repository.Delete(id)
}
