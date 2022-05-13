package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"travesty/internal/app/model"
)

func StartAdminServer(host string, port int, services []model.Service) error {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, services)
	})
	return r.Run(fmt.Sprintf("%s:%d", host, port))
}
