package auth

// Suppose your project has these APIs.
// POST /api/v1/urls GET  /api/v1/urls  DELETE /api/v1/urls/:id /api/v1/urls/:id
// Should anyone on the internet be able to call these?  No. Only logged-in users.  Without middleware, every handler would start like this:
// func CreateURL(c *gin.Context){
//     // Read JWT
//     // Verify JWT
//     // Extract UserID
//     // Continue...
// }

import (
	"fmt"
	"net/http"
	"strings"

	myjwt "github.com/AbhayBajpai10/url-shortener/pkg/jwt"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "missing authorization header",
			})
			c.Abort()
			return
		}
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "invalid authorization header",
			})
			c.Abort()
			return
		}

		token := parts[1]
		claims, err := myjwt.ValidateToken(token, secret)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid token",
			})
			c.Abort()
			return
		}
		fmt.Println("JWT user id:", claims.UserID)
		fmt.Println("JWT email:", claims.Email)
		//c.Set() aur c.Get() ki keys exactly same honi chahiye.
		c.Set("userID", claims.UserID)
		c.Set("email", claims.Email)
		c.Next()
	}
}
