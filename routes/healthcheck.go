package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/remoteview/service-blocks/healthcheck"
	"github.com/remoteview/service-blocks/version"
)

// HealthCheckHandler - route
// @Summary Health check
// @Description Health check
// @Accept  json
// @Produce  json
// @Router /_health [get]
func HealthCheckHandler(c *gin.Context) {
	version := version.GetVersion()
	c.JSON(200, healthcheck.HealthCheck{Status: "Ok", Version: version})
}
