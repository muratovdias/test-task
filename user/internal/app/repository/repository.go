package repository

type Repository struct {
	User
}

func NewRepository(mongo *MongoDB) *Repository {
	return &Repository{
		User: NewUserRepo(mongo.users),
	}
}
