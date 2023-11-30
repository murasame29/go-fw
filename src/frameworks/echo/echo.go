package echo

import (
	"murasame29/go-fw/src/domain"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ginServer struct {
	router *echo.Echo
}

func NewGinServer() *echo.Echo {
	server := &ginServer{
		router: echo.New(),
	}

	server.helloWorld()
	return server.router
}

func (s *ginServer) helloWorld() {
	s.router.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, domain.Response{
			Message: "Hello World",
		})
	})
}
