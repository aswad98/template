package server

import (
	"net/http"

	"capregsoft.com/template/models"
	server_errors "capregsoft.com/template/services/server/errors"
	"github.com/labstack/echo"
)

func (s *Server) MakeUserId(c echo.Context) error {
	userIdData := models.NewUserData()
	if err := c.Bind(&userIdData); err != nil {
		return server_errors.NewInternalServerError(err.Error())
	}
	userIdDataResp, err := s.api.CreateUserId(userIdData)
	if err != nil {
		return server_errors.NewInternalServerError(err.Error())
	}
	result := models.Response{
		Status:  http.StatusOK,
		Message: "Your id created succesfully ",
		Data:    userIdDataResp,
	}
	return c.JSON(http.StatusOK, result)
}
