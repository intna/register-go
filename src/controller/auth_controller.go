package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"register/src/constants"
	"register/src/models/schemas"
	"register/src/services"
)

var validate = validator.New(validator.WithRequiredStructEnabled())

func RegisterHandler(c *gin.Context) {
	var req schemas.RegisterSchema
	// Bind JSON input
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(constants.BAD_REQUEST, gin.H{"error": err.Error()})
		return
	}

	// Custom validation for matching passwords
	if err := validate.Struct(req); err != nil {
		c.JSON(constants.BAD_REQUEST, gin.H{"error": err.Error()})
		return
	}

	if req.Password != req.ConfirmPassword {
		c.JSON(constants.BAD_REQUEST, gin.H{"error": "invalid confirmPassword!"})
		return
	}

	services.Register(c, req)
}

// LogoutHandler handles user logout
func LogoutHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Logout successful",
	})
}
