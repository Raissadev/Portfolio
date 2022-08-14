package repositories

type RepositoryInterface interface {
	All() string
	Create() string
	Update() string
	Show() string
	Delete() string
}
