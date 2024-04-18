package routes

import (
	"example.com/rest-api/models"
	"example.com/rest-api/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func signup(ctx *gin.Context) {
	var user models.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Couldn't parse data"})
		return
	}
	err = user.Save()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Couldn't Save User"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "User Saved Successfully"})
}
func login(ctx *gin.Context) {
	var user models.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Couldn't parse data"})
		return
	}
	err = user.ValidateCredentials()
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		//ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Couldn't Authenticate Credentials"})
		return
	}

	token, err := utils.GenerateJWTToken(user.Email, user.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Couldn't Authenticate user"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Login Successful!", "token": token})

}
