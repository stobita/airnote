package server

import "github.com/gin-gonic/gin"

// Run api server
func Run() error {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "airnote",
		})
	})
	return r.Run()
}
