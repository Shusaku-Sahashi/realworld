package router

import "github.com/gin-gonic/gin"

/*
InitRouter is initializer of http router.
*/
func InitRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	return r
}
