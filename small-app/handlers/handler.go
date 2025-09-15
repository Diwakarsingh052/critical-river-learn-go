package handlers

import (
	"net/http"
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

func (h handler) Signup(c *gin.Context) {
	// Conn // CreateUser
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
