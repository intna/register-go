package controller

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"register/src/models/schemas"
	"register/src/services"
)

var validate = validator.New(validator.WithRequiredStructEnabled())

func RegisterHandler(c *gin.Context) error {
	var req schemas.RegisterSchema
	// Bind JSON input
	if err := c.ShouldBind(&req); err != nil {
		return errors.New("invalid params")
	}

	// Custom validation for matching passwords
	if err := validate.Struct(req); err != nil {
		return errors.New("invalid params")
	}

	if req.Password != req.ConfirmPassword {
		return errors.New("ConfirmPassword should as same as password")
	}

	if err := services.Register(req); err != nil {
		return err
	}

	c.JSON(http.StatusOK, gin.H{"message": "Registration successful"})
	return nil
}

// LogoutHandler handles user logout
func LogoutHandler(c *gin.Context) error {
	c.JSON(http.StatusOK, gin.H{
		"message": "Logout successful",
	})
	return nil
}
