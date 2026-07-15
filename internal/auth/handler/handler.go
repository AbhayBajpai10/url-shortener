package handler 
import (
	"net/http"
	"github.com/AbhayBajpai10/url-shortener/internal/auth/service"
	"github.com/AbhayBajpai10/url-shortener/internal/auth/model"
	"github.com/AbhayBajpai10/url-shortener/internal/auth/dto"
	"github.com/gin-gonic/gin"
)
type Handler struct {
	Service *service.Service
}
func NewHandler(s *service.Service) *Handler {
	return &Handler{
		Service: s,
	}
}
func (handle *Handler) Register(c *gin.Context) {
	var req dto.RegisterRequest;
	if err:= c.ShouldBindJSON(&req); err!= nil {
		c.JSON(http.StatusBadRequest, gin.H{ //gin.H is a shortcut type create onflying json objects
			"error": err.Error(),
		})
		return;
	}
	user:= &model.User{
		Name: req.Name,
		Email: req.Email,
		Password: req.Password,
	}
	err:= handle.Service.RegisterUser(user);
	if err!= nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return;
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User Registered Successfully",
	})
}

func (handle *Handler) Login(c *gin.Context) {
	var req dto.LoginRequest;
	if err:=c.ShouldBindJSON(&req); err!=nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return;
	}
	res, err:= handle.Service.Login(&req);
	if err!=nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, res);
}