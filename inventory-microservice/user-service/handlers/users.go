package handlers

import (
	"errors"
	"log/slog"
	"net/http"
	"time"
	"user-service/internal/auth"
	"user-service/internal/users"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
)

/*
signup request json

	{
	  "name": "Diwakar",
	  "email": "diwakar@example.com",
	  "password": "securepassword",
	  "roles": ["user"]
	}
*/
/*
login request json
{
  "email": "diwakar@example.com",
  "password": "securepassword"
}
*/

type Handler struct {
	conf *users.Conf
	k    *auth.Keys
}

func NewHandler(conf *users.Conf, k *auth.Keys) (*Handler, error) {
	if conf == nil || k == nil {
		return nil, errors.New("conf or auth is nil")
	}
	return &Handler{conf: conf, k: k}, nil

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
	var claim auth.Claims
	claim.RegisteredClaims = jwt.RegisteredClaims{
		Issuer:    "user-service",
		Subject:   user.ID,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}
	claim.Roles = user.Roles
	token, err := h.k.GenerateToken(claim)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "problem in logging in"})
		slog.Error("Error generating token", slog.String("error", err.Error()))
		return
	}
	c.JSON(http.StatusOK, token)
}
