package controller

// RegisterSchema defines the structure and validation rules for the registration request
type RegisterSchema struct {
	Username        string `json:"username" form:"username" validate:"required,min=3,max=255"`
	Password        string `json:"password" form:"password" validate:"required,min=6,max=255"`
	ConfirmPassword string `json:"confirmPassword" form:"confirmPassword" validate:"required,min=6,max=255"`
}
