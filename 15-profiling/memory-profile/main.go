package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type TextAnalysis struct {
	Content       string
	WordFrequency map[string]int
	AvgWordLength float64
	LongestWord   string
}

func AnalyzeText(filename string) (TextAnalysis, error) {
	// reading the entire file in memory
	content, err := os.ReadFile(filename)
	if err != nil {
		return TextAnalysis{}, err
	}

	// Convert the byte slice of the file content to a string.
	text := string(content)

	// Split the `text` string into a slice of words.
	words := strings.Fields(text)

	// Initialize a map to hold the frequency of occurrence of each word.
	frequency := make(map[string]int)

	// Initialize totalLength to cumulate the length of all words.
	totalLength := 0

	// Initialize longest to hold the longest word.
	longest := ""

	// Loop over every word in the `words` slice.
	for _, word := range words {
		frequency[word]++ // Increment the count for the word in the frequency map.
		if len(word) > len(longest) {
			longest = word // Update the longest word if the current word is longer.
		}
		totalLength += len(word) // Add the length of the current word to totalLength.
	}

	// Calculating the average word length.
	avgLength := float64(totalLength) / float64(len(words))

	return TextAnalysis{
		Content:       text,
		WordFrequency: frequency,
		AvgWordLength: avgLength,
		LongestWord:   longest,
	}, nil

}

var cache = make(map[string]TextAnalysis)

func OptimizedAnalyzeText(fileName string) (TextAnalysis, error) {
	if res, ok := cache[fileName]; ok {
		return res, nil
	}

	file, err := os.Open(fileName)
	if err != nil {
		return TextAnalysis{}, err
	}
	defer file.Close()

	frequency := make(map[string]int, 100000)
	totalLength := 0
	longest := ""

	scanner := bufio.NewScanner(file)
	// scanner by default splits by line
	// changing it to split by words, it would read one word at a time
	scanner.Split(bufio.ScanWords)

	// loop over the scanner
	// it keeps reading the file until it reaches EOF
	for scanner.Scan() {
		word := scanner.Text()
		frequency[word]++
		if len(word) > len(longest) {
			longest = word
		}
		totalLength += len(word)
	}

	if scanner.Err() != nil {
		return TextAnalysis{}, scanner.Err()
	}

	avgLength := float64(totalLength) / float64(len(frequency))
	textAnalysis := TextAnalysis{
		Content:       "",
		WordFrequency: frequency,
		AvgWordLength: avgLength,
		LongestWord:   longest,
	}
	cache[fileName] = textAnalysis
	return textAnalysis, nil

}

func main() {
	fileName := "book.txt"

	analysis, err := AnalyzeText(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Average Word Length:", analysis.AvgWordLength)
	fmt.Println("Longest Word:", analysis.LongestWord)

}
