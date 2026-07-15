package service

import (
	"errors"

	"github.com/AbhayBajpai10/url-shortener/config"
	"github.com/AbhayBajpai10/url-shortener/internal/auth/dto"
	"github.com/AbhayBajpai10/url-shortener/internal/auth/model"
	"github.com/AbhayBajpai10/url-shortener/internal/auth/repository"
	"github.com/AbhayBajpai10/url-shortener/pkg/jwt"
	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	Repo *repository.Repository //Repository ko store karta hai) are jaise cpp me store krte the car engine ko same vaise hi
	Config *config.Config 
}

func NewService(repo *repository.Repository, cfg *config.Config) *Service {
	//today service has only one field. but it may grow in future
	//then for every object we create, 've to write service  := Service{
	// Repo: repo,
	// Logger: log,
	// Config: cfg,
	// Cache: redis,
	//btt using constructer we can simple write service := NewService(repo, logger, config, cache);
	return &Service{
		Repo: repo,
		Config: cfg,
	}
}

// register user
func (s *Service) RegisterUser(user *model.User) error {
	existingUser, err := s.Repo.GetUserByEmail(user.Email)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return err
	}
	if existingUser != nil {
		return errors.New("user already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(user.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	return s.Repo.CreateUser(user)
}

// login user
func (s *Service) Login(req *dto.LoginRequest) (*dto.LoginResponse, error) {
	user, err := s.Repo.GetUserByEmail(req.Email)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.New("invalid email or password")
		}
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(req.Password),
	)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}
	token, err:= jwt.GenerateToken(user.ID, user.Email, s.Config.JWTSecret)
	if err!=nil {
		return nil, err
	}
	return &dto.LoginResponse{
		Token: token,
		User: dto.UserResponse{
			ID: user.ID,
			Name: user.Name,
			Email: user.Email,
		},
	},nil
}
