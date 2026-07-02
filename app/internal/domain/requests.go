package domain

type RegisterUserRequest struct {
	Username        string `json:"name"           form:"name" validate:"required,min=2,max=100"`
	Email           string `json:"email"          form:"email" validate:"required,email"`
	Password        string `json:"password"       form:"password" validate:"required,min=4,max=20"`
	ConfirmPassword string `form:"confirmPassword" validate:"required,eqfield=Password"`
	Gender          string `json:"gender"         form:"gender" validate:"required"`
	Dob             string `json:"dob"            form:"dob" validate:"required"`
	Currency        int    `json:"currency"       form:"currency" validate:"required"`
}
