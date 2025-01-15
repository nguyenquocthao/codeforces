package main

import (
	"bufio"
	"container/heap"
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

type HItem struct {
	Size, I int
}

type Heap []HItem

func (h Heap) Len() int           { return len(h) }
func (h Heap) Less(i, j int) bool { return h[i].Size > h[j].Size }
func (h Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *Heap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(HItem))
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

type Info struct {
	K  int
	A  []int64
	Uf UnionFind
	M  []*Counter
	H  *Heap
}

func NewInfo(a []int64, k int) *Info {
	n := len(a)
	info := &Info{k, a, UnionFind(makeRange(0, n)), make([]*Counter, n), &Heap{}}
	for i, v := range a {
		if info.Validv(v) {
			info.M[i] = NewCounter()
			info.M[i].Inc(int(v), 1)
			if i > 0 && info.Validv(a[i-1]) {
				info.Merge(i-1, i)
			}
		}
	}
	info.InitH()

	return info
}

func (info *Info) InitH() {
	info.H = &Heap{}
	for _, i := range info.Uf.GetRoots() {
		info.CheckUpdateH(i)
	}
}

func (info *Info) CheckUpdateH(i int) {
	m := info.GetCounter(i)
	if m != nil && m.Nunique() == info.K {
		heap.Push(info.H, HItem{m.Nitem, i})
	}
}

func (info *Info) GetLongest() int {
	if info.H.Len() > len(info.A) {
		info.InitH()
	}
	for info.H.Len() > 0 {
		it := (*info.H)[0]
		m := info.GetCounter(it.I)
		if it.Size != m.Nitem || m.Nunique() != info.K {
			// continue
			heap.Pop(info.H)
		} else {
			return it.Size
		}
	}
	return 0
}

func (info *Info) Validv(v interface{}) bool {
	return ToInt64(v) < int64(info.K)
}

func (info *Info) Merge(i, j int) {
	i, j = info.Uf.Find(i), info.Uf.Find(j)
	if i == j {
		return
	}
	if info.M[i].Nunique() < info.M[j].Nunique() {
		i, j = j, i
	}

	info.Uf[j] = i
	info.M[i].Merge(info.M[j])
	info.M[j] = nil
}

func (info *Info) GetCounter(i int) *Counter {
	return info.M[info.Uf.Find(i)]
}

func (info *Info) Dec(i int, v int) {
	if info.Validv(info.A[i]) {
		info.GetCounter(i).Inc(int(info.A[i]), -1)
	}
	if info.Validv(info.A[i] - int64(v)) {
		v2 := int(info.A[i] - int64(v))
		m := info.GetCounter(i)
		if m == nil {
			m = NewCounter()
			info.M[i] = m
		}
		m.Inc(v2, 1)
		if i > 0 && info.Validv(info.A[i-1]) {
			info.Merge(i-1, i)
		}
		if i+1 < len(info.A) && info.Validv(info.A[i+1]) {
			info.Merge(i, i+1)
		}
		info.CheckUpdateH(i)
	}
}

func (info *Info) Debug() {
	fmt.Println("k=", info.K, info.H)
	for _, i := range info.Uf.GetRoots() {
		m := info.M[i]
		fmt.Println(i, m)
	}

}

func run(a []int64, queries [][]int) {
	for _, q := range queries {
		a[q[0]-1] += int64(q[1])
	}
	// fmt.Println(a)
	infos, k := []*Info{}, 1
	for k <= len(a) {
		infos = append(infos, NewInfo(a, k))
		k <<= 1
	}
	nq := len(queries)
	res := make([]int, nq)
	for i := nq - 1; i >= 0; i-- {
		ind, v := queries[i][0]-1, queries[i][1]
		for _, info := range infos {
			// fmt.Println(info)
			// info.Debug()
			res[i] = max(res[i], info.GetLongest())
			info.Dec(ind, v)
		}
		a[ind] -= int64(v)
		// fmt.Println(a)
	}
	for _, v := range res {
		fmt.Println(v)
	}

}

func main() {

	ntest := readInt()
	// ntest := 1
	for nt := 0; nt < ntest; nt++ {
		queries := make([][]int, readSliceInt()[1])
		a := readSliceInt64()
		for i := range queries {
			queries[i] = readSliceInt()
		}
		run(a, queries)
	}
}
