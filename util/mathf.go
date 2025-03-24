package util

// (x+1)**n == sum(comb(n, i) * x**i)
// => sum(comb(n, i)) = 2**n
// n(x+1)**(n-1) = sum(comb(n,i) * i * x**(i-1))
// sum(comb(n,i) * i) == n*2**(n-1)
// n*(n-1)*(x+1)**(n-2) = sum(comb(n,i) * i * (i-1) * x**(i-2))
// sum(comb(n,i) * i * (i-1)) == n*(n-1) * 2**(n-2)
// sum(comb(n,i) * i*i) = n*(n-1)* 2**(n-2) + n*2**(n-1) = n*(n+1) * 2**(n-2)

func ToPolyMod(f func(int64) int64, p int) func(int64) int64 {
	cal := make([]int64, p+1)
	for i := 0; i <= p; i++ {
		cal[i] = f(int64(i))
	}
	coeff := make([]int64, p+1)
	if p == 0 {
		coeff[0] = cal[0]
	} else if p == 1 {
		// co0 == cal0
		// co0 + c01 == cal1
		coeff[0] = cal[0]
		coeff[1] = mod(cal[1] - cal[0])
	} else if p == 2 {
		// c0 + 0*c1 + 0*c2 = cal0
		// c0 + c1 + c2 = cal1
		// c0 + 2*c1 + 4*c2 = cal2
		coeff[0] = cal[0]
		// c1 + c2 == cal1 - c0
		// 2*c1 + 4*c2 = cal2 - c0
		coeff[2] = mod(mod_inverse(int64(2)) * mod(cal[2]-2*cal[1]+cal[0]))
		coeff[1] = mod(cal[1] - coeff[0] - coeff[2])
	}
	return func(x int64) int64 {
		res := int64(0)
		mul := int64(1)
		for _, v := range coeff {
			res = mod(res + mod(mul*v))
			mul = mod(mul * x)
		}
		return res
	}
}

func Sump2(n int64) int64 {
	// n*(n+1)/2
	n = mod(n)
	return mod(mod_inverse(int64(2)) * mod(n*(n+1)))
}

func Sump3(n int64) int64 {
	// n*(n+1)*(2*n+1)/6
	n = mod(n)
	temp := mod(n * (n + 1))
	return mod(mod_inverse(int64(6)) * mod(temp*mod(2*n+1)))
}
