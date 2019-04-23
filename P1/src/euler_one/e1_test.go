package euler_one

import "testing"

func Test_e1(t *testing.T) {
	m, k := euler_one_better(1000)
	t.Log("Result ", m," - ", k)
}

func Test_e1f(t *testing.T) {
	m, k := euler_one_fastest(999)
	t.Log("Result ", m," - ", k)
}

func Test_e1n(t *testing.T) {
	m, k := euler_one_naive(1000)
	t.Log("Result ", m," - ", k)
}

func Benchmark_e1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = euler_one_better(999)
	}
}

func Benchmark_e1f(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = euler_one_fastest(999)
	}
}

func Benchmark_e1n(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = euler_one_naive(999)
	}
}
