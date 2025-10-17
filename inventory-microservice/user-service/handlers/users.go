package handlers

import (
	"errors"
	"log/slog"
	"net/http"
	"user-service/internal/users"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	conf *users.Conf
}

func NewHandler(conf *users.Conf) (*Handler, error) {
	if conf == nil {
		return nil, errors.New("conf is nil")
	}
	return &Handler{conf: conf}, nil

}

func (h *Handler) SignUp(c *gin.Context) {

	var newUser users.NewUser

	// Parse the incoming JSON request into a `NewUser` struct.
	err := c.ShouldBindJSON(&newUser)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": http.StatusText(http.StatusBadRequest),
		})

		slog.Error("Error in parsing request body:",
			slog.String("error", err.Error()))
		return
	}

	v := validator.New()
	err = v.Struct(newUser)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "please provide values in correct format",
		})
		return
	}
	ctx := c.Request.Context()
	user, err := h.conf.InsertUser(ctx, newUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "User Creation Failed",
		})
		return
	}
	c.JSON(http.StatusCreated, user)
}
