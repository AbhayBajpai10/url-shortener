package repository

import (
	"context"
	"time"

	"github.com/AbhayBajpai10/url-shortener/internal/url/model"
	"github.com/jackc/pgx/v5"
	"github.com/redis/go-redis/v9"
)

type Repository struct {
	DB *pgx.Conn
	Redis *redis.Client
}

func NewRepository(db *pgx.Conn, redisClient *redis.Client) *Repository {
	return &Repository{
		DB: db,
		Redis: redisClient,
	}
}
func (r *Repository) CreateURL(url *model.URL) error {
	err := r.DB.QueryRow(
		context.Background(),
		`INSERT INTO urls(user_id, original_url, short_code)
		VALUES ($1, $2, $3)
		RETURNING id`,
		url.UserID, url.OriginalURL, url.ShortCode,
	).Scan(&url.ID)
	if err != nil {
		return err
	}
	return nil
}
func (r *Repository) GetByShortCode(shortCode string) (*model.URL, error) {
	url := &model.URL{}
	err := r.DB.QueryRow(
		context.Background(),
		`SELECT id,
		user_id,
		original_url,
		short_code,
		clicks
		FROM urls
		WHERE short_code = $1`,
		shortCode,
	).Scan(&url.ID, &url.UserID, &url.OriginalURL, &url.ShortCode, &url.Clicks)
	if err != nil {
		return nil, err
	}
	return url, nil
}

func (r *Repository) GetURLsByUserId(userID int, limit int, offset int) ([]model.URL, error) {
	rows, err := r.DB.Query(
		context.Background(),
		`SELECT
	id,
	original_url,
	short_code,
	clicks,
	created_at
FROM urls
WHERE user_id = $1
ORDER BY created_at DESC
limit $2
offset $3;`,
		userID,
		limit,
		offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var urls []model.URL
	for rows.Next() {
		var url model.URL
		err := rows.Scan(
			&url.ID,
			&url.OriginalURL,
			&url.ShortCode,
			&url.Clicks,
			&url.CreatedAt,
		)
		urls = append(urls, url)
		if err != nil {
			return nil, err
		}
	}
	return urls, nil
}

func (r *Repository) GetURLByID(id int) (*model.URL, error) {
	url:=&model.URL{}
	err:=r.DB.QueryRow(
		context.Background(),
		`SELECT 
		id,
		user_id,
		original_url,
		short_code,
		clicks FROM urls
		WHERE id=$1`,
		id,
	).Scan(&url.ID, &url.UserID, &url.OriginalURL, &url.ShortCode, &url.Clicks)
	if err!=nil {
		return nil,err
	}
	return url,nil
}

func (r *Repository) DeleteURL(id int) error {
	_, err:=r.DB.Exec(
		context.Background(),
		`DELETE FROM urls
		WHERE id=$1`,
		id,
	)
	if err!=nil {
		return err
	}
	return nil
}

func (r *Repository) IncrementClicks(shortCode string,) error {
	_,err:=r.DB.Exec(
		context.Background(),
		`UPDATE urls
		SET
		clicks=clicks+1
		WHERE short_code=$1`,
		shortCode,
	)
	if err!=nil {
		return err;
	}
	return nil
}

func (r *Repository) GetFromCache (shortCode string) (string, error) {
	return r.Redis.Get(
		context.Background(),
		shortCode,
	).Result()
}

func (r *Repository) SaveToCache (shortCode string, originalURL string) error {
	return r.Redis.Set(
		context.Background(),
		shortCode,
		originalURL,
		2*time.Hour,
	).Err()
}