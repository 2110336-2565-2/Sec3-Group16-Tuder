package schemas

type SchemaChangePassword struct {
	Email       string `json:"email" validate:"required,email"`
	NewPassword string `json:"newpassword" validate:"required"`
}
