package main

import (
	"bufio"
	"fmt"
	"math"
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

// const MOD = 1000000007

// const MOD = (1 << 60) + 7

// const maxn = 1000000

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

// func pow(x, n int64) int64 {
// 	x = x % MOD
// 	res := int64(1)
// 	for n > 0 {
// 		if n%2 == 1 {
// 			res = (res * x) % MOD
// 		}
// 		x = (x * x) % MOD
// 		n = n / 2
// 	}
// 	return res
// }

// func mod_inverse(x int64) int64 {
// 	return pow(x, MOD-2)
// }

// func comb(n, k int64) int64 {
// 	if n < 0 || k > n {
// 		return 0
// 	}
// 	inv := (IFAC[k] * IFAC[n-k]) % MOD
// 	return (FAC[n] * inv) % MOD
// }

// func mod[T int | int64](v T) T {
// 	res := v % MOD
// 	if res < 0 {
// 		res += MOD
// 	}
// 	return res
// }

func getlex(a []int) []int {
	maxv := Max(a...)
	target := 2*maxv + 1
	canrm := Sum(a...) - target
	if canrm < 0 {
		return []int{}
	}
	st := NewStack[int]()
	for _, v := range a {
		for st.Len() > 0 && st.Top() < v && st.Top() <= canrm {
			canrm -= st.Pop()
		}
		st.Push(v)
	}
	return st.ToList()
}

func lexmax(a, b []int) []int {
	if len(a) > len(b) {
		a, b = b, a
	}
	for i, v := range a {
		if v < b[i] {
			return b
		} else if v > b[i] {
			return a
		}
	}
	return b
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

func nbitat(x int64, i int) int64 {
	p := int64(1) << i
	a, b := x/(2*p), x%(2*p)
	return a*p + max(0, b-p+1)
}

func comb2(n int64) int64 {
	return mod(n * (n - 1) / 2)
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

type Heap []int

func (h Heap) Len() int           { return len(h) }
func (h Heap) Less(i, j int) bool { return h[i] < h[j] }
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

const inf int64 = 1 << 62
const v9 = 2_000_000_000

func sqrt(v int64) int64 {
	res := int64(math.Sqrt(float64(v)))
	for res*res < v {
		res += 1
	}
	for res*res > v {
		res -= 1
	}
	return res
}

type Item struct {
	V     int64
	Thres int64
}

func NewItem(data []int64) Item {
	v := int64(0)
	for _, x := range data {
		v = sqrt(v + x)
	}
	thres := v + 1
	for i := len(data) - 1; i >= 0; i-- {
		x := data[i]
		if thres >= v9 {
			thres = inf
			break
		} else {
			thres = thres*thres - x
		}
	}
	return Item{v, thres}
}

type Node struct {
	Item
	Lo, Hi      int
	Left, Right *Node
}

func (node *Node) Fix() {
	if node.Left == nil {
		return
	}
	if node.Left.V < node.Right.Thres-1 {
		node.V = node.Right.V
		node.Thres = inf
	} else if node.Left.V == node.Right.Thres-1 {
		node.V = node.Right.V
		node.Thres = node.Left.Thres
	} else {
		node.V = node.Right.V + 1
		node.Thres = inf
	}
}

func (node *Node) Update(i int, it Item) {
	if node.Lo == i && node.Hi == i {
		node.Item = it
		return
	}
	if node.Hi < i || node.Lo > i {
		return
	}
	node.Left.Update(i, it)
	node.Right.Update(i, it)
	node.Fix()

}

type NodeRoot struct {
	Data   []int64
	Offset int
	Root   *Node
}

func NewNodeRoot(a []int64) *NodeRoot {
	offset := 6 - (len(a) % 6)
	if offset == 6 {
		offset = 0
	} else {
		a = append(make([]int64, offset), a...)
	}
	var dp func(i, j int) *Node
	dp = func(i, j int) *Node {
		if i == j {
			return &Node{Item: NewItem(a[i*6 : i*6+6]), Lo: i, Hi: i, Left: nil, Right: nil}
		}
		mid := (i + j) / 2
		node := &Node{Lo: i, Hi: j, Left: dp(i, mid), Right: dp(mid+1, j)}
		node.Fix()
		return node
	}
	return &NodeRoot{a, offset, dp(0, len(a)/6-1)}
}

func (root *NodeRoot) Update(i, v int64) int64 {
	i = int64(root.Offset) + i - 1
	root.Data[i] = v
	ni := int(i / 6)
	root.Root.Update(ni, NewItem(root.Data[ni*6:ni*6+6]))
	return root.Root.V
}

func run(a []int64, queries [][]int64) {
	root := NewNodeRoot(a)
	for _, q := range queries {
		fmt.Println(root.Update(q[0], q[1]))
	}
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

	// ntest := readInt()
	ntest := 1
	// debug = ntest == 10000
	for nt := 0; nt < ntest; nt++ {
		queries := make([][]int64, readSliceInt()[1])
		a := readSliceInt64()
		for i := range queries {
			queries[i] = readSliceInt64()
		}
		run(a, queries)
	}
}
