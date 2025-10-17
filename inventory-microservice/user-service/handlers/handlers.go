package handlers

import (
	"user-service/internal/auth"
	"user-service/internal/users"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, conf *users.Conf, k *auth.Keys) {
	h, err := NewHandler(conf, k)
	if err != nil {
		panic(err)
	}
	r.GET("/ping", status)
	endpointPrefix := "/users"
	v1 := r.Group(endpointPrefix)
	{
		v1.POST("/signup", h.SignUp)
		v1.POST("/login", h.Login)
		v1.GET("/onlyLoggedin", func(c *gin.Context) {})
	}

}
func status(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "ok",
	})
}
