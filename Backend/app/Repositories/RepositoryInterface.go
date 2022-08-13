package repositories

type RepositoryInterface interface {
	All()
}

type Repository struct {
	Interface RepositoryInterface
}

func (r *Repository) All() string {
	return "hiol"
}
