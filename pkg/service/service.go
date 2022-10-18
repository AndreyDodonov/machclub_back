package service

import "github.com/AndreyDodonov/machclub_back/pkg/repository"

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

type Service struct {
	Authorization
	News
	NewsItem
	Article
	ArticleItem
}

func NewService(rep *repository.Repository) *Service {
	return &Service{}
}