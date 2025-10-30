package main_test

import "testing"

func BenchmarkSomething(b *testing.B) {
	x := 0
	y := 1
	for b.Loop() {
		x += 1
		y *= 2
		b.Log("Hello")
	}
}
