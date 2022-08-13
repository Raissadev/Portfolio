package services

import repositories "api/app/Repositories"

var repository repositories.Repository

type ServiceInterface interface {
}

type Service struct {
	//
}

func (s *Service) All() string {
	return repository.All()
}
