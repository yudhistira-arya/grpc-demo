package main

import (
	"sync"
	"testing"
)

var result []Meteorite

func Benchmark1HttpRequest(b *testing.B) {
	var response []Meteorite
	for n := 0; n < b.N; n++ {
		// always record the result to prevent the compiler eliminating the function call.
		response = GetMeteorite()
	}
	// always store the result to a package level variable
	// so the compiler cannot eliminate the Benchmark itself.
	result = response
}

func Benchmark10HttpRequest(b *testing.B) {
	concurrentRequest := 10

	var response []Meteorite
	for n := 0; n < b.N; n++ {
		var wg sync.WaitGroup
		wg.Add(concurrentRequest)
		for c := 0; c < concurrentRequest; c++ {
			go func() {
				defer wg.Done()
				// always record the result to prevent the compiler eliminating the function call.
				response = GetMeteorite()
			}()
		}
		wg.Wait()
	}
	// always store the result to a package level variable
	// so the compiler cannot eliminate the Benchmark itself.
	result = response
}

func Benchmark100HttpRequest(b *testing.B) {
	concurrentRequest := 10

	var response []Meteorite
	for n := 0; n < b.N; n++ {
		var wg sync.WaitGroup
		wg.Add(concurrentRequest)
		for c := 0; c < concurrentRequest; c++ {
			go func() {
				defer wg.Done()
				// always record the result to prevent the compiler eliminating the function call.
				response = GetMeteorite()
			}()
		}
		wg.Wait()
	}
	// always store the result to a package level variable
	// so the compiler cannot eliminate the Benchmark itself.
	result = response
}


