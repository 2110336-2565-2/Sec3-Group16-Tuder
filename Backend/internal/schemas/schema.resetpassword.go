package schemas

type SchemaResetPassword struct {
	Email string `json:"email" validate:"required,email"`
}

type SchemaNewPassword struct {
	Token           string `json:"token" validate:"required"`
	Password        string `json:"password" validate:"required"`
	ConfirmPassword string `json:"confirm_password" validate:"required"`
}
