package handler

import (
	"fmt"
	"net/http"

	"github.com/app/realworld/handler/resource"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Login(c *gin.Context) {
	var req resource.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Print(req)
	userInfo, token, err := h.Auth.Authenticate(req.Email(), req.Password())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resource.NewLoginResponse(userInfo, token))
}
