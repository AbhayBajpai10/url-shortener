package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct { //we don't send whole database, we send identity card as userid and email
	UserID int    `json:"user_id"`
	Email  string `json:"email"`

	jwt.RegisteredClaims //it contains issuer,subject, expires, issuetime, id and everything
}

func GenerateToken(userID int, email string, secret string) (string, error) {
	claims := Claims{
		UserID: userID,
		Email:  email,

		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)
	signedToken, err := token.SignedString([]byte(secret))

	if err != nil {
		return "", err
	}

	return signedToken, nil
}
//validate token
func ValidateToken(tokenString string, secret string) (*Claims, error) {
	// returning claims because auth will need claims.userid and email id 
	claims:=&Claims{}; //empty
	token,err := jwt.ParseWithClaims( //filled
		tokenString,
		claims,
		func(token *jwt.Token) (interface{}, error) { //callback which takes secret to verify the token
		return []byte(secret), nil
	},
	)
	if err!=nil {
		return nil,err;
	}
	if !token.Valid {
		return nil, errors.New("Invalid JWT Token");
	}
	return claims, nil;
}