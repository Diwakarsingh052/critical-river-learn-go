package sum

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

// create file with _test.go to register them as test file

// go test -v ./... // run all the tests for the project

// test function must be prefixed with Test
// and it must have a pointer to testing.T as its first argument
func TestSumInt(t *testing.T) {
	//1st test case
	input := []int{1, 2, 3, 4, 5}
	want := 15
	got := SumInt(input)

	if got != want {
		// test would continue on if test case fail in case of Errorf
		t.Errorf("sum of 1 to 5 should be %d, go %d", want, got)

		// test would stop on if test case fail in case of Fatalf
		//t.Fatalf()
	}

	//2nd test case
	input = nil
	want = 0
	got = SumInt(input)
	if got != want {
		t.Errorf("sum of nil should be %d, got %d", want, got)
	}

}

func TestSumIntV2(t *testing.T) {

	// table test
	// ... would be an array of struct
	// size of the array is the number of test cases
	tt := [...]struct {
		name  string
		input []int
		want  int
	}{
		{
			name:  "one to five numbers",
			input: []int{1, 2, 3, 4, 5},
			want:  15,
		},
		{
			name:  "nil slice",
			input: nil,
			want:  0,
		},
		{
			name:  "empty slice",
			input: []int{},
			want:  0,
		},
	}
	fmt.Printf("%T\n", tt)

	for _, tc := range tt {
		// t.Run creates a subtest, you can run each index of slice as a subtest
		t.Run(tc.name, func(t *testing.T) {
			got := SumInt(tc.input)
			require.Equal(t, tc.want, got)
		})

	}
}
