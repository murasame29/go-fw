package gin

import (
	"murasame29/go-fw/src/domain"

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
		c.JSON(200, domain.Response{
			Message: "Hello World",
		})
	})
}
