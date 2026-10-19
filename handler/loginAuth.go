package handler

import (
	"Assignment4/auth"
	"Assignment4/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-crypt/x/bcrypt"
)

type LoginRequest struct{
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
		return
	}

	user, found := model.FindbyUsername(req.Username)
	if !found{
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil{
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Bad Request"})
		return
	}

	token, err := auth.GenerateToken(user)
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"accessToken": token})

}

func Me(c *gin.Context){
	value, exists := c.Get("claims")
	if !exists{
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unatuthorized"})
		return
	}

	claims, ok := value.(*auth.Claims)
	if !ok{
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unatuthorized"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id" : claims.ID,
		"username": claims.Username,
		"name": claims.Name,
	})

}