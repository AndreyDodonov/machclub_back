package service

import (
	"github.com/AndreyDodonov/machclub_back/pkg/models"
	"github.com/AndreyDodonov/machclub_back/pkg/repository"
)

type NewsService struct {
	repo repository.News
}

func NewNewsService(repo repository.News) *NewsService {
	return &NewsService{repo: repo}
}

func (s *NewsService) CreateNews(news models.News) (int, error) {
	return s.repo.CreateNews(news)
}
