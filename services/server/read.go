package server

import (
	"net/http"
	"strconv"

	"capregsoft.com/template/models"
	server_errors "capregsoft.com/template/services/server/errors"
	"github.com/labstack/echo"
)

func (s *Server) GetUserDataById(c echo.Context) error {
	userIdstr := c.Param("user_id")
	userId, err := strconv.Atoi(userIdstr)
	if err != nil {
		return server_errors.NewInternalServerError(err.Error())
	}
	userDataResp, err := s.api.GetUserDataById(userId)
	if err != nil {
		return server_errors.NewUnauthorizedError("wrong user id")
	}
	result := models.Response{
		Status:  http.StatusOK,
		Message: "data fetched successfully",
		Data:    userDataResp,
	}
	return c.JSON(http.StatusOK, result)

}
