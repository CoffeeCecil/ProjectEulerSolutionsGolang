package primes

import (
	"fmt"
	"math/big"
	"testing"
)


func Test_Miller3(t *testing.T) {
	primeslist := [...]int64{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113, 127, 131, 137, 139, 149, 151, 157, 163, 167, 173, 179, 181, 191, 193, 197, 199}
	for i := int64(0); i < 200; i++ {
		if MillerRabin644(big.NewInt(i)) {
			for j, k := range primeslist {
				if k == i {
					primeslist[j] = 1
				}
			}
		}
	}
}

func Test_Miller2(t *testing.T) {
	primeslist := [...]int64{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113, 127, 131, 137, 139, 149, 151, 157, 163, 167, 173, 179, 181, 191, 193, 197, 199}
	for i := int64(0); i < 200; i++ {
		if MillerRabin643(big.NewInt(i)) {
			for j, k := range primeslist {
				if k == i {
					primeslist[j] = 1
				}
			}
		}
	}

	for _, k := range primeslist {
		if k != 1 {
			t.Log(k)
			panic("Didn't find all primes in primes list.")
		}
	}
}

func Test_Miller(t *testing.T) {
	primeslist := [...]int64{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113, 127, 131, 137, 139, 149, 151, 157, 163, 167, 173, 179, 181, 191, 193, 197, 199}
	for i := int64(0); i < 200; i++ {
		if Miller_Rabin_Custom(big.NewInt(i), toMiller64(), 0) {
			for j, k := range primeslist {
				if k == i {
					primeslist[j] = 1
				}
			}
		}
	}

	for _, k := range primeslist {
		if k != 1 {
			t.Log(k)
			panic("Didn't find all primes in primes list.")
		}
	}
}

func TestErasthones( t *testing.T){
	primeslist := [...]int64{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113, 127, 131, 137, 139, 149, 151, 157, 163, 167, 173, 179, 181, 191, 193, 197, 199}
	r := SieveEra(200)
	for _, v := range r {
		for j, w := range primeslist{
			if int64(v) == w {
				primeslist[j] = 1
			}
		}
	}

	for _, v := range primeslist{
		if v != 1{
			t.Log(r)
			t.Log(primeslist)
			panic("Primeslist incomplete!")
		}
	}
}

func TestNaiveErasthones( t *testing.T){
	primeslist := [...]int64{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113, 127, 131, 137, 139, 149, 151, 157, 163, 167, 173, 179, 181, 191, 193, 197, 199}
	r, _ := erasthonesNaiveArray(200)
	found:=true
	for i, v := range r {
		if !v  {
			found = false
			for j, w := range primeslist{
				if w == int64(i) {
					primeslist[j]=1
					found = true
					break
				}
			}
		}
		if !found{
			t.Log(i)
			panic("Couldn't find prime")
		}
	}

	for _, v := range primeslist {
		if v != 1 {
			t.Log(primeslist)
			t.Log("Primeslist incomplete!")
		}
	}
}

func Test_StressSizeErasthones(t *testing.T){
	v:= 1 << 32
	
	//Don't worry, you'll run out of memory.
	for {
		fmt.Println(v)
		_, vp := ErasthonesArray(v)
		fmt.Println(vp)
		v = v << 1
	}
}

func TestOddball( t *testing.T){
	primeslist := [...]int64{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113, 127, 131, 137, 139, 149, 151, 157, 163, 167, 173, 179, 181, 191, 193, 197, 199}
	r := OddballPrimesTest(0, big.NewInt(200) )
	temp:=int64(0)
	found := true
	for i,v := range r {
		if i > 0 && v {
			temp = int64(2*i + 1)
			found = false
		}
		for j, w := range primeslist {
			if w == temp {
				primeslist[j] = 1
				found = true
				break
			}
		}
		if found == false{
			t.Log(r)
			t.Log(i )
			t.Log( 2*i + 1)
			panic( "Didn't find prime!")
		}
	}
	primeslist[0] = 1
	for _, v := range primeslist {
		if v != 1 {
			t.Log(r)
			t.Log(primeslist)
			panic("Prime not found!")
		}
	}
}

func Benchmark_SieveEar( b *testing.B){
	v := int(^uint(0) >> 32) 

	for i:=0; i< b.N; i++{
		ErasthonesArray( v )
	}
}

func Benchmark_SieveEarNaive( b *testing.B){
	v := int(^uint(0) >> 32)
	for i:=0; i< b.N; i++{
		erasthonesNaiveArray(v)
	}
}

func Test_Fact(t *testing.T) {
	k := big.NewInt(5)
	Factorial(k, k)
	t.Log(k)
	if k.Cmp(big.NewInt(120)) != 0 {
		panic("Factorial failed to produce good value")
	}
}

func Benchmark_TLn_63_eqv(b *testing.B) {
	x := big.NewFloat(2)
	prec := big.NewFloat(18)
	for i := 0; i < b.N; i++ {
		TaylorLogEExtended(x, prec, 53)
	}
}

func Benchmark_TLn_101_digits(b *testing.B) {
	x := big.NewFloat(2)
	iter := big.NewFloat(6 * 18)
	prec := uint(7 * 53)
	for i := 0; i < b.N; i++ {
		TaylorLogEExtended(x, iter, prec)
	}
}

//So with calculating LN, this and the next test show where to use LN unbound.
//Specifically, notice how this test passes with the inputs given, whereas
//a naive ln function will take much, much, much more iterations?
//Now, this function completely fails to do that because when a function is divisible by 10
//you need 0 iterations, and the function will automatically compute ln0 and add m*ln10
func Test_Ln1234_Unbound25(t *testing.T) {
	//ln 1234 according to wolfram alpha
	stx := "7.118016204465333123414803800068367392789935050999118454826"[0:25]
	iter := big.NewFloat(50)
	prec := uint(100)
	l2 := LogEUnbound(big.NewFloat(1234), iter, prec)
	st2 := l2.Text('g', len(stx))
	for i, c := range st2 {
		if i >= len(stx) {
			break
		}
		if rune(stx[i]) != c {
			t.Log("Digits: ")
			t.Log(i, " of ", len(stx))
			if i < 20 {
				t.Log("Should be: ")
				t.Log(stx)
				t.Log("Got: ")
				t.Log(st2)
				panic("Doesn't Match!")
			}
			break
		}
	}
	if len(st2) < len(stx) {
		t.Log("Digits: ")
		t.Log(len(st2))
		t.Log(st2)
		t.Log(stx)
		panic("Too few digits.")
	}
}

//ln is more efficient when it is closer to 1. ln(2) requires less iterations
//than ln2 unbound.
//For example, this gets 4 correct digits with 1234 iterations.
func Test_Ln1234_to25(t *testing.T) {
	//ln1234
	stx := "7.118016204465333123414803800068367392789935050999118454826"[0:25]
	iter := big.NewFloat(1234)
	prec := uint(100)
	l2 := TaylorLogEExtended(big.NewFloat(1234), iter, prec)
	st2 := l2.Text('g', len(stx))
	for i, c := range st2 {
		if i >= len(stx) {
			break
		}
		if rune(stx[i]) != c {
			t.Log("Digits: ")
			t.Log(i, " of ", len(stx))
			if i < 20 {
				t.Log("Should be: ")
				t.Log(stx)
				t.Log("Got: ")
				t.Log(st2)
				panic("Doesn't Match!")
			}
			break
		}
	}
	if len(st2) < len(stx) {
		t.Log("Digits: ")
		t.Log(len(st2))
		t.Log(st2)
		t.Log(stx)
		panic("Too few digits.")
	}
}

func Test_Ln_Unbound25(t *testing.T) {
	stx := "0.69314718055994530941723"
	iter := big.NewFloat(59)
	prec := uint(100)
	l2 := LogEUnbound(big.NewFloat(2), iter, prec)
	st2 := l2.Text('g', len(stx))
	for i, c := range st2 {
		if i >= len(stx) {
			break
		}
		if rune(stx[i]) != c {
			t.Log("Digits: ")
			t.Log(i, " of ", len(stx))
			if i < 20 {
				t.Log("Should be: ")
				t.Log(stx)
				t.Log("Got: ")
				t.Log(st2)
				panic("Doesn't Match!")
			}
			break
		}
	}
	if len(st2) < len(stx) {
		t.Log("Digits: ")
		t.Log(len(st2))
		t.Log(st2)
		t.Log(stx)
		panic("Too few digits.")
	}
}

//ln is more efficient when it is closer to 1. ln(2) requires less iterations
//than ln2 unbound.
func Test_Ln_to25(t *testing.T) {
	stx := "0.69314718055994530941723"
	iter := big.NewFloat(59)
	prec := uint(100)
	l2 := TaylorLogEExtended(big.NewFloat(2), iter, prec)
	st2 := l2.Text('g', len(stx))
	for i, c := range st2 {
		if i >= len(stx) {
			break
		}
		if rune(stx[i]) != c {
			t.Log("Digits: ")
			t.Log(i, " of ", len(stx))
			if i < 20 {
				t.Log("Should be: ")
				t.Log(stx)
				t.Log("Got: ")
				t.Log(st2)
				panic("Doesn't Match!")
			}
			break
		}
	}
	if len(st2) < len(stx) {
		t.Log("Digits: ")
		t.Log(len(st2))
		t.Log(st2)
		t.Log(stx)
		panic("Too few digits.")
	}
}

//practical testing with ln shows that you can get 16 digits of accuracy
//with the equivalent of a float64. Love to test this against an actual float 64 for timing.
//with 18 iterations
func Test_Ln(t *testing.T) {
	//from http://oeis.org/A002162/constant for log_e(2)
	stx := "0.693147180559945309417232121458176568075500134360255254120680009493393621969694715605863326996418687"
	iter := big.NewFloat(6 * 18)
	prec := uint(7 * 53)

	l2 := TaylorLogEExtended(big.NewFloat(2), iter, prec)
	st2 := l2.Text('g', len(stx))
	for i, c := range st2 {
		if i >= len(stx) {
			break
		}
		if rune(stx[i]) != c {
			t.Log("Digits: ")
			t.Log(i, " of ", len(stx))
			if i < 20 {
				t.Log("Should be: ")
				t.Log(stx)
				t.Log("Got: ")
				t.Log(st2)
				panic("Doesn't Match!")
			}
			break
		}
	}
	if len(st2) < len(stx) {
		t.Log("Digits: ")
		t.Log(len(st2))
		t.Log(st2)
		t.Log(stx)
		panic("Too few digits.")
	}
}

func Test_E_Const(t *testing.T) {
	stx := "2.718281828459045235360287471352662497757247093699959574966967627724076630353547594571382178525166427427466391932003059921817413596629043572900334295260"
	//for the record, setting to 5000 and prec = 4066 gives 1234 digits
	arg := big.NewFloat(5000)
	prec := uint(8 * 4096)
	arg.SetPrec(prec)
	e := MakeE(arg, prec)
	st2 := e.Text('g', len(stx))
	_ = fmt.Println
	for i, c := range stx {
		if rune(st2[i]) != c {
			fmt.Println(i)
			if i < 20 {
				panic("Doesn't match!")
			}
			break
		}
	}
}

func Benchmark_PerfectPowers(b *testing.B){
	for i:=int64(0) ; i < int64(b.N); i++{
		perfectPowersTester(big.NewInt(i))
	}
}

func Benchmark_PerfectPowersAKS(b *testing.B){
	for i:=int64(0) ; i < int64(b.N); i++{
		PerfectPowerAKSCheck(big.NewInt(i))
	}
}

func TestPerfectPowers(t *testing.T){
	cnt := 0
	for i:=int64(0) ; i < 1766 ; i++ {
		if(perfectPowersTester(big.NewInt(i))){
			t.Log("Is perfect power: ")
			t.Log(i)
			cnt++
		}
	}
	t.Log(cnt)
}

func TestPerfectAKSPowers(t *testing.T){
	cnt := 0
	for i:=int64(0) ; i < 1766 ; i++ {
		if(PerfectPowerAKSCheck(big.NewInt(i))){
			t.Log("Is perfect power: ")
			t.Log(i)
			cnt++
		}
	}
	t.Log("cnt is: ", cnt)
}