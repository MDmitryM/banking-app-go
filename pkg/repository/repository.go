package repository

type Authorization interface {
}

type Transaction interface {
}

type Statistic interface {
}

type Category interface {
}

type Repository struct {
	Authorization
	Transaction
	Statistic
	Category
}

func NewRepository(db *MongoDB) *Repository {
	return &Repository{}
}
