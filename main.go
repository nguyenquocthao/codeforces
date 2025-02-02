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

type Stack[T any] struct {
	data []T
	i    int
}

func (s *Stack[T]) Top() T {
	return s.data[s.i-1]
}

func (s *Stack[T]) Len() int {
	return s.i
}

func (s *Stack[T]) Push(v T) {
	if s.i == len(s.data) {
		s.data = append(s.data, v)
	} else {
		s.data[s.i] = v
	}
	s.i += 1
}

func (s *Stack[T]) Pop() T {
	s.i -= 1
	if s.i < 0 {
		panic("Invalid stack pop: stack is empty")
	}
	return s.data[s.i]
}

func (s *Stack[T]) ToList() []T {
	return s.data[:s.i]
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{data: []T{}, i: 0}
}

const MOD = 998244353

// const MOD = 1000000007
const maxn = 1000000

var FAC = make([]int64, maxn+1)
var IFAC = make([]int64, maxn+1)

func init() {
	FAC[0], FAC[1] = 1, 1
	IFAC[0], IFAC[1] = 1, 1
	for i := int64(2); i < maxn+1; i++ {
		FAC[i] = (i * FAC[i-1]) % MOD
		IFAC[i] = mod_inverse(FAC[i])
	}
}

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

func mod_inverse(x int64) int64 {
	return pow(x, MOD-2)
}

func comb(n, k int64) int64 {
	if n < 0 || k > n {
		return 0
	}
	inv := (IFAC[k] * IFAC[n-k]) % MOD
	return (FAC[n] * inv) % MOD
}

func mod[T int | int64](v T) T {
	res := v % MOD
	if res < 0 {
		res += MOD
	}
	return res
}

func Catalan(n int) int64 {
	// fmt.Println("Catalan", n)
	return mod(FAC[2*n] * mod(IFAC[n+1]*IFAC[n]))
}

func iCatalan(n int) int64 {
	return mod(IFAC[2*n] * mod(FAC[n+1]*FAC[n]))
}

func run(a []int, k int) int {
	// stay the same
	// swap with first
	n := len(a)
	k -= 1
	badi := -1
	for i, v := range a {
		if v > a[k] {
			badi = i
			break
		}
	}
	if badi < 0 {
		return n - 1
	}
	if badi > k {
		return badi - 1
	}
	// swap with first bad
	res := 0
	if badi > 0 {
		res += 1
	}
	val := a[k]
	a[k], a[badi] = a[badi], a[k]
	for j := badi + 1; j < n; j++ {
		if val > a[j] {
			res += 1
		} else {
			break
		}
	}
	return Max(badi-1, res)

}

func main() {
	ntest := readInt()
	// ntest := 1
	for nt := 0; nt < ntest; nt++ {
		l := readSliceInt()
		fmt.Println(run(readSliceInt(), l[1]))

	}
}
