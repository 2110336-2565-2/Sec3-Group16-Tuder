package controllers

import (
	"net/http"

	schema "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/schemas"
	service "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/services"
	echo "github.com/labstack/echo/v4"
)


type controllerSearch struct {
	service service.ServiceSearch
}

func NewControllerSearch(s service.ServiceSearch) *controllerSearch {
	return &controllerSearch{service: s}
}

func (cR * controllerSearch) SearchContent(c echo.Context) (err error) {

	var searchContent *schema.SchemaSearch

	if err := c.Bind(&searchContent); err != nil {
		c.JSON(http.StatusBadRequest, schema.SchemaResponses{
			Success : false,
			Message : "invalid request payload",
			Data : err.Error(),
		})
		return err
	}

	searchContentInfo, err := cR.service.SearchService(searchContent)

	if err != nil {
		c.JSON(http.StatusBadRequest, schema.SchemaResponses{
			Success : false,
			Message : err.Error(),
			Data : nil,
		})
		return err
	}

	c.JSON(http.StatusOK, schema.SchemaResponses{
		Success : true,
		Message : "Search successfully",
		Data : searchContentInfo,
	})
	return nil
}