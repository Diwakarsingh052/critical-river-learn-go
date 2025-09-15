package users

import (
	"sync"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type CacheStore map[string]User

type Conn struct {
	cache CacheStore
	mu    *sync.RWMutex
}

func NewConn() Conn {
	return Conn{
		cache: make(CacheStore, 100),
		mu:    new(sync.RWMutex),
	}
}

func (c *Conn) CreatUser(n NewUser) (User, error) {
	//creating hash of password
	// from hash output we can't get the original value
	// safe for passwords
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(n.Password), bcrypt.DefaultCost)
	if err != nil {
		return User{}, err
	}

	u := User{
		Id:           uuid.NewString(),
		Email:        n.Email,
		Name:         n.Name,
		Age:          n.Age,
		PasswordHash: string(hashedPass),
	}

	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache[u.Email] = u

	return u, nil

}

func (c *Conn) FetchUsers() map[string]User {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.cache
}
