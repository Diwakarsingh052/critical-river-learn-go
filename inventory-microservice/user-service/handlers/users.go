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
		slog.Error("Error parsing request body", slog.String("error", err.Error()))
		return
	}

	v := validator.New()
	err = v.Struct(newUser)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "please provide values in correct format",
		})

		slog.Error("Error validating fields of user",
			slog.String("error", err.Error()))
		return

	}
	ctx := c.Request.Context()
	user, err := h.conf.InsertUser(ctx, newUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "User Creation Failed",
		})

		slog.Error("Error inserting user", slog.String("error", err.Error()))
		return
	}
	c.JSON(http.StatusCreated, user)
}

func (h *Handler) Login(c *gin.Context) {
	// Declare a struct to hold the login request payload
	var loginPayload struct {
		Email    string `json:"email" validate:"required,email"` // Email must be valid and required
		Password string `json:"password" validate:"required"`    // Password required
	}
	// converting json to struct
	err := c.ShouldBindJSON(&loginPayload)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": http.StatusText(http.StatusBadRequest),
		})
		slog.Error("Error parsing request body", slog.String("error", err.Error()))
		return
	}

	v := validator.New()
	// validate the struct
	err = v.Struct(loginPayload)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "please provide values in correct format",
		})
		slog.Error("Error validating fields of user",
			slog.String("error", err.Error()))
		return
	}

	user, err := h.conf.AuthenticateUser(c.Request.Context(), loginPayload.Email, loginPayload.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "User Authentication Failed",
		})
		slog.Error("Error authenticating user", slog.String("error", err.Error()))
		return
	}

	c.JSON(http.StatusOK, user)
}
