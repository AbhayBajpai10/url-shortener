package model 
import ( "time" );
//db schema
type User struct {
	ID int
	Name string
	Email string
	Password string
	CreatedAt time.Time
	UpdatedAt time.Time
}
