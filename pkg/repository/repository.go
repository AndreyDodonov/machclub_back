package repository

import "github.com/jmoiron/sqlx"

type Authorization interface {
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
	return &Repository{}
}