package services

import repositories "api/app/Repositories"

var repository repositories.Repository

type ServiceInterface interface {
	All() string
}

type Service struct {
	Interface ServiceInterface
}

func (s *Service) All() string {
	return repository.All()
}
