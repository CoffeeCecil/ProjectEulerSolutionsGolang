package main

import (
	"./primes/src"
	"math"
	"math/big"
	"fmt"
	"os"
	"strconv"
)

func fBFMR32(n int64) int64 {
	tn := big.NewInt(n)
	cn := big.NewInt(n)
	zero:= big.NewInt(0)
	i := big.NewInt(n/2)
	onei :=big.NewInt(1)
	for ; i.Cmp(onei) > -1; i.Sub(i, onei){
		if primes.MillerRabin32(i){
			tn.Set(cn)
			tn.Mod(cn, i)
			if(tn.Cmp(zero) == 0){
				return i.Int64()
			}
		}
	}
	return n
}

func fBFMR642(n int64) int64 {
	tn := big.NewInt(n)
	cn := big.NewInt(n)
	zero:= big.NewInt(0)
	i := big.NewInt(n/2)
	onei :=big.NewInt(1)
	for ; i.Cmp(onei) > -1; i.Sub(i, onei){
		if primes.MillerRabin642(i){
			tn.Set(cn)
			tn.Mod(cn, i)
			if(tn.Cmp(zero) == 0){
				return i.Int64()
			}
		}
	}
	return n
}

//checks from sqrt of the number, which is the biggest possible factor.
func fBFMR643(n int64) int64 {
	tn := big.NewInt(n)
	cn := big.NewInt(n)
	zero:= big.NewInt(0)
	i := big.NewInt(n)
	i.Sqrt(i)
	onei :=big.NewInt(1)
	i.Add(i, onei)
	for ; i.Cmp(onei) > -1; i.Sub(i, onei){
		if primes.MillerRabin642(i){
			tn.Set(cn)
			tn.Mod(cn, i)
			if(tn.Cmp(zero) == 0){
				return i.Int64()
			}
		}
	}
	return n
}

//checks from sqrt of the number, which is the biggest possible factor.
func fBFMR644(n int64) int64 {
	tn := big.NewInt(n)
	cn := big.NewInt(n)
	zero:= big.NewInt(0)
	i := big.NewInt(n)
	i.Sqrt(i)
	onei :=big.NewInt(1)
	i.Add(i, onei)
	for ; i.Cmp(onei) > -1; i.Sub(i, onei){
		if primes.MillerRabin64(i){
			tn.Set(cn)
			tn.Mod(cn, i)
			if(tn.Cmp(zero) == 0){
				return i.Int64()
			}
		}
	}
	return n
}

//checks from sqrt of the number, which is the biggest possible factor.
func fBFMR645(n int64) int64 {
	tn := big.NewInt(n)
	cn := big.NewInt(n)
	zero:= big.NewInt(0)
	i := big.NewInt(n)
	i.Sqrt(i)
	onei :=big.NewInt(1)
	i.Add(i, onei)
	for ; i.Cmp(onei) > -1; i.Sub(i, onei){
		if primes.MillerRabin643(i){
			tn.Set(cn)
			tn.Mod(cn, i)
			if(tn.Cmp(zero) == 0){
				return i.Int64()
			}
		}
	}
	return n
}


//checks from sqrt of the number, which is the biggest possible factor.
func fBFMR646(n int64) int64 {
	tn := big.NewInt(n)
	cn := big.NewInt(n)
	zero:= big.NewInt(0)
	i := big.NewInt(n)
	i.Sqrt(i)
	onei :=big.NewInt(1)
	i.Add(i, onei)
	for ; i.Cmp(onei) > -1; i.Sub(i, onei){
		if primes.MillerRabin644(i){
			tn.Set(cn)
			tn.Mod(cn, i)
			if(tn.Cmp(zero) == 0){
				return i.Int64()
			}
		}
	}
	return n
}

//checks from sqrt of the number, which is the biggest possible factor.
func fBFMR647(n int64) int64 {
	m := big.NewInt(n)
	tn := big.NewInt(0)
	cn := big.NewInt(n)
	zero:= big.NewInt(0)
	i := big.NewInt(n)
	onei :=big.NewInt(1)
	i.Set(zero)
	
	for ; i.Cmp(cn) < 1; i.Add(i,onei){
		if primes.MillerRabin644(i){
			m.Set(cn)
			m.DivMod(cn, i, tn)
			if(tn.Cmp(zero) == 0){
				if m.Cmp(onei) == 0{
					return cn.Int64()
				}
				cn.Set(m)
			}
		}
	}
	return cn.Int64()
}

//is it worth it to check only primes?
func findBiggestFactorPrimes(n int64) int64{
	tn := big.NewInt(n)
	cn := big.NewInt(n)
	zero:= big.NewInt(0)
	i := big.NewInt(n/2)
	onei :=big.NewInt(1)
	for ; i.Cmp(onei) > -1; i.Sub(i, onei){
		if primes.MillerRabin64(i){
			tn.Set(cn)
			tn.Mod(cn, i)
			if(tn.Cmp(zero) == 0){
				return i.Int64()
			}
		}
	}
	return n
}

//is it worth it to check only primes?
func fbEAFP(n int64) int64{
	var m,tn int
	cn := int(n)
	arr, end := primes.ErasthonesArray(int(math.Sqrt(float64(n))) + 1)
	index:= func(i int) int{
		v := i - 1
		return v >> 1
	}
	//factor out all the powers of two.
	if n%2==0{
		for cn % 2 == 0{
			cn = cn >> 1
		}
		if arr[index(cn)]{
			return int64(cn)
		}
	}
	divmod:= func(n, d int) (int, int) {
		return n/d, n % d
	}
	v:=0
	for i, j:=1,1; j <= end ; i++{
		if !arr[i] {
			j++
			v = i << 1
			v++
			m, tn = divmod(cn , v)
			for tn == 0 {
				if m==1{
					return int64(cn)
				}
				cn = m
				m, tn = divmod(cn , v)
			}
		}
	}
	return int64(cn)
}

//FindBiggestFactorMiller A function using Miller Rabin 
//to find the biggest prime underneath
//an int64 that is a factor of the number n.
func FindBiggestFactorMiller(n int64) int64 {
	return fBFMR647(n)
}


//FindBiggestFactorErasthones : A function using Sieve of Erasthones 
//to find the biggest prime underneath
//an int64 that is a factor of the number n.
func FindBiggestFactorErasthones(n int) int {
	return int(fbEAFP(int64(n)))
}

func main(){
	n:= 600851475143
	if len(os.Args) > 1 {
		v, e := strconv.Atoi(os.Args[1])
		if e != nil {
			panic( e.Error() )
		}
		n = v
	}
	s:= math.Sqrt(float64(n))
	
	if s < float64(17179869184) {
		i := FindBiggestFactorMiller(int64(n))
		fmt.Println("The biggest factor of ", n , " is ", i, ".")
	} else {
		i := FindBiggestFactorErasthones(n)
		fmt.Println("The biggest factor of ", n , " is ", i, ".")
	
	}
}