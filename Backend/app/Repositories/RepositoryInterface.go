package repositories

type RepositoryInterface interface {
	All() string
}

type Repository struct {
	//
}

func (r *Repository) All() string {
	return "helow"
}
