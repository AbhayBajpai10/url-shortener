package repository

//only talk to db

import (
	"context"

	"github.com/AbhayBajpai10/url-shortener/internal/auth/model"
	"github.com/jackc/pgx/v5"
)

type Repository struct { //every repository object has access to the database connection.
	DB *pgx.Conn
}

func NewRepository(db *pgx.Conn) *Repository { //constructer
	return &Repository{
		DB: db,
	}
}
func (r *Repository) CreateUser(user *model.User) error {

	_, err := r.DB.Exec( //exec because it will not return the rows, only execute the query
		context.Background(),
		`INSERT INTO users(name, email, password)
		 VALUES ($1, $2, $3)`, //placeholders
		user.Name,
		user.Email,
		user.Password,
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetUserByEmail(email string) (*model.User, error) {
	row := r.DB.QueryRow(
		context.Background(),
		`
	SELECT id, name, email, password, created_at, updated_at
	FROM users
	WHERE email = $1
	`,
		email,
	)
	user:= &model.User{};
	err:=row.Scan(
		&user.ID,
		&user.Email,
		&user.Name,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err!=nil {
		return nil,err	
	}
	return user,nil
}
