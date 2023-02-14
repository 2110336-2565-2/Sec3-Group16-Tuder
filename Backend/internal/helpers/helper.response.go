package helpers

import(
	echo "github.com/labstack/echo/v4"
	schema "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/schemas"
)

func APIResponse(c echo.Context, Message string, StatusCode int, Method string, Data interface{}) {

	jsonResponse := schema.SchemaResponses{
		StatusCode: StatusCode,
		Method:     Method,
		Message:    Message,
		Data:       Data,
	}

	if StatusCode >= 400 {
		c.JSON(StatusCode, jsonResponse)
	} else {
		c.JSON(StatusCode, jsonResponse)
	}
}

func ValidatorErrorResponse(e echo.Context, StatusCode int, Method string, Error interface{}) {
	errResponse := schema.SchemaErrorResponse{
		StatusCode: StatusCode,
		Method:     Method,
		Error:      Error,
	}

	e.JSON(StatusCode, errResponse)
}