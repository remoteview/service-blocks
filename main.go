package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/remoteview/service-blocks/blocks"
	"github.com/remoteview/service-blocks/version"
)

var dbBlocks []blocks.Block

// HealthCheck - status
type HealthCheck struct {
	Version string `json:"version"`
	Status  string `json:"status"`
}

func main() {
	dbBlocks = append(dbBlocks, blocks.Block{ID: "1", StartTime: blocks.JSONTime(time.Date(0, 0, 0, 9, 0, 0, 0, time.UTC)), EndTime: blocks.JSONTime(time.Date(0, 0, 0, 17, 0, 0, 0, time.UTC)), Timezone: "UTC", UserID: "1"})
	dbBlocks = append(dbBlocks, blocks.Block{ID: "2", StartTime: blocks.JSONTime(time.Date(0, 0, 0, 9, 0, 0, 0, time.UTC)), EndTime: blocks.JSONTime(time.Date(0, 0, 0, 17, 0, 0, 0, time.UTC)), Timezone: "UTC", UserID: "2"})
	dbBlocks = append(dbBlocks, blocks.Block{ID: "3", StartTime: blocks.JSONTime(time.Date(0, 0, 0, 9, 0, 0, 0, time.UTC)), EndTime: blocks.JSONTime(time.Date(0, 0, 0, 17, 0, 0, 0, time.UTC)), Timezone: "UTC", UserID: "3"})

	r := gin.Default()
	r.GET("/_health", healthCheckHandler)
	r.GET("/blocks", listBlocksHandler)
	r.GET("/blocks/:id", listBlockHandler)

	r.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "OPTIONS"},
		AllowHeaders:     []string{"Origin"},
		AllowAllOrigins:  true,
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	fmt.Println("Running server!")
	err := r.Run(":3001")
	if err != nil {
		fmt.Println("Error starting server")
	}
}

func healthCheckHandler(c *gin.Context) {
	version, err := version.GetVersion()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, HealthCheck{Status: "Ok", Version: version})
}

func listBlocksHandler(c *gin.Context) {
	c.JSON(200, dbBlocks)
}

func listBlockHandler(c *gin.Context) {
	for _, item := range dbBlocks {
		if item.ID == c.Param("id") {
			c.JSON(http.StatusBadRequest, item)
			return
		}
	}
}
