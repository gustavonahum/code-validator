package http

import (
	"github.com/labstack/echo"
)

type StudentCodeHandler struct {
	SCUsecase studentCode.Usecase
}

func NewStudentCodeHandler(e *echo.Echo, us studentCode.Usecase) {
	handler := &StudentCodeHandler{
		SCUsecase: us,
	}
	e.GET("/student-code", handler.FetchStudentCode)
	e.POST("/student-code", handler.Store)
	e.GET("/student-code/:id", handler.GetByID)
	e.DELETE("/student-code/:id", handler.Delete)
}