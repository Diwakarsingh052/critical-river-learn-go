package handlers

import (
	"small-app/internal/users"

	"github.com/gin-gonic/gin"
)

/*


`
curl -X POST http://localhost:8080/signup \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "email": "john.doe@example.com",
    "age": 25,
    "password": "mySecurePassword123"
  }'

curl localhost:3000/view
`


*/

// we created the handler struct to inject dependencies
// handler package depends on users package
// we can't call user package methods from handler package
// so we added a field to the handler struct that holds the users.Conn
// now using this field we can call users.Conn methods from handler package

// anytime a package depends on another package,
// we should create a struct that holds the dependencies
// some companies use framework like fX for dependency injection
type handler struct {
	uc users.Conn
}

func InitRoutes(con users.Conn) *gin.Engine {
	r := gin.New()
	h := handler{uc: con}
	r.POST("/signup", h.Signup)
	r.GET("/view", h.ViewAll)
	return r
}
