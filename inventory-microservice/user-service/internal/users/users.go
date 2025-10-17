package users

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Conf struct {
	db *sql.DB
}

func NewConf(db *sql.DB) (*Conf, error) {
	if db == nil {
		return nil, errors.New("db is nil")
	}
	return &Conf{db: db}, nil
}

func (c *Conf) InsertUser(ctx context.Context, newUser NewUser) (User, error) {
	uid := uuid.NewString()

	// Hash the user's password using bcrypt to store it securely in the database.
	// `bcrypt.DefaultCost` determines the cost of the hashing algorithm for computational overhead.
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		return User{}, err // Return an error if password hashing fails.
	}

	// Get the current UTC time for `createdAt` and `updatedAt` timestamps for the new user.
	createdAt := time.Now().UTC()
	updatedAt := time.Now().UTC()

	var user User

	f := func(tx *sql.Tx) error {
		query := `
      INSERT INTO users
      (id, name, email, password_hash, created_at, updated_at, roles)
      VALUES ($1, $2, $3, $4, $5, $6, $7)
      RETURNING id, name, email, created_at, updated_at, roles
      `

		// Execute the `INSERT` query within the transaction to add the new user.
		// `QueryRowContext` executes the query and scans the resulting row into the `user` struct.
		err = tx.QueryRowContext(ctx, query, uid, newUser.Name, newUser.Email, hashedPassword,
			createdAt, updatedAt, newUser.Roles).
			Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt, &user.Roles)

		if err != nil {
			return fmt.Errorf("failed to insert user: %w", err)
		}

		// If the query is successful, return nil to indicate no errors.
		return nil
	}
	err = c.withTx(ctx, f)
	// If the transaction or insertion fails, return an error.
	if err != nil {
		return User{}, fmt.Errorf("failed to insert user: %w", err)
	}
	return user, nil

}

func (c *Conf) AuthenticateUser(ctx context.Context, email string, password string) (User, error) {
	var user User
	f := func(tx *sql.Tx) error {
		query := `
SELECT id, name, email, password_hash, created_at, updated_at, roles
		FROM users
		WHERE email = $1
`
		var passwordHash string

		err := tx.QueryRowContext(ctx, query, email).
			Scan(&user.ID, &user.Name, &user.Email, &passwordHash,
				&user.CreatedAt, &user.UpdatedAt, &user.Roles)

		if err != nil {
			return fmt.Errorf("failed to fetch user: %w", err)
		}

		// Compare the stored hashed password with the provided password
		err = bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
		if err != nil {
			return fmt.Errorf("invalid password: %w", err)
		}
		return nil

	}

	err := c.withTx(ctx, f)
	if err != nil {
		return User{}, fmt.Errorf("failed to authenticate user: %w", err)
	}
	return user, nil
}

// withTx is a helper function that simplifies the usage of SQL transactions.
// It begins a transaction using the provided context (`ctx`), executes the given function (`fn`),
// and handles commit or rollback based on the success or failure of the function.
func (c *Conf) withTx(ctx context.Context, fn func(*sql.Tx) error) error {
	// Start a new transaction using the context.
	tx, err := c.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to begin tx: %w", err) // Return an error if the transaction cannot be started.
	}

	// Execute the provided function (`fn`) within the transaction.
	err = fn(tx)
	if err != nil {
		// If the function returns an error, attempt to roll back the transaction.
		er := tx.Rollback()
		if er != nil {
			// If rollback also fails (and it's not because the transaction is already done),
			// return an error indicating the failure to roll back.
			return fmt.Errorf("failed to rollback withTx: %w", err)
		}
		// Return the original error from the function execution.
		return fmt.Errorf("failed to execute withTx: %w", err)
	}

	// If no errors occur, commit the transaction to apply the changes.
	err = tx.Commit()
	if err != nil {
		// Return an error if the transaction commit fails.
		return fmt.Errorf("failed to commit withTx: %w", err)
	}

	// Return nil if the function executes successfully and the transaction is committed.
	return nil
}
