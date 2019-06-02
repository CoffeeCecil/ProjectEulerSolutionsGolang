package main
import ("testing"
	"math/big"
)

func TestEuler2(t *testing.T){
	t.Log(Fib_naive(big.NewInt(4000000)))
	
}

func TestEuler2Fast(t *testing.T){
	t.Log(Fib_even(big.NewInt(4000000)))
}
func TestEuler2Faster(t *testing.T){
	t.Log(Fib_even_faster(big.NewInt(4000000)))
}

func BenchmarkEuler2(b *testing.B){
	for i:=0; i< b.N; i++{
		Fib_naive(big.NewInt(4000000))
	}
}
//need a big number to show how the speedup effects
//the algorithm
func BenchmarkEuler2FastBig(b *testing.B){
	for i:=0; i< b.N; i++{
		Fib_even(big.NewInt(160000000))
	}
}

func BenchmarkEuler2FastestBig(b *testing.B){
	for i:=0; i< b.N; i++{

		Fib_even_faster(big.NewInt(160000000))
	}
}

func BenchmarkEuler2Fast(b *testing.B){
	for i:=0; i< b.N; i++{
		Fib_even(big.NewInt(4000000))
	}
}

func BenchmarkEuler2Fastest(b *testing.B){
	for i:=0; i< b.N; i++{

		Fib_even_faster(big.NewInt(4000000))
	}
}