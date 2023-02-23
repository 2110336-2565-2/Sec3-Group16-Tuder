package controllers

import (
	"net/http"

	schema "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/schemas"
	service "github.com/2110336-2565-2/Sec3-Group16-Tuder/internal/services"
	echo "github.com/labstack/echo/v4"
)

type controllerCourseSearch struct {
	service service.ServiceCourseSearch
}

func NewControllerCourseSearch(s service.ServiceCourseSearch) *controllerCourseSearch {
	return &controllerCourseSearch{service: s}
}

func (cR *controllerCourseSearch) SearchContent(c echo.Context) (err error) {

	var searchContent *schema.CourseSearch

	if err := c.Bind(&searchContent); err != nil {
		c.JSON(http.StatusBadRequest, schema.SchemaResponses{
			Success: false,
			Message: "invalid request payload",
			Data:    err.Error(),
		})
		return err
	}
	//fmt.Println(searchContent)

	Course_search_result, err := cR.service.CourseSearchService(searchContent)
	if err != nil {
		c.JSON(http.StatusBadRequest, schema.SchemaResponses{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
		return err
	}

	c.JSON(http.StatusOK, schema.CourseSearchResponse{
		Success: true,
		Message: "Search successfully",
		Data:    Course_search_result,
	})
	return
}

func (cR *controllerCourseSearch) GetAllCourse(c echo.Context) (err error) {
	var searchContent *schema.CourseSearch

	if err := c.Bind(&searchContent); err != nil {
		c.JSON(http.StatusBadRequest, schema.SchemaResponses{
			Success: false,
			Message: "invalid request payload",
			Data:    err.Error(),
		})
		return err
	}
	//fmt.Println(searchContent)

	Course_search_result, err := cR.service.SearchAllCourse(searchContent)

	if err != nil {
		c.JSON(http.StatusBadRequest, schema.SchemaResponses{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
		return err
	}

	c.JSON(http.StatusOK, schema.CourseSearchResponse{
		Success: true,
		Message: "Search successfully",
		Data:    Course_search_result,
	})
	return
}
