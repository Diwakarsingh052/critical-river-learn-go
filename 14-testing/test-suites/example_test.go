package test_suites

import (
	"database/sql"
	"fmt"
	"testing"

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

// SetupSuite is called once before running the tests in the suite.
// Use this method for expensive setup operations that should only happen once,
// such as establishing database connections or loading test fixtures.
func (s *ExampleTestSuite) SetupSuite() {
	fmt.Println("SetupSuite: Called once before all tests in the suite")
	// Example operations you might perform in SetupSuite:
	// - Connect to a database
	// - Create test fixtures
	// - Initialize external services
	//s.db = "some connection"
}

// TearDownSuite is called once after all tests in the suite have been run.
// Use this method for cleanup operations that should happen once at the end,
// such as closing database connections or removing test fixtures.
func (s *ExampleTestSuite) TearDownSuite() {
	fmt.Println("TearDownSuite: Called once after all tests in the suite")
	// Example operations you might perform in TearDownSuite:
	// - Disconnect from a database
	// - Remove test fixtures
	// - Shut down external services
}

// SetupTest is called before each test in the suite.
// Use this method for setup operations that should happen before every test,
// such as initializing test data or resetting state.
func (s *ExampleTestSuite) SetupTest() {
	fmt.Println("SetupTest: Called before each test")
	s.TestData = "test data initialized"
	// Example operations you might perform in SetupTest:
	// - Reset test data to a known state
	// - Create test objects
	// - Set up mock expectations
}

// TearDownTest is called after each test in the suite.
// Use this method for cleanup operations that should happen after every test,
// such as clearing test data or resetting state.
func (s *ExampleTestSuite) TearDownTest() {
	fmt.Println("TearDownTest: Called after each test")
	s.TestData = ""
	// Example operations you might perform in TearDownTest:
	// - Clean up test data
	// - Verify and reset mocks
	// - Close any resources opened specifically for the test
}

// TestExample1 is a test function within the suite.
// Any method that starts with "Test" will be run as a test.
func (s *ExampleTestSuite) TestExample1() {
	fmt.Println("Running TestExample1")
	// The s.Equal method comes from the embedded suite.Suite
	// It asserts that the two values are equal
	s.Equal("test data initialized", s.TestData, "TestData should be initialized")

	// Other assertion examples:
	// s.NotEqual(expected, actual, "message")
	// s.True(value, "message")
	// s.False(value, "message")
	// s.Nil(value, "message")
	// s.NotNil(value, "message")
}
func (s *ExampleTestSuite) TestExample2() {
	fmt.Println("Running TestExample2")
	s.Equal("test data initialized", s.TestData, "TestData should be initialized")
}

func TestExampleSuite(t *testing.T) {
	// Create a new instance of our test suite
	testSuite := new(ExampleTestSuite)

	suite.Run(t, testSuite)
	// The execution order will be:
	// 1. SetupSuite()
	// 2. For each test method:
	//    a. SetupTest()
	//    b. Test method (e.g., TestExample1)
	//    c. TearDownTest()
	// 3. TearDownSuite()

}
