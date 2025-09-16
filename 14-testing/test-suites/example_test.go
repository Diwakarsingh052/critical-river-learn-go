package test_suites

import (
	"database/sql"

	"github.com/stretchr/testify/suite"
)

type ExampleTestSuite struct {
	// suite.Suite must be embedded to provide the functionality of a testify suite
	suite.Suite
	// dependencies for the test
	db       *sql.DB
	TestData string
}

// SUITE LIFECYCLE METHODS
// ============================================================================
// Testify provides four lifecycle methods that you can implement:
// 1. SetupSuite - runs once in the suite
// 2. TearDownSuite - runs once after all tests in the suite have finished
// 3. SetupTest - runs before each test in the suite
// 4. TearDownTest - runs after each test in the suite
