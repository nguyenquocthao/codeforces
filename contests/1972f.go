package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime/debug"
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

const MAXV = int64(1) << 60

func processFunctionalGraph(a []int) func(int, int) int {
	n := len(a)
	root := Repeat(-1, n)
	var getroot func(i int) int
	getroot = func(i int) int {
		if root[i] < 0 {
			root[i] = i
			root[i] = getroot(a[i])
		}
		return root[i]
	}
	mchild := make([][]int, n)
	for i, j := range a {
		if i == getroot(i) {
			continue
		}
		mchild[j] = append(mchild[j], i)
	}
	depth, size, order := make([]int, n), make([]int, n), make([]int, n)
	curorder := 0
	var dp func(i int)
	dp = func(i int) {
		size[i] = 1
		order[i] = curorder
		curorder += 1
		for _, j := range mchild[i] {
			depth[j] = depth[i] + 1
			dp(j)
			size[i] += size[j]
		}
	}
	for i := 0; i < n; i++ {
		if i == getroot(i) {
			dp(i)
		}
	}
	isChildOf := func(x, y int) bool {
		return order[x] >= order[y] && order[x] < order[y]+size[y]
	}
	dist := func(x, y int) int {
		if getroot(x) != getroot(y) {
			return n
		}
		r := getroot(x)
		if isChildOf(x, y) {
			return depth[x] - depth[y]
		} else if isChildOf(a[r], y) {
			// x->r->a[r]->y
			return depth[x] + 1 + depth[a[r]] - depth[y]
		} else {
			return n
		}
	}
	return dist
}

func run(a, b []int) int {
	n := len(b)
	b = append([]int{0}, b...)
	dist := processFunctionalGraph(b)
	// for i := 1; i <= n; i++ {
	// 	for j := 1; j <= n; j++ {
	// 		fmt.Println(i, j, dist(i, j))
	// 	}
	// }
	feasible := func(ntrick int) bool {
		cur := 1
		for _, v := range a {
			for cur <= n && dist(v, cur) > ntrick {
				cur += 1
			}
			if cur > n {
				return false
			}
		}
		return true
	}
	res, l, r := -1, 0, n-1
	for l <= r {
		mid := (l + r) / 2
		if feasible(mid) {
			res, r = mid, mid-1
		} else {
			l = mid + 1
		}
	}
	return res
}

func main() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
			debug.PrintStack()
			// fmt.Println(data)

			// buf := make([]byte, 1024)
			// n := runtime.Stack(buf, false)
			// fmt.Println("Panic recovered:", r)
			// runtime.Stack()
			// fmt.Println(strings.ReplaceAll(string(buf[:n]), "\n", " "))
		}
	}()
	// fmt.Println(permutation([]int{0, 1, 2}))
	// ntest := 1
	// run(1073741824)
	// for i := 0; i < 100; i++ {
	// 	var x int = 1 << i
	// 	fmt.Println(i, x)
	// }
	ntest := readInt()
	// ntest := 1

	// debug = ntest == 10000
	for nt := 0; nt < ntest; nt++ {
		readSliceInt()
		fmt.Println(run(readSliceInt(), readSliceInt()))
	}
}
