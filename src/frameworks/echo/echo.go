package echo

import (
	"murasame29/go-fw/src/domain"
	"net/http"

	"github.com/labstack/echo/v4"
)

type echoServer struct {
	router *echo.Echo
}

func NewGinServer() *echo.Echo {
	server := &echoServer{
		router: echo.New(),
	}

	server.helloWorld()
	return server.router
}

func (s *echoServer) helloWorld() {
	s.router.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, domain.Response{
			Message: "Hello World",
		})
	})
}
