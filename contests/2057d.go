package main

import (
	"bufio"
	"fmt"
	"os"
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

func zfunction(s string) []int {
	n := len(s)
	z, left, right := make([]int, n), 0, 0
	for i := 1; i < n; i++ {
		z[i] = Max(0, Min(right-i, z[i-left]))
		for i+z[i] < n && s[i+z[i]] == s[z[i]] {
			z[i] += 1
		}
		if i+z[i] > right {
			left, right = i, i+z[i]
		}
	}
	return z
}

func mincomponent(s string) int {
	n := len(s)
	table, i := make([]int, n), 0
	for j := 1; j < n; j++ {
		ch := byte(s[j])
		for i > 0 && s[i] != ch {
			i = table[i-1]
		}
		if s[i] == ch {
			i += 1
		}
		table[j] = i
	}
	// fmt.Println(s, table, i)
	for n%(n-i) > 0 {
		i = table[i-1]
	}
	table = nil
	return n - i

}

type SegNode struct {
	Lo, Hi      int
	Left, Right *SegNode
	V           int
	Data        [4]int
}

func NewSegTree(a []int) *SegNode {
	var create func(lo, hi int) *SegNode
	create = func(lo, hi int) *SegNode {
		if lo == hi {
			// [min x+ax, max x+ax, min x-ax, max x-ax]
			return &SegNode{lo, hi, nil, nil, 0, [4]int{a[lo] + lo, a[lo] + lo, lo - a[lo], lo - a[lo]}}
		} else {
			mid := (lo + hi) / 2
			res := &SegNode{lo, hi, create(lo, mid), create(mid+1, hi), 0, [4]int{0, 0, 0, 0}}
			res.Fix()
			return res

		}
	}
	return create(0, len(a)-1)
}

// func BaseNode(i, v int) *SegNode {
// 	return &SegNode{i, i, nil, nil, 0, [4]int{i + v, i + v, i - v, i - v}}
// }

func (node *SegNode) Fix() {
	node.V = max(node.Left.V, node.Right.V, node.Left.Data[1]-node.Right.Data[0], node.Left.Data[3]-node.Right.Data[2])
	node.Data = [4]int{min(node.Left.Data[0], node.Right.Data[0]), max(node.Left.Data[1], node.Right.Data[1]),
		min(node.Left.Data[2], node.Right.Data[2]), max(node.Left.Data[3], node.Right.Data[3])}
}

func (node *SegNode) Update(i, v int) {
	// fmt.Println("update", i, v, node.Lo, node.Hi)
	if node.Lo == i && i == node.Hi {
		node.Data = [4]int{i + v, i + v, i - v, i - v}
		return
	}
	if i <= node.Left.Hi {
		node.Left.Update(i, v)
	} else {
		node.Right.Update(i, v)
	}
	node.Fix()

}

// func (s *SegTree) Query(i, j int) int64 {
// 	i, j = i+s.N, j+s.N
// 	chain, rvChain := []Matrix{}, []Matrix{}
// 	for i < j {
// 		if i&1 > 0 {
// 			chain = append(chain, s.Data[i])
// 			i += 1
// 		}
// 		if j&1 > 0 {
// 			j -= 1
// 			rvChain = append(rvChain, s.Data[j])
// 		}
// 		i, j = i>>1, j>>1
// 	}
// 	slices.Reverse(rvChain)
// 	chain = append(chain, rvChain...)
// 	res := chain[0]
// 	for i := 1; i < len(chain); i++ {
// 		res = CombineMatrix(res, chain[i])
// 	}
// 	return res[0][4]
// }

func run(a []int, k int) int {
	slices.Sort(a)
	n := len(a)
	dif := 0
	for i := n - 1; i > 0; i -= 2 {
		dif += a[i] - a[i-1]
	}
	added := 0
	if n%2 == 1 {
		added += a[0]
	}
	return added + max(0, dif-k)
}

func main() {
	ntest := readInt()
	// ntest := 1
	for nt := 0; nt < ntest; nt++ {
		l := readSliceInt()
		a := readSliceInt()
		fmt.Println(run(a, l[1]))

	}
}
