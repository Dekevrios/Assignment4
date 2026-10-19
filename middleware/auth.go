package middleware

import (
	"Assignment4/auth"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Authenticate(c *gin.Context){
	// ambil header bearer
	header := c.GetHeader("Authorization")

	parts := strings.SplitN(header, " ",2)
	if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer"){
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	claims, err :=auth.ParseToken(parts[1])
	if err != nil{
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// if !(c.Request.Header.Get("Token") == "auth"){
	// 	c.AbortWithStatusJSON(401, gin.H{
	// 		"error" : "Unauthorized",
	// 	})
	// 	return
	// }

	c.Set("claims", claims)
	c.Next()
}

