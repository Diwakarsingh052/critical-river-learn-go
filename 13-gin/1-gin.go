package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	//Default returns an Engine instance with the Logger and Recovery middleware already attached.
	//r := gin.Default()
	r := gin.New()

	// this would apply middleware to all the routes
	r.Use(gin.Logger(), gin.Recovery())
	r.GET("/hello", hello)
	r.GET("/name", GetName)

	r.GET("/user", User)
	//r.POST("/user", UserPost)
	//r.PUT("/user", UserPut)

	// creating a new group
	// benefit of this is
	// we can add specific middlewares to a group of routes
	v1 := r.Group("/v1")

	{
		// this would apply middleware to all the routes in this group
		v1.Use(gin.Logger(), gin.Recovery())
		v1.GET("/user", User)   // /v1/user
		v1.GET("/posts", Posts) // /v1/posts
	}

	panic(r.Run(":8080"))

	///user/Posts
	///user/profile
	//
	///products/edit
	///producs/
}
func hello(c *gin.Context) {

	//m := map[string]string{}
	//gin.H // it is a map of type string to any
	// setting the status code
	// setting the content type
	// converting the json
	//sending the response
	c.JSON(http.StatusOK, gin.H{"message": "hello"})
}

func GetName(c *gin.Context) {

	// getting the query params
	// e.g. localhost:8080/search?firstname=John&lastname=Does // anything after ? is query params
	firstName := c.Query("firstname")
	lastName := c.DefaultQuery("lastname", "none")

	c.String(http.StatusOK, "Hello %s %s", firstName, lastName)
}

func User(c *gin.Context) {
	c.String(http.StatusOK, "user")
}

func Posts(c *gin.Context) {
	c.String(http.StatusOK, "posts")
}
