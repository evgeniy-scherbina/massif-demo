package main

import "math/rand"

const (
	KB = 1000
	MB = KB * 1000
	GB = MB * 1000
)

func allocate(bytes int) []byte {
	x := make([]byte, bytes)
	for i := 0; i < bytes; i++ {
		x[i] = byte(rand.Int())
	}
	return x
}

func main() {
	// allocate 100KB
	n := 100
	x := make([][]byte, n)
	for i := 0; i < n; i++ {
		x[i] = allocate(KB)
	}

	// allocate 1MB
	mb := allocate(MB)
	_ = mb

	// allocate 100KB
	n = 100
	y := make([][]byte, n)
	for i := 0; i < n; i++ {
		y[i] = allocate(KB)
	}
}
