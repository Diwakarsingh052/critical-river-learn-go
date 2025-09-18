package main

import "fmt"

// UniqueSubstrings Function  accepts a string and return its unique substrings
func UniqueSubstrings(s string) []string {
	var result []string //Initialize an empty list to hold unique substrings
	n := len(s)         //Calculate length of input string

	//Iterate over entire string
	for i := 0; i < n; i++ {

		// For each character in the string, iterate over the rest of the string
		for j := i + 1; j <= n; j++ {

			// Get the substring from i to j
			sub := s[i:j]

			// Flag to check if a substring already exists in the result
			exists := false

			// Iterate over the result slice to see if the substring already exists.
			for _, elem := range result {
				if sub == elem {
					exists = true
					break
				}
			}

			// If the substring didn't exist in the result, append it to result slice.
			if !exists {
				result = append(result, sub)
			}
		}
	}
	return result //Return the result slice containing unique substrings.
}

// OptimizedUniqueSubstrings Function accepts a string and optimises the way we find unique substrings by using a map
func OptimizedUniqueSubstrings(s string) []string {
	result := make([]string, 0) // Initialize an empty slice to hold unique substrings

	// Initialize a map to hold unique substrings and give it an initial size of 1000.
	substringMap := make(map[string]bool, 1000)

	n := len(s) // Calculate the length of the input string

	// Iterate over each character of the input string
	// Iterate over each character of the input string
	for i := 0; i < n; i++ {

		// For each character in the string, iterate over the rest of the string
		for j := i + 1; j <= n; j++ {

			// Get the substring from i to j
			sub := s[i:j]

			// If the substring doesn't exist in substringMap (it's unique until now)
			if _, exists := substringMap[sub]; !exists {

				// Add the substring to substringMap and set its value to true
				substringMap[sub] = true

				// Append the unique substring to the result slice.
				result = append(result, sub)
			}
		}
	}

	// Return the result slice containing unique substrings
	return result
}

func main() {
	s := "hello"
	result := UniqueSubstrings(s)
	fmt.Println("Unique substrings: ", result)
	//result = OptimizedUniqueSubstrings(s)
	fmt.Println("Unique substrings: ", result)
}
