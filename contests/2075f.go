package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
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

// const MOD = 998_244_353

const MOD = 1_000_000_007

// const maxn = 200_000

// var FAC = make([]int64, maxn+1)
// var IFAC = make([]int64, maxn+1)

// func init() {
// 	FAC[0], FAC[1] = 1, 1
// 	IFAC[0], IFAC[1] = 1, 1
// 	for i := int64(2); i < maxn+1; i++ {
// 		FAC[i] = (i * FAC[i-1]) % MOD
// 		IFAC[i] = mod_inverse(FAC[i])
// 	}
// }

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

// func comb(n, k int64) int64 {
// 	if n < 0 || k > n {
// 		return 0
// 	}
// 	inv := (IFAC[k] * IFAC[n-k]) % MOD
// 	return (FAC[n] * inv) % MOD
// }

func mod[T int | int64](v T) T {
	res := v % MOD
	if res < 0 {
		res += MOD
	}
	return res
}

// type UnionFind []int

// func (u UnionFind) Find(x int) int {
// 	if u[x] != x {
// 		u[x] = u.Find(u[x])
// 	}
// 	return u[x]
// }

// func (u UnionFind) Union(x, y int) {
// 	u[u.Find(y)] = u.Find(x)
// }

// func (u UnionFind) NGroup() int {
// 	res := 0
// 	for i, v := range u {
// 		if i == v {
// 			res++
// 		}
// 	}
// 	return res
// }

// func (u UnionFind) GetRoots() []int {
// 	res := make([]int, 0, len(u))
// 	for i, v := range u {
// 		if i == v {
// 			res = append(res, i)
// 		}
// 	}
// 	return res
// }

// type Item struct {
// 	Start, End int
// }
// type Heap []Item

// func (h Heap) Len() int           { return len(h) }
// func (h Heap) Less(i, j int) bool { return (h[i].End - h[i].Start) < (h[j].End - h[j].Start) }
// func (h Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// func (h *Heap) Push(x any) {
// 	// Push and Pop use pointer receivers because they modify the slice's length,
// 	// not just its contents.
// 	*h = append(*h, x.(Item))
// }

// func (h *Heap) Pop() any {
// 	old := *h
// 	n := len(old)
// 	x := old[n-1]
// 	*h = old[0 : n-1]
// 	return x
// }

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

type Item struct {
	A0, A1, B0, B1 int
}

type Node struct {
	Lo, Hi      int
	Left, Right *Node
	Lazy, Maxv  int
}

func (node *Node) Fix() {
	if node.Left != nil {
		node.Maxv = node.Lazy + max(node.Left.Maxv, node.Right.Maxv)
	}
}

func (node *Node) Flush() {
	if node.Left != nil {
		node.Left.Lazy += node.Lazy
		node.Left.Maxv += node.Lazy
		node.Right.Lazy += node.Lazy
		node.Right.Maxv += node.Lazy
	}
	node.Lazy = 0
}

func createnode(lo, hi int) *Node {
	if lo == hi {
		return &Node{lo, hi, nil, nil, 0, 0}
	}
	mid := (lo + hi) / 2
	return &Node{lo, hi, createnode(lo, mid), createnode(mid+1, hi), 0, 0}
}

func (node *Node) Update(l, r, v int) {
	if node.Hi < l || node.Lo > r {
		return
	}
	if l <= node.Lo && node.Hi <= r {
		node.Maxv += v
		node.Lazy += v
		return
	}
	node.Flush()
	node.Left.Update(l, r, v)
	node.Right.Update(l, r, v)
	node.Fix()

}

func run(arr []int) int {
	n := len(arr)
	noninc := true
	for i := 1; i < n; i++ {
		if arr[i-1] < arr[i] {
			noninc = false
			break
		}
	}
	if noninc {
		return 1
	}
	spre, ssuf := NewStack[int](), NewStack[int]()
	spre.Push(0)
	for i, v := range arr {
		if arr[spre.Top()] > v {
			spre.Push(i)
		}
		for ssuf.Len() > 0 && arr[ssuf.Top()] <= v {
			ssuf.Pop()
		}
		ssuf.Push(i)
	}
	pre, suf := spre.ToList(), ssuf.ToList()
	items := make([]Item, n)
	ai, bi := -1, 0
	for i := 0; i < n; i++ {
		if suf[bi] == i {
			bi += 1
		}
		if ai+1 < len(pre) && pre[ai+1] < i {
			ai += 1
		}
		items[i].A1 = ai
		items[i].B0 = bi
	}
	indexes := makeRange(0, n)
	slices.SortFunc(indexes, func(i, j int) int {
		return arr[i] - arr[j]
	})
	slices.Reverse(indexes)
	ai, bi = 0, -1
	for _, i := range indexes {
		v := arr[i]
		// v decreasing
		if ai < len(pre) && arr[pre[ai]] == v {
			ai += 1
		}
		if bi+1 < len(suf) && arr[suf[bi+1]] > v {
			bi += 1
		}
		items[i].A0 = ai
		items[i].B1 = bi
	}
	// fmt.Println(pre, suf)
	// fmt.Println(items)

	acc := make([][]Item, len(suf)+1)
	for _, it := range items {
		if it.A0 > it.A1 || it.B0 > it.B1 {
			continue
		}
		acc[it.B0] = append(acc[it.B0], Item{A0: it.A0, A1: it.A1, B0: 1})
		acc[it.B1+1] = append(acc[it.B1+1], Item{A0: it.A0, A1: it.A1, B0: -1})
	}
	// fmt.Println(acc)
	node := createnode(0, len(pre)+1)
	res := 0
	for _, l := range acc {
		for _, it := range l {
			node.Update(it.A0, it.A1, it.B0)
		}
		res = max(res, node.Maxv)
	}

	return res + 2

}

var debug = false

func main() {
	defer func() {
		if r := recover(); r != nil {
			// fmt.Println(data)

			buf := make([]byte, 1024)
			n := runtime.Stack(buf, false)
			// fmt.Println("Panic recovered:", r)

			fmt.Println(strings.ReplaceAll(string(buf[:n]), "\n", " "))
		}
	}()
	// cached := make([]int64, 100_000+1)
	// inv2 := mod_inverse(2)
	// for i := 2; i < len(cached); i++ {
	// 	cached[i] = mod(1 + mod(cached[i/2]+cached[(i+1)/2])*inv2)
	// }
	ntest := readInt()
	// debug = ntest == 10000
	for nt := 0; nt < ntest; nt++ {
		readInt()
		fmt.Println(run(readSliceInt()))
	}
}
