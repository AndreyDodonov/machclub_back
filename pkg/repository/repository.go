package repository

import (
	"github.com/AndreyDodonov/machclub_back/pkg/models"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GetUser(email, password string) (models.User, error)
}

type News interface {
}

type NewsItem interface {
}

type Article interface {
}

type ArticleItem interface {
}

type Repository struct {
	Authorization
	News
	NewsItem
	Article
	ArticleItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),

	}
}
