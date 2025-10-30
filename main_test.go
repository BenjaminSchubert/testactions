package main_test

import "testing"

func BenchmarkSomething(b *testing.B) {
	x := 0
	for b.Loop() {
		x += 1
		b.Log("Hello")
	}
}
