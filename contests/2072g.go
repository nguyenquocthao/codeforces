package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func readInt() int {
	n := 0
	fmt.Fscanf(reader, "%d\n", &n)
	return n
}

func readInt64() int64 {
	n := int64(0)
	fmt.Fscanf(reader, "%d\n", &n)
	return n
}

func readString() string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func readSliceInt() []int {
	s := strings.TrimSpace(readString())
	if s == "" {
		return []int{}
	}
	data := strings.Split(s, " ")
	res := make([]int, len(data))
	for i, v := range data {
		res[i], _ = strconv.Atoi(v)
	}
	return res
}

func readSliceInt64() []int64 {
	s := strings.TrimSpace(readString())
	if s == "" {
		return []int64{}
	}
	data := strings.Split(s, " ")
	res := make([]int64, len(data))
	for i, v := range data {
		res[i], _ = strconv.ParseInt(v, 10, 64)
	}
	return res
}

func readSliceString() []string {
	s := strings.TrimSpace(readString())
	if s == "" {
		return []string{}
	}
	return strings.Split(s, " ")
}

func printSlice[T any](l []T) {
	output := make([]string, len(l))
	for i, v := range l {
		output[i] = fmt.Sprint(v)
	}
	fmt.Println(strings.Join(output, " "))
}

func Max[T int | float32 | string | int64 | byte](args ...T) T {
	res := args[0]
	for i := 1; i < len(args); i++ {
		if args[i] > res {
			res = args[i]
		}
	}
	return res
}

func Min[T int | float32 | string | int64 | byte](args ...T) T {
	res := args[0]
	for i := 1; i < len(args); i++ {
		if args[i] < res {
			res = args[i]
		}
	}
	return res
}

func Sum[T int | float32 | int64](args ...T) T {
	var res T
	for _, v := range args {
		res += v
	}
	return res
}

func abs[T int | int64](v T) T {
	if v < 0 {
		return -v
	}
	return v
}

func reverseStr(s string) string {
	rns := []rune(s) // convert to rune
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		// swap the letters of the string,
		// like first with last and so on.
		rns[i], rns[j] = rns[j], rns[i]
	}

	// return the reversed string.
	return string(rns)
}

func Unique[T comparable](l []T) []T {
	m := map[T]bool{}
	for _, v := range l {
		m[v] = true
	}
	res := make([]T, 0, len(l))
	for v := range m {
		res = append(res, v)
	}
	return res
}

func divceil[T int | int64](a, b T) T {
	res := a / b
	if a%b > 0 {
		res += 1
	}
	return res
}

func makeRange(i, j int) []int {
	res := make([]int, j-i)
	for k := 0; k < len(res); k++ {
		res[k] = i + k
	}
	return res
}

func Repeat[T any](v T, n int) []T {
	res := make([]T, n)
	for i := 0; i < n; i++ {
		res[i] = v
	}
	return res
}

func divneg[T int | int64](a, b T) T {
	res := a / b
	m := a % b
	if m < 0 {
		res -= 1
	}
	return res
}

func accumulate(a []int64) []int64 {
	res := make([]int64, len(a)+1)
	for i, v := range a {
		res[i+1] = res[i] + v
	}
	return res
}

func Count[T comparable](l []T) map[T]int {
	m := map[T]int{}
	for _, v := range l {
		m[v] += 1
	}
	return m
}

func Keys[K comparable, V any](m map[K]V) []K {
	res := []K{}
	for k := range m {
		res = append(res, k)
	}
	return res
}

func IfElse[T any](condition bool, a, b T) T {
	if condition {
		return a
	} else {
		return b
	}
}

func gcd[T int | int64](a, b T) T {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

func lcm[T int | int64](a, b T) T {
	return a * (b / gcd(a, b))
}

func tol(s string) []int {
	res := make([]int, len(s))
	for i, ch := range s {
		if ch == '1' {
			res[i] = 1
		} else {
			res[i] = 0
		}
	}
	return res
}

func ToSet[T comparable](l []T) map[T]bool {
	m := map[T]bool{}
	for _, v := range l {
		m[v] = true
	}
	return m
}
func Map[T, T1 any](l []T, f func(v T) T1) []T1 {
	res := make([]T1, len(l))
	for i, v := range l {
		res[i] = f(v)
	}
	return res
}

func Contains[T comparable](l []T, x T) bool {
	for _, v := range l {
		if v == x {
			return true
		}
	}
	return false
}

// const MOD = 998244353
const MOD = 1_000_000_007

func pow(x, n int64) int64 {
	x = x % MOD
	res := int64(1)
	for n > 0 {
		if n%2 == 1 {
			res = (res * x) % MOD
		}
		x = (x * x) % MOD
		n = n / 2
	}
	return res
}

func mod_inverse[T int | int64](x T) T {
	res := pow(int64(x), MOD-2)
	return T(res)
}

func mod[T int | int64](v T) T {
	res := v % MOD
	if res < 0 {
		res += MOD
	}
	return res
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

func calx(a, b, stepb int64) int64 {
	// a*b + (a+1)*(b-stepb) + (a+2)*(b-2*stepb)...
	n := (b / stepb)
	res := mod((n + 1) * mod(a*b))
	res = mod(res + mod(Sump2(n)*mod(b-a*stepb)))
	res = mod(res - mod(stepb*Sump3(n)))
	return res
}

// var cacheinv = map[[2]int][2]int{}

func calinv(n, p int64) (int64, int64) {
	if n < p {
		return n, 1
	}
	// key := [2]int{int(n), int(p)}
	// if v, ok := cacheinv[key]; ok {
	// 	return int64(v[0]), int64(v[1])
	// }
	first := n % p
	a, b := calinv(n/p, p)
	a = mod(a + mod(first*pow(p, b)))
	b += 1
	// cacheinv[key] = [2]int{int(a), int(b)}
	return a, b
}

const THRES = 600

func run(n, k int64) int64 {
	res := int64(0)
	hi := min(THRES, k)
	for p := int64(2); p <= hi; p++ {
		added, _ := calinv(n, p)
		res = mod(res + added)
		// fmt.Println(p, res)
	}
	if k <= THRES {
		return res
	}
	if n <= THRES {
		return mod(res + mod(mod(n)*mod(k-THRES)))
	}

	p := int64(THRES) + 1
	for p <= n {
		// fmt.Println(p-1, res)
		step, b := n/p, n%p
		ns := (b / step) + 1
		added := mod(calx(p, b, step) + mod(ns*step))
		res = mod(res + added)
		// fmt.Println(p+ns-1, res)
		if p+ns > k {
			k += 1
			if n/k == step {
				b = n % k
				ns = (b / step) + 1
				added := mod(calx(k, b, step) + mod(ns*step))
				res = mod(res - added)
			}
			return res
		}
		p += ns
	}
	return mod(res + mod(mod(n)*mod(k-n)))

}

var debug = false

func main() {
	// fmt.Println(calx(14, 12, 2))
	ntest := readInt()
	// ntest := 1
	// debug = ntest == 7752
	for nt := 0; nt < ntest; nt++ {
		l := readSliceInt64()
		if nt == 0 && l[0] == 69995 && l[1] == 693110035550211703 {
			debug = false
		}
		if !debug {
			fmt.Println(run(l[0], l[1]))
		} else if nt == 1102 {
			fmt.Println(l)
		}

	}
}
