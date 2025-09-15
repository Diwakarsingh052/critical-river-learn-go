package users

type User struct {
	Id           string
	Email        string `json:"email"`
	Name         string `json:"name"` // giving the name of the field in the json ouptut
	Age          int64  `json:"age"`
	PasswordHash string `json:"password_hash"`
}

// NewUser is the struct that we use to create a new user
// we should avoid using domain struct directly to accept client json data
type NewUser struct {
	Name     string `json:"name" binding:"required,max=60"`
	Email    string `json:"email" binding:"required,email,max=60"`
	Age      int64  `json:"age" binding:"required,min=18"`
	Password string `json:"password" binding:"required,max=60"`
}
