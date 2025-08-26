package main

import "strconv"

// err must be the last value to be returned
// use types in return instead of variable names
func convertToFloat(s string) (float64, error) {

	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, err // whenever err happens set other values to default
	}

	// success case should be clearly visible
	return f, nil // don't write err , even err is going to be nil
}
