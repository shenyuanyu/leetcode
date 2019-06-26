package task50

func myPow(x float64, n int) float64 {
	if x == 0 {
		return 0
	}
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	if n < 0 {
		n = -n
		x = 1 / x
	}

	t := myPow(x, n/2)
	t *= t
	if n%2 != 0 {
		t *= x
	}

	return t
}
