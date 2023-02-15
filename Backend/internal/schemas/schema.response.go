package schemas

type SchemaResponses struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type SchemaRegisterResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Error   error  `json:"error"`
}

type SchemaLoginResponses struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}

type SchemaErrorResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Error   error  `json:"error"`
}
