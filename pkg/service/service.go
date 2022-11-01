package service

import (
	"github.com/AndreyDodonov/machclub_back/pkg/models"
	"github.com/AndreyDodonov/machclub_back/pkg/repository"
)

type Authorization interface {
	// создать пользователя
	CreateUser(user models.User) (int, error)
	// сгенерировать токен
	GenerateToken(email, password string) (string, error)
	// распарсить токен
	ParseToken(token string) (int, error)
}

type News interface {
	// создать новость
	CreateNews(news models.News) (int, error)
	// получить все новости
	getAllNews() ([]models.News, error)
	// получить одну новость по id
	getNewsById(newsId int) (models.News, error)
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
	return &Service{
		Authorization: NewAuthService(rep.Authorization),
	}
}
