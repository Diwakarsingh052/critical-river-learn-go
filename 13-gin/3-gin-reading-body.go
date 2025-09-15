package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

/*
	{
	  "id": 1,
	  "name": "Alice Johnson",
	  "email": "alice.johnson@example.com",
	  "age": 28,
	  "created_at": "2025-09-13T10:30:00Z"  // time.Time
	}
*/

type user struct {
	UserId    int       `json:"id" binding:"required"`
	Name      string    `json:"name"  binding:"required,min=3,max=50"`
	Email     string    `json:"email" binding:"required,email"`
	Age       int       `json:"age" binding:"required,gte=18,lte=120"`
	CreatedAt time.Time `json:"created_at"`
}

/*
curl -X POST "http://localhost:8080/users" \
  -H "Content-Type: application/json" \
  -d '{
    "id": 1,
    "name": "Alice Johnson",
    "email": "alice.johnson@example.com",
    "age": 28,
    "created_at": "2025-09-13T10:30:00Z"
  }'
*/

// create a gin server
// register one post endpoint
// read the body
// parse the body and store json it in a struct

// print the user locally
func main() {
	r := gin.New()
	r.POST("/users", createUser)
	// read the body
	//parse the body and store json it in a struct
	// c.ShouldBindJSON(&struct)
	r.Run(":8080")
}

func createUser(c *gin.Context) {
	var user user
	//ShouldBindJSON(&user)
	// Reading the body
	// Convert json to type
	// validation of the data
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Printf("Received user: %+v\n", user)
	c.JSON(http.StatusOK, user)
}
