package services

type ServiceInterface interface {
	All() string
	Create() string
	Update() string
	Show() string
	Delete() string
}
