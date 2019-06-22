package router

import "github.com/gin-gonic/gin"

/*
InitRouter is initializer of http router.
*/
func InitRouter() *gin.Engine {
	r := gin.Default()

	return r
}
