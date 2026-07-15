package service

import (
	"errors"
	"fmt"

	"github.com/AbhayBajpai10/url-shortener/internal/url/dto"
	"github.com/AbhayBajpai10/url-shortener/internal/url/model"
	"github.com/AbhayBajpai10/url-shortener/internal/url/repository"
	"github.com/AbhayBajpai10/url-shortener/pkg/shortener"
	"github.com/redis/go-redis/v9"
)

type Service struct {
	Repo *repository.Repository
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Repo: repo,
	}
}

// rec. req->generate shortcode->ccreate model-> repos.createurl()->build response.
func (s *Service) CreateURL(userID int, req *dto.CreateURLRequest) (*dto.CreateURLResponse, error) {
	shortcode := shortener.GenerateShortCode(6)
	url := &model.URL{
		UserID:      userID,
		OriginalURL: req.OriginalURL,
		ShortCode:   shortcode,
	}
	fmt.Println(
		"Creating URL with:",
		"userID =", userID,
		"originalURL =", req.OriginalURL,
	)
	err := s.Repo.CreateURL(url)
	if err != nil {
		return nil, err
	}
	return &dto.CreateURLResponse{
		ID:          url.ID,
		OriginalURL: url.OriginalURL,
		ShortCode:   url.ShortCode,
		ShortURL:    "http://localhost:8080/" + shortcode,
	}, nil
}

func (s *Service) GetOriginalURL(shortCode string) (string, error) {
	cachedURL, err := s.Repo.GetFromCache(shortCode)
	//cache hit
	if err == nil {
		_ = s.Repo.IncrementClicks(shortCode)
		return cachedURL, nil
	}
	if err != redis.Nil {
		fmt.Println("Redis Error", err)
	}

	OrgURL, err := s.Repo.GetByShortCode(shortCode)
	if err != nil {
		return "", err
	}
	//save to cache
	err = s.Repo.SaveToCache(shortCode, OrgURL.OriginalURL)
	if err != nil {
		fmt.Println("Redis Error", err)
	}
	err = s.Repo.IncrementClicks(shortCode)
	if err != nil {
		return "", err
	}
	return OrgURL.OriginalURL, nil
}

func (s *Service) GetUrls(userID int, pageNo int, limit int) ([]dto.URLResponse, error) {
	offset:=(pageNo-1)*limit
	urls, err := s.Repo.GetURLsByUserId(userID, limit, offset)
	if err != nil {
		return nil, err
	}
	var response []dto.URLResponse
	for _, url := range urls {
		response = append(response, dto.URLResponse{
			ID:          url.ID,
			OriginalURL: url.OriginalURL,
			ShortCode:   url.ShortCode,
			Clicks:      url.Clicks,
			ShortURL: "http://localhost:8080/" + url.ShortCode,
		})
	}
	return response, nil
}

func (s *Service) DeleteURL(id int, userID int) error {
	url, err := s.Repo.GetURLByID(id)
	if err != nil {
		return err
	}
	if url.UserID != userID {
		return errors.New("Forbidden")
	}
	return s.Repo.DeleteURL(id)
}
