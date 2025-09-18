package main

import "testing"

var txt = `
The Project Gutenberg EBook of Moby Dick; or The Whale, by Herman
Melville

This eBook is for the use of anyone anywhere at no cost and with almost
no restrictions whatsoever.  You may copy it, give it away or re-use it
under the terms of the Project Gutenberg License included with this
eBook or online at www.gutenberg.org

`

func BenchmarkOptimizedUniqueSubstrings(b *testing.B) {

	for b.Loop() {
		OptimizedUniqueSubstrings(txt)
	}
}

func BenchmarkUniqueSubstrings(b *testing.B) {

	for b.Loop() {
		UniqueSubstrings(txt)
	}
}

//go test -run none -bench . -benchtime 3s -benchmem -v -cpuprofile p.out
//go tool pprof -http=:8080 p.out
