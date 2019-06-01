package main

import "testing"

func Test_e1(t *testing.T) {
	m, k := eulerOneBetter(1000)
	t.Log("Result ", m," - ", k)
}

func Test_e1f(t *testing.T) {
	m, k := eulerOneFastest(999)
	t.Log("Result ", m," - ", k)
}

func Test_e1n(t *testing.T) {
	m, k := eulerOneNaive(1000)
	t.Log("Result ", m," - ", k)
}

func Benchmark_e1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = eulerOneBetter(999)
	}
}

func Benchmark_e1f(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = eulerOneFastest(999)
	}
}

func Benchmark_e1n(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = eulerOneNaive(999)
	}
}
