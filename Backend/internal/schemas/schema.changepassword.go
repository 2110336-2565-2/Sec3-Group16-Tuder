package schemas

type SchemaChangePassword struct {
	Username        string `json:"username" validate:"required"`
	Password        string `json:"password" validate:"required"`
	ConfirmPassword string `json:"confirmpassword" validate:"required"`
}

type SchemaCheckPassword struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}
