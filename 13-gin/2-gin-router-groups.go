package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	// any logger registered using r.Use() will be applied to all routes, and all groups
	r.Use(gin.Logger())

	v1 := r.Group("/v1")
	{
		// v1.Use would apply middleware to all routes in this group written after it
		v1.Use(SpecialMiddleware())
		v1.GET("/users", func(c *gin.Context) {
			//panic("something went wrong")
			c.String(200, "Users v1 (Gin)")
		})

		// v1.Use(auth) would be applied to endpoint written after it not before
		// which means user would not have the auth middleware attached
		v1.Use(Auth())
		v1.GET("/posts", func(c *gin.Context) {
			c.String(200, "Posts v1 (Gin)")
		})
	}

	v2 := r.Group("/v2")
	{
		v2.GET("/users", func(c *gin.Context) {
			//panic("something went wrong")
			c.String(200, "Users v2 (Gin)")
		})

		v2.GET("/posts", func(c *gin.Context) {
			c.String(200, "Posts v2 (Gin)")
		})
	}
	r.Run(":8080")
}

func SpecialMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Time before request
		t := time.Now()
		log.Println("Special middleware started")
		c.Next()

		// After request
		latency := time.Since(t)
		log.Printf("%d | %s | %s | %s\n",
			c.Writer.Status(), c.Request.Method, c.Request.URL, latency)
	}

}

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("auth mid layer started")
		c.Next()
	}
}
