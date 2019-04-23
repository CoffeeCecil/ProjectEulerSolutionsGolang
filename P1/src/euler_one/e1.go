package euler_one

func euler_one_better(n int) (int, int) {
	m, i, k := 0, 0, 0
	s := [7]int{3, 2, 1, 3, 1, 2, 3}
	for k < n {
		m += k
		k += s[i]
		i++
		if i > 6 {
			i = 0
		}
	}
	k -= s[i-1]
	return m, k
}

func euler_one_fastest(n int) (int, int) {
	e := n
	for e%3 != 0 {
		e -= 1
	}
	e3 := e / 3
	e = n - 1
	for e%5 != 0 {
		e -= 1
	}
	e5 := e / 5
	for e%15 != 0 {
		e -= 1
	}
	e15 := e / 15
	sumn := func(n int) int {
		return (n*n + n) / 2
	}
	return 3*sumn(e3) + 5*sumn(e5) - 15*sumn(e15), n - 1
}

func euler_one_naive(n int) (int, int) {
	m, i := 0, 0
	for i < n {

		if i%3 == 0 {
			m += i
		} else if i%5 == 0 {
			m += i
		}
		i++
	}
	return m, i
}
