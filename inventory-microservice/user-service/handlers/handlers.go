package handlers

import (
	"user-service/internal/users"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, conf *users.Conf) {
	h, err := NewHandler(conf)
	if err != nil {
		panic(err)
	}
	r.GET("/ping", status)
	endpointPrefix := "/users"
	v1 := r.Group(endpointPrefix)
	{
		v1.POST("/signup", h.SignUp)
		v1.POST("/login", h.Login)
	}

}
func status(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "ok",
	})
}
