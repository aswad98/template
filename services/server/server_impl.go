package server

import (
	"capregsoft.com/template/services/api"
	"github.com/labstack/echo"
)

type ServerImpl interface {
	MakeUserId(c echo.Context) error
	GetUserDataById(c echo.Context) error
}
type Server struct {
	api *api.TemplateApiImpl
}

func NewServer() *Server {
	return &Server{
		api: api.NewTemplateApiImpl(),
	}
}
func NewServerImpl(e *echo.Echo) {
	server := NewServer()
	e.POST("/createuserid", server.MakeUserId)
	e.GET("getdatabyuserid", server.GetUserDataById)

}

var _ ServerImpl = &Server{}
