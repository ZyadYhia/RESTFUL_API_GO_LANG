package middleware

import (
	"example.com/rest-api/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Authenticate(ctx *gin.Context) {
	token := ctx.Request.Header.Get("Authorization")
	if token == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not Authorized"})
		return
	}
	userId, err := utils.VerifyJWTToken(token)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not Verified"})
		return
	}
	ctx.Set("UserId", userId)
	ctx.Next()
}
