package handler

import (
	"github.com/gin-gonic/gin"
)

/*
	各Handlerをrouteに紐づけていく
*/
func (h *Handler) Register(r *gin.Engine) *gin.Engine {
	user := r.Group("/users")
	user.POST("/login", func(c *gin.Context) { h.Login(c) })

	return r
}
