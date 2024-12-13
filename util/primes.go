package util

import "slices"

var notprimes = make([]bool, 300001)
var primes = []int{}

func init() {
	for i := 2; i < len(notprimes); i++ {
		if notprimes[i] {
			continue
		}
		primes = append(primes, i)
		for j := i * i; j < len(notprimes); j += i {
			notprimes[j] = true
		}
	}
}

var factors = make([][]int, 1000001)

func init() {
	for i := 2; i < len(factors); i++ {
		if len(factors[i]) > 0 {
			continue
		}
		for j := i; j < len(factors); j += i {
			factors[j] = append(factors[j], i)
		}
	}
}

// func factorize(v int) []int {
// 	lo, hi := []int{}, []int{}
// 	for i := 1; i*i <= v; i++ {
// 		if i*i == v {
// 			lo = append(lo, i)
// 			break
// 		} else if v%i == 0 {
// 			lo = append(lo, i)
// 			hi = append(hi, v/i)
// 		}
// 	}
// 	for j := len(hi) - 1; j >= 0; j-- {
// 		lo = append(lo, hi[j])
// 	}
// 	return lo
// }

func gcd[T int | int64](a, b T) T {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

func lcm[T int | int64](a, b T) T {
	return a * (b / gcd(a, b))
}

var mfactors = make([]int, 3000001)

func init() {
	for i := 2; i < len(mfactors); i++ {
		if mfactors[i] > 0 {
			continue
		}
		mfactors[i] = i

		for j := i * i; j < len(mfactors); j += i {
			if mfactors[j] == 0 {
				mfactors[j] = i
			}
		}
	}
}

var cachedFactorize = map[int][]int{}

func factorize(v int) []int {
	if v == 1 {
		return []int{1}
	}
	if x, ok := cachedFactorize[v]; ok {
		return x
	}
	key := v

	minf, c := mfactors[v], 0
	for mfactors[v] == minf {
		c += 1
		v = v / minf
	}

	prev := factorize(v)
	res := make([]int, 0, len(prev)*(c+1))
	for _, v := range prev {
		res = append(res, v)
		for i := 0; i < c; i++ {
			v *= minf
			res = append(res, v)
		}
	}
	slices.Sort(res)
	cachedFactorize[key] = res
	return res
}

func countGCD(n int) [][2]int {
	factors := factorize(n)
	m := map[int]int{}
	for _, f := range factors {
		m[f] = n / f
	}
	for i := len(factors) - 1; i >= 0; i-- {
		v := factors[i]
		for _, mul := range factorize(n / v) {
			if mul > 1 {
				m[v] -= m[v*mul]
			}
		}
	}
	res := make([][2]int, len(factors))
	for i, f := range factors {
		res[i] = [2]int{f, m[f]}
	}
	return res
}
