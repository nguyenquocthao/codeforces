package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
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

func query(l, r int) int {
	fmt.Println("?", l, r)
	return readInt()
}

type Counter struct {
	M     map[int]int
	Nitem int
}

func NewCounter() *Counter {
	return &Counter{map[int]int{}, 0}
}

func (c *Counter) Inc(v, added int) {
	c.Nitem += added
	c.M[v] += added
	if c.M[v] == 0 {
		delete(c.M, v)
	}
}

func (c *Counter) Nunique() int {
	return len(c.M)
}

func (c *Counter) Merge(c2 *Counter) {
	c.Nitem += c2.Nitem
	for k, v := range c2.M {
		c.M[k] += v
	}
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
			res++
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

type Heap []int

func (h Heap) Len() int           { return len(h) }
func (h Heap) Less(i, j int) bool { return h[i] > h[j] }
func (h Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *Heap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *Heap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func ToInt64(v interface{}) int64 {
	switch val := v.(type) {
	case int:
		return int64(val)
	case int64:
		return val
	default:
		// return 0, fmt.Errorf("cannot convert type %s to int64", reflect.TypeOf(v))
		panic(fmt.Errorf("cannot convert type %s to int64", reflect.TypeOf(v)))
	}
}

func run(a, b [][]int) bool {
	nrow, ncol := len(a), len(a[0])
	domask := func(a [][]int, mask int) [][]bool {
		res := make([][]bool, nrow)
		for i, r := range a {
			res[i] = make([]bool, ncol)
			for j, v := range r {
				if v&mask > 0 {
					res[i][j] = true
				}
			}
		}
		return res
	}
	keyf := func(isrow bool, i int) int {
		if isrow {
			return ncol + i
		} else {
			return i
		}
	}
	check := func(ba, bb [][]bool) bool {
		graph := make([][]int, nrow+ncol)
		for i, r := range bb {
			for j, v := range r {
				if v {
					// row operation => col operation
					graph[keyf(true, i)] = append(graph[keyf(true, i)], keyf(false, j))
				} else {
					// col operation => row operation
					graph[keyf(false, j)] = append(graph[keyf(false, j)], keyf(true, i))
				}
			}
		}
		marked := make([]int, len(graph))
		var hascycle func(i int) bool
		hascycle = func(i int) bool {
			if marked[i] > 0 {
				return marked[i] == 2
			}
			marked[i] = 2
			for _, j := range graph[i] {
				if hascycle(j) {
					return true
				}
			}
			marked[i] = 1
			return false
		}
		for i, r := range ba {
			for j := range r {
				if ba[i][j] && !bb[i][j] {
					if hascycle(keyf(true, i)) {
						return false
					}
				} else if !ba[i][j] && bb[i][j] {
					if hascycle(keyf(false, j)) {
						return false
					}
				}
			}
		}
		return true
	}
	for nb := 0; nb <= 30; nb++ {
		ba, bb := domask(a, 1<<nb), domask(b, 1<<nb)
		if !check(ba, bb) {
			return false
		}
	}
	return true
}

func main() {

	ntest := readInt()
	// ntest := 1
	for nt := 0; nt < ntest; nt++ {
		l := readSliceInt()
		a, b := make([][]int, l[0]), make([][]int, l[0])
		for i := range a {
			a[i] = readSliceInt()
		}
		for i := range b {
			b[i] = readSliceInt()
		}
		if run(a, b) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
