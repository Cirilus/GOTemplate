package health_check

import "github.com/gin-gonic/gin"

// @Summary Show the status of server.
// @Description get the status of server.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /status [get]
func RegisterHTTPEEndpoints(router *gin.Engine) {
	router.GET("/status", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "OK",
		})
	})
}
