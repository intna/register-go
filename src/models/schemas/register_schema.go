package schemas

// RegisterSchema defines the structure and validation rules for the registration request
type RegisterSchema struct {
	Email           string `json:"email" form:"email" validate:"required,min=3,max=255,email"`
	Password        string `json:"password" form:"password" validate:"required,min=6,max=255"`
	ConfirmPassword string `json:"confirmPassword" form:"confirmPassword" validate:"required,min=6,max=255"`
}
