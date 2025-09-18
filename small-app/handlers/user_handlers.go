package handlers

import (
	"net/http"
	"small-app/internal/users"

	"github.com/gin-gonic/gin"
)

func (h handler) Signup(c *gin.Context) {

	var n users.NewUser
	err := c.ShouldBindJSON(&n)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	u, err := h.uc.CreatUser(n)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, u)
}

func (h handler) ViewAll(c *gin.Context) {
	u := h.uc.FetchUsers()
	c.JSON(http.StatusOK, u)
}
