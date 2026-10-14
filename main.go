package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime/debug"
	"slices"
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

// func Unique[T comparable](l []T) []T {
// 	m := map[T]bool{}
// 	for _, v := range l {
// 		m[v] = true
// 	}
// 	res := make([]T, 0, len(l))
// 	for v := range m {
// 		res = append(res, v)
// 	}
// 	return res
// }

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

func Accumulate[T int | int64](a []T) []T {
	res := make([]T, len(a)+1)
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

// const MOD = 1_000_000_007

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

func mod[T int | int64](v T) T {
	res := v % MOD
	if res < 0 {
		res += MOD
	}
	return res
}

type Deque[T any] struct {
	Data       []T
	Start, End int
}

func (dq *Deque[T]) Len() int {
	return dq.End - dq.Start
}

func NewDeque[T any](l []T, capacity int) *Deque[T] {
	res := &Deque[T]{}
	if capacity < 2 {
		capacity = 2
	}
	res.Start = 0
	res.End = len(l)
	if len(l) >= capacity {
		res.Data = l
	} else {
		res.Data = make([]T, capacity)
		copy(res.Data, l)
	}
	// fmt.Println(res)
	return res
}

func (dq *Deque[T]) upgradeFull() {
	if dq.Len() < len(dq.Data) {
		return
	}
	// fmt.Println("upgrade", dq)
	data := append(dq.Data[dq.Start:], dq.Data[:dq.Start]...)
	dq.Start = 0
	dq.End = len(data)
	dq.Data = make([]T, 2*len(data)+1)
	copy(dq.Data, data)
}

func (dq *Deque[T]) First() T {
	// fmt.Println(dq)
	return dq.Data[dq.Start]
}
func (dq *Deque[T]) Last() T {
	// fmt.Println(dq)
	j := dq.End - 1
	if j >= len(dq.Data) {
		j -= len(dq.Data)
	}
	return dq.Data[j]
}

func (dq *Deque[T]) At(i int) T {
	i += dq.Start
	if i >= len(dq.Data) {
		i -= len(dq.Data)
	}
	return dq.Data[i]
}

func (dq *Deque[T]) AppendLeft(v T) {
	dq.upgradeFull()
	if dq.Start == 0 {
		dq.Start, dq.End = dq.Start+len(dq.Data), dq.End+len(dq.Data)
	}
	dq.Start -= 1
	dq.Data[dq.Start] = v
}

func (dq *Deque[T]) Append(v T) {
	dq.upgradeFull()
	j := dq.End
	if j >= len(dq.Data) {
		j -= len(dq.Data)
	}
	dq.Data[j] = v
	dq.End += 1
}

func (dq *Deque[T]) PopLeft() T {
	res := dq.First()
	dq.Start += 1
	if dq.Start == len(dq.Data) {
		dq.Start, dq.End = dq.Start-len(dq.Data), dq.End-len(dq.Data)
	}
	return res
}

func (dq *Deque[T]) Pop() T {
	res := dq.Last()
	dq.End -= 1
	return res
}

func (dq Deque[T]) ToList() []T {
	res := make([]T, dq.Len())
	for i := range res {
		res[i] = dq.At(i)
	}
	return res
}

func (dq Deque[T]) String() string {
	return fmt.Sprint(dq.ToList())
}

type UnionFind []int

func (u UnionFind) Find(x int) int {
	if u[x] != x {
		u[x] = u.Find(u[x])
	}
	return u[x]
}

func (u UnionFind) Union(x, y int) {
	u[u.Find(y)] = u.Find(x)
}

func (u UnionFind) NGroup() int {
	res := 0
	for i, v := range u {
		if i == v {
			res += 1
		}
	}
	return res
}

func (u UnionFind) GetRoots() []int {
	res := make([]int, 0, len(u))
	for i, v := range u {
		if i == v {
			res = append(res, i)
		}
	}
	return res
}

func Filter[T any](l []T, f func(v T) bool) []T {
	res := []T{}
	for _, v := range l {
		if f(v) {
			res = append(res, v)
		}
	}
	return res
}

func cal(a []int) (int, int) {
	maxv, minv, cur := 0, 0, 0
	for _, v := range a {
		cur += v
		maxv = max(maxv, cur)
		minv = min(minv, cur)
	}
	return maxv, minv
}

func calsub(a []int) (int, int) {
	maxv, minv, cur := 0, 0, 0
	submax, submin := 0, 0
	for _, v := range a {
		cur += v
		maxv = max(maxv, cur)
		minv = min(minv, cur)
		submax = max(submax, cur-minv)
		submin = min(submin, cur-maxv)
	}
	return submax, submin
}

func run(a []int) []int {
	ii := -1
	for i, v := range a {
		if v < -1 || v > 1 {
			ii = i
			break
		}
	}
	if ii < 0 {
		maxv, minv := calsub(a)
		return makeRange(minv, maxv+1)
	}
	aa := a[:ii]
	slices.Reverse(aa)
	max0, min0 := cal(aa)
	val := a[ii]
	max1, min1 := cal(a[ii+1:])
	ra := makeRange(val+min0+min1, val+max0+max1+1)
	max0, min0 = calsub(aa)
	// fmt.Println("calsub", aa, max0, min0)
	max1, min1 = calsub(a[ii+1:])
	rb := makeRange(min(min0, min1), max(max0, max1)+1)
	if ra[0] > rb[0] {
		ra, rb = rb, ra
	}
	for len(ra) > 0 && ra[len(ra)-1] >= rb[0] {
		ra = ra[:len(ra)-1]
	}
	ra = append(ra, rb...)
	return ra
}

func main() {
	// x := makeRange(0, 21)
	// for i := 1; i <= 10; i++ {
	// 	y := []int{}
	// 	for _, v := range x {
	// 		y = append(y, v%i)
	// 	}
	// 	fmt.Println(i, y)
	// }

	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
			debug.PrintStack()
			// buf := make([]byte, 1024)
			// n := runtime.Stack(buf, false)
			// fmt.Println("Panic recovered:", r)
			// runtime.Stack()
			// fmt.Println(strings.ReplaceAll(string(buf[:n]), "\n", " "))
		}
	}()
	// for n := 1000; n < 1001; n++ {
	// 	fmt.Println(n, generate(n), len(generate(n)))
	// }
	// fmt.Println(permutation([]int{0, 1, 2}))
	// ntest := 1
	// run(1073741824)
	// for i := 0; i < 100; i++ {
	// 	var x int = 1 << i
	// 	fmt.Println(i, x)
	// }
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i, Catalan(i))
	// }
	// info := NewInfo()
	// for i := 0; i < 3; i++ {
	// 	v := info.Add(2)
	// 	// fmt.Println(info.st.ToList(), info.n, v)
	// }
	// fmt.Println(factors[:20])

	ntest := readInt()
	// ntest := 1
	// debug = ntest == 10000
	for nt := 0; nt < ntest; nt++ {
		readInt()
		res := run(readSliceInt())
		fmt.Println(len(res))
		printSlice(res)
	}
}
