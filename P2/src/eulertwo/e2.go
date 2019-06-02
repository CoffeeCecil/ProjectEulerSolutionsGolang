package main

import (
	"math/big"
	"fmt"
	"os"
)

//Fib_even_faster (n *big.Int) *big.Int : Function for generating the sum of even Fib numbers.
//Uses the recurrence Fib_even = 4*F_(n-3) + F_(n-6)
//Derived from: https://www.mathblog.dk/project-euler-problem-2/
func Fib_even_faster(n *big.Int) *big.Int{
	a:= big.NewInt(2)
	if n.Cmp(a) < 0{
		a.SetInt64(0)
		return a
	}
	a.SetInt64(8)
	if n.Cmp(a) < 0{
		a.SetInt64(2)
		return a
	}
	
	a.SetInt64(2)
	var four big.Int
	four.SetInt64(4)
	b := big.NewInt(0)
	c := big.NewInt(0)
	r := big.NewInt(0)
	for a.Cmp(n) < 0{
		r.Add(r, a)
		c.Set(b)
		b.Set(a)
		a.Mul(a,&four)
		a.Add(a,c)
	}
	return r
}

//Fib_even (n *big.Int) *big.Int Function for generating even fib numbers.
//Relies on the fact that the 3rd even fib number will always be even.
func Fib_even(n *big.Int) *big.Int{
	a:= big.NewInt(2)
	if n.Cmp(a) < 0{
		a.SetInt64(0)
		return a
	}
	a.SetInt64(8)
	if n.Cmp(a) < 0{
		a.SetInt64(2)
		return a
	}
	a.SetInt64(1)
	b := big.NewInt(2)
	r := big.NewInt(2)
	for true{
		a.Add(a, b)//fn-1 = 3
		b,a = a, b
		a.Add(a, b)//fn-1 = 5
		b,a = a, b//move
		a.Add(a,b)//fn = 8
		b,a = a,b
		//
		//if b = fn_even is over n, break
		if (b.Cmp(n) >= 0){
			break
		}
		//otherwise accumulate
		r.Add(r,b)
	}
	return r
}
//Fib_naive (n *big.Int) *big.Int Function for verification purposes.
func Fib_naive(n *big.Int) *big.Int {
	//a = f_(n-2)
	a:= big.NewInt(0)
	if n.Cmp(a) <= 0{
		return a
	}
	//b = f_n(-1)
	b := big.NewInt(1)
	if n.Cmp(b) == 0{
		return a
	}
	operator := big.NewInt(2)
	var zero big.Int;
	var two big.Int; 
	zero.SetInt64(0)
	two.SetInt64(2)
	a.SetInt64(1)
	fn := big.NewInt(0)
	//must if it isn't 1, and then it must be fib4 = 2
	r:= big.NewInt(0)
	for (fn.Cmp(n))<0 {
		fn.Add(a, b)
		a.Set(b)
		b.Set(fn)

		if operator.Mod(fn, &two).Cmp(&zero) == 0 {
			r.Add(r, fn)
		}
		//fmt.Println("Fib is ", fn)
	}
	return r
}

func main(){
	n, two := big.NewInt(4000000), big.NewInt(2)
	success := false
	if len(os.Args) > 1 {
		n,success = n.SetString(os.Args[1], 10)
		if !success{
			fmt.Println("Could not parse argument. Program takes a number greater than two or no input.")
		}
		if n.Cmp(two) < 0 {
			fmt.Println("Program takes a number greater than two or no input.")
		}
	}
	m := Fib_even_faster(n)
	fmt.Println("The sum of the even Fib terms under ", n , " is ", m, ".")
}