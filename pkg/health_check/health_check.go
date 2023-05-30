package health_check

import "github.com/gin-gonic/gin"

func RegisterHTTPEEndpoints(router *gin.Engine) {
	router.GET("/status", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "OK",
		})
	})
}
