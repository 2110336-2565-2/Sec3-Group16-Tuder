package schemas

type SchemaResponses struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type SchemaRegisterResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Token   string `json:"token"`
	Error   error  `json:"error"`
}

type SchemaLoginResponses struct {
	Username string `json:"username"`
	Token    string `json:"token"`
	Role     string `json:"role"`
}

type SchemaErrorResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Error   error  `json:"error"`
}

type CourseSearchResponse struct {
	Success bool                  `json:"success"`
	Message string                `json:"message"`
	Data    []*CourseSearchResult `json:"Course_serach_result"`
}
