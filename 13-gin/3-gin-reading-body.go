package main

/*
	{
	  "id": 1,
	  "name": "Alice Johnson",
	  "email": "alice.johnson@example.com",
	  "age": 28,
	  "created_at": "2025-09-13T10:30:00Z"  // time.Time
	}
*/

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

	// read the body
	//parse the body and store json it in a struct
	// c.ShouldBindJSON(&struct)
}

func createUser(c *gin.Context) {
	var user user
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	fmt.Printf("Received user: %+v\n", user)
	c.JSON(200, user)
}
