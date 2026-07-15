package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/AbhayBajpai10/url-shortener/internal/url/dto"
	"github.com/AbhayBajpai10/url-shortener/internal/url/service"
	"github.com/gin-gonic/gin"
)

//bind json-> getuseridfrom middleware-> call service -> return json
type Handler struct{
	Service *service.Service
}
func NewHandler(s *service.Service) *Handler{
	return &Handler{
		Service: s,
	}
}
func (h * Handler) CreateURL(c *gin.Context) {
	var req dto.CreateURLRequest
	if err:=c.ShouldBindJSON(&req); err!=nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	userID:=c.GetInt("userID")
	fmt.Println("Handler user id", userID)
	res,err:=h.Service.CreateURL(
		userID,
		&req,
	)
	if err!=nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return;
	}
	c.JSON(http.StatusCreated,res)
}

func (h *Handler) RedirectURL(c* gin.Context) {
	shortCode:=c.Param("shortCode")
	url,err:= h.Service.GetOriginalURL(shortCode)
	if err!=nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return;
	}
	c.Redirect(http.StatusFound,url)
}

func (h *Handler) GetMyURLs (c *gin.Context) {
	userID:=c.GetInt("userID")
	pageStr:=c.DefaultQuery("page","1")
	limitStr:=c.DefaultQuery("limit","5")
	pageNo,_:=strconv.Atoi(pageStr)
	limit,_:=strconv.Atoi(limitStr)
	res, err:= h.Service.GetUrls(userID, pageNo, limit)
	if err!=nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *Handler) DeleteURL (c *gin.Context) {
	idStr := c.Param("id")
	id, err:= strconv.Atoi(idStr);
	if err!=nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		return
	}
	userID := c.GetInt("userID")
	err = h.Service.DeleteURL(id, userID)
	if err!=nil {
		c.JSON(http.StatusForbidden, gin.H{
			"error": err.Error(),
		})
	}
	c.JSON(http.StatusAccepted, gin.H{
		"message": "URL DELETED SUCCESSFULLY",
	})
}