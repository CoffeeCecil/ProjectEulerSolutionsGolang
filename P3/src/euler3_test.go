package main

import(
	"testing"
	"math/big"
)


func Test_findBiggestFactorPrimes(t *testing.T){
	i:=findBiggestFactorPrimes(600851475143)
	k:=big.NewInt(i)

	if !k.ProbablyPrime(100){
		panic("K probably isn't prime!")
	}

}

func Test_findBiggestFactorPrimes4(t *testing.T){
	i:= fbEAFP(600851475143)
	t.Log(i)
}

func Test_findBiggestFactorPrimes2(t *testing.T){
	v:=int64(600851475143)
	i:=fBFMR644(v)
	k:=big.NewInt(i)

	if !k.ProbablyPrime(100){
		panic("K probably isn't prime!")
	}
	j:=big.NewInt(v)
	j.Div(j, k)
	if j.ProbablyPrime(100) && j.Cmp(k) != 0{
		panic("J is probably a prime number!")
	}
	t.Log(i)
}


func Test_findBiggestFactorPrimes3(t *testing.T){
	v:=int64(600851475143)
	i:=fBFMR647(v)
	k:=big.NewInt(i)

	if !k.ProbablyPrime(100){
		t.Log(k)
		panic("K probably isn't prime!")
	}
	j:=big.NewInt(v)
	j.Div(j, k)
	if j.ProbablyPrime(100) && j.Cmp(k) != 0{
		panic("J is probably a prime number!")
	}
	t.Log(i)
}

func Test_findBiggestFactorPrimesSmall(t *testing.T){
	t.Log(findBiggestFactorPrimes(13195))
}


func Test_findBiggestFactorPrimesSmall2(t *testing.T){
	t.Log(fBFMR642(13195))
}

func Benchmark_findBiggestFactorPrimesSmall(b *testing.B){
	for i:=0; i < b.N; i++{
		findBiggestFactorPrimes(13195)
	}
}

func Benchmark_findBiggestFactorPrimesSmallB(b *testing.B){
	for i:=0; i < b.N; i++{
		findBiggestFactorPrimes(131950)
	}
}

func Benchmark_findBiggestFactorPrimesSmallC(b *testing.B){
	for i:=0; i < b.N; i++{
		findBiggestFactorPrimes(1319500)
	}
}

func Benchmark_findBiggestFactorPrimesSmallD(b *testing.B){
	for i:=0; i < b.N; i++{
		fBFMR32(1319500)
	}
}


func Benchmark_findBiggestFactorPrimesSmallE(b *testing.B){
	for i:=0; i < b.N; i++{
		fBFMR642(1319500)
	}
}

func Benchmark_findBiggestFactorPrimesSmallF(b *testing.B){
	for i:=0 ; i < b.N; i++ {
		fBFMR643(1319500)
	}
}

func Benchmark_findBiggestFactorPrimesSmallG(b *testing.B){
	
	for i:=0 ; i < b.N; i++ {
		fBFMR644(1319500)
	}
}

func Benchmark_findBiggestFactorPrimesSmallH(b *testing.B){
	for i:=0 ; i < b.N; i++ {
		fBFMR645(1319500)
	}
}

func Benchmark_findBiggestFactorPrimesSmallI(b *testing.B){
	for i:=0 ; i < b.N; i++ {
		fBFMR646(1319500)
	}
}

func Benchmark_findBiggestFactorPrimes(b *testing.B){
	v:=int64(600851475143)
	for i:=0 ; i < b.N; i++ {
		fBFMR643(v)
	}
}

func Benchmark_findBiggestFactorPrimesB(b *testing.B){
	v:=int64(600851475143)
	for i:=0 ; i < b.N; i++ {
		fBFMR647(v)
	}
}


func Benchmark_findBiggestFactorPrimesC(b *testing.B){
	v:=int64(600851475143)
	for i:=0 ; i < b.N; i++ {
		fbEAFP(v)
	}
}