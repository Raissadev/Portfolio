package repositories

import models "api/app/Models"

type UserRepository struct {
	Model     models.User
	Interface RepositoryInterface
}
