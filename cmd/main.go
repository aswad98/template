package main

import (
	"capregsoft.com/template/services/server"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	server.NewServerImpl(e)
	e.Logger.Fatal(e.Start(":8000"))
}
