package gin

import (
	"murasame29/go-fw/src/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ginServer struct {
	router *gin.Engine
}

func NewGinServer() *gin.Engine {
	server := &ginServer{
		router: gin.Default(),
	}
	server.helloWorld()

	return server.router
}

func (s *ginServer) helloWorld() {
	s.router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, domain.Response{
			Message: "Hello World",
		})
	})
}
