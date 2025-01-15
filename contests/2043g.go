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

const BLOCK_SIZE = 300

type Gridg struct {
	N      int
	Prefix [][]int
	Data   [][]int
	RowAcc [][]int
	Dirty  []bool
}

func NewGridg(a []int) *Gridg {
	n := divceil(len(a), BLOCK_SIZE)
	g := &Gridg{
		N:      n,
		Prefix: make([][]int, n+1),
		Data:   make([][]int, n),
		RowAcc: make([][]int, n),
		Dirty:  make([]bool, n),
	}
	for i := range g.Prefix {
		g.Prefix[i] = make([]int, len(a)+1)
	}
	for i := range g.Data {
		g.Data[i] = make([]int, n)
		g.RowAcc[i] = make([]int, n+1)
	}
	for i, v := range a {
		g.Inc(i/BLOCK_SIZE, v)
	}
	return g
}

func (g *Gridg) Countr(i, j, v int) int {
	return g.Prefix[j+1][v] - g.Prefix[i][v]
}

func (g *Gridg) Inc(i, v int) {
	if g.Prefix[g.N][v] > 0 {
		for j := 0; j < g.N; j++ {
			g.Data[i][j] += g.Countr(j, j, v)
		}
		g.Dirty[i] = true
	}
	for j := i + 1; j <= g.N; j++ {
		g.Prefix[j][v] += 1
	}
}

func (g *Gridg) Dec(i, v int) {
	for j := i + 1; j <= g.N; j++ {
		g.Prefix[j][v] -= 1
	}
	if g.Prefix[g.N][v] > 0 {
		for j := 0; j < g.N; j++ {
			g.Data[i][j] -= g.Countr(j, j, v)
		}
		g.Dirty[i] = true
	}
}

func (g *Gridg) DoAcc(i int) {
	for j := 0; j < g.N; j++ {
		g.RowAcc[i][j+1] = g.RowAcc[i][j] + g.Data[i][j]
	}
}

func (g *Gridg) Query(i, j int) int64 {
	res := int64(0)
	for k := i; k <= j; k++ {
		if g.Dirty[k] {
			g.DoAcc(k)
			g.Dirty[k] = false
		}
		res += int64(g.RowAcc[k][j+1] - g.RowAcc[k][i])
	}
	return res
}

func run(arr []int, queries [][]int) {
	// fmt.Println(arr, queries)
	n := len(arr)
	last := 0
	mod := func(v int) int {
		x := v + last
		if x >= n {
			x -= n
		}
		return x + 1
	}
	res := make([]string, 0, len(queries))
	grid := NewGridg(arr)
	count := make([]int64, n+1)
	for _, q := range queries {
		if q[0] == 1 {
			p, x := mod(q[1])-1, mod(q[2])
			// fmt.Println("query1", arr, p, x)
			if arr[p] != x {
				grid.Dec(p/BLOCK_SIZE, arr[p])
				arr[p] = x
				grid.Inc(p/BLOCK_SIZE, arr[p])
			}
		} else {
			l, r := mod(q[1])-1, mod(q[2])-1
			if l > r {
				l, r = r, l
			}
			npair, size := int64(0), int64(r-l+1)

			// fmt.Println("query2", arr, l, r)
			bi, bj := (l/BLOCK_SIZE)+1, (r/BLOCK_SIZE)-1
			if bi > bj {
				for k := l; k <= r; k++ {
					npair += count[arr[k]]
					count[arr[k]] += 1
				}
				for k := l; k <= r; k++ {
					count[arr[k]] = 0
				}
			} else {
				npair += grid.Query(bi, bj)
				lo, hi := bi*BLOCK_SIZE, (bj+1)*BLOCK_SIZE
				for k := l; k < lo; k++ {
					npair += count[arr[k]] + int64(grid.Countr(bi, bj, arr[k]))
					count[arr[k]] += 1
				}
				for k := hi; k <= r; k++ {
					npair += count[arr[k]] + int64(grid.Countr(bi, bj, arr[k]))
					count[arr[k]] += 1
				}
				for k := l; k < lo; k++ {
					count[arr[k]] = 0
				}
				for k := hi; k <= r; k++ {
					count[arr[k]] = 0
				}

			}

			npair = size*(size-1)/2 - npair
			res = append(res, fmt.Sprint(npair))
			if npair < 0 {
				break
			}
			last = int(npair % int64(n))

		}
	}
	fmt.Println(strings.Join(res, " "))

}

func main() {

	// ntest := readInt()
	ntest := 1
	for nt := 0; nt < ntest; nt++ {
		readInt()
		a := readSliceInt()
		queries := make([][]int, readInt())
		for i := range queries {
			queries[i] = readSliceInt()
		}
		run(a, queries)

	}
}
