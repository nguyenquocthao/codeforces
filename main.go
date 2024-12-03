package main

import (
	"bufio"
	"fmt"
	"math/bits"
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

const MOD = 998244353

// const MOD = 1000000007
const maxn = 1000000

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

func reverse[T any](l []T) {
	i, j := 0, len(l)-1
	for i < j {
		l[i], l[j] = l[j], l[i]
		i, j = i+1, j-1
	}
}

type RMQ struct {
	data [][]int
	l    []int
}

func (r *RMQ) minAt(i, j int) int {
	if r.l[i] <= r.l[j] {
		return i
	}
	return j
}

func NewRMQ(l []int) *RMQ {
	r := &RMQ{nil, l}
	n := len(l)
	data := [][]int{make([]int, n)}
	for i := 0; i < n; i++ {
		data[0][i] = i
	}
	for p := 1; (1 << p) <= n; p++ {
		prev := data[p-1]

		cur := make([]int, n-(1<<p)+1)
		for i := range cur {
			cur[i] = r.minAt(prev[i], prev[i+(1<<(p-1))])
		}
		data = append(data, cur)
	}
	r.data = data
	return r
}

func (r *RMQ) Query(i, j int) int {
	if i > j {
		i, j = j, i
	}
	if i == j {
		return i
	}
	nb := bits.Len(uint((j - i + 1))) - 1
	return r.minAt(r.data[nb][i], r.data[nb][j-(1<<nb)+1])
}

type LCA struct {
	query func(a, b int) int
}

func NewLCA(root int, mchild [][]int) *LCA {
	nnode := len(mchild)
	order := make([]int, 0, 2*nnode)
	depth := make([]int, 0, 2*nnode)

	var dfs func(node, d int)
	dfs = func(node, d int) {
		order = append(order, node)
		depth = append(depth, d)
		for _, child := range mchild[node] {
			dfs(child, d+1)
			order = append(order, node)
			depth = append(depth, d)
		}
	}
	dfs(root, 0)

	rvIndex := make([]int, nnode)
	// for i := range rvIndex {
	// 	rvIndex[i] = -1
	// }
	for i, node := range order {
		rvIndex[node] = i
	}

	if len(order) < 32 {
		rmq := NewRMQ(depth)
		return &LCA{
			query: func(a, b int) int {
				i, j := rvIndex[a], rvIndex[b]
				return order[rmq.Query(i, j)]
			},
		}
	}

	bsize := bits.Len(uint(len(depth))) / 2
	for len(depth)%bsize > 0 {
		depth = append(depth, depth[len(depth)-1]+1)
	}

	blocks := make([]*RMQ, 0, len(depth)/bsize)
	mblock := map[int]*RMQ{}
	for i := 0; i < len(depth); i += bsize {
		key := 0
		for j := i; j < i+bsize-1; j++ {
			key <<= 1
			if depth[j] < depth[j+1] {
				key |= 1
			}
		}
		if _, ok := mblock[key]; !ok {
			mblock[key] = NewRMQ(depth[i : i+bsize])
		}
		blocks = append(blocks, mblock[key])
	}

	bQuery := func(bi, i, j int) int {
		return blocks[bi].Query(i, j) + bi*bsize
	}

	getMin := func(indexes ...int) int {
		minIdx := indexes[0]
		for _, idx := range indexes {
			if depth[idx] < depth[minIdx] {
				minIdx = idx
			}
		}
		return order[minIdx]
	}

	blockDepth := []int{}
	for bi := 0; bi < len(blocks); bi++ {
		blockDepth = append(blockDepth, depth[bQuery(bi, 0, bsize-1)])
	}
	rmqAll := NewRMQ(blockDepth)

	return &LCA{
		query: func(a, b int) int {
			i, j := rvIndex[a], rvIndex[b]
			if i > j {
				i, j = j, i
			}
			bi, bj := i/bsize, j/bsize
			if bi == bj {
				return getMin(bQuery(bi, i-bi*bsize, j-bi*bsize))
			}
			indexes := []int{
				bQuery(bi, i-bi*bsize, bsize-1),
				bQuery(bj, 0, j-bj*bsize),
			}
			if bj-bi >= 2 {
				bk := rmqAll.Query(bi+1, bj-1)
				indexes = append(indexes, bQuery(bk, 0, bsize-1))
			}
			return getMin(indexes...)
		},
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

func ListRemove[T comparable](l []T, v T) []T {
	for i, x := range l {
		if x == v {
			l[i], l[len(l)-1] = l[len(l)-1], l[i]
			return l[:len(l)-1]
		}
	}
	return l

}

func run(a []int, edges [][]int) {
	n := len(a) / 2
	m := make([][]int, 2*n)
	for _, e := range edges {
		i, j := e[0]-1, e[1]-1
		m[i] = append(m[i], j)
		m[j] = append(m[j], i)
	}
	uf := UnionFind(makeRange(0, 2*n))
	colors := make([]map[int]bool, 2*n)
	for i, col := range a {
		colors[i] = map[int]bool{col: true}
	}
	union := func(i, j int) {
		i, j = uf.Find(i), uf.Find(j)
		if i == j {
			return
		}
		if len(colors[i]) < len(colors[j]) {
			i, j = j, i
		}
		for c := range colors[j] {
			colors[i][c] = true
		}
		colors[j] = nil
		uf[j] = i
	}
	root := 0
	for i, nb := range m {
		for _, j := range nb {
			if j < i {
				union(i, j)
			}
		}
		if len(colors[uf.Find(i)]) == n {
			root = i
			break
		}
	}
	parent := make([]int, 2*n)
	parent[root] = root
	var maketree func(i int)
	maketree = func(i int) {
		for _, j := range m[i] {
			m[j] = ListRemove(m[j], i)
			parent[j] = i
			maketree(j)
		}
	}
	maketree(root)
	// fmt.Println(root, parent, m)
	mustkeep, removed, lca := make([]bool, 2*n), make([]bool, 2*n), NewLCA(root, m)
	colornodes := make([][]int, n+1)
	for i, c := range a {
		colornodes[c] = append(colornodes[c], i)
	}
	markkeep := func(i int) {
		if removed[i] {
			return
		}
		for !mustkeep[i] {
			mustkeep[i] = true
			i = parent[i]
		}
	}
	for c := 1; c <= n; c++ {
		markkeep(lca.query(colornodes[c][0], colornodes[c][1]))
	}
	// fmt.Println(colornodes, mustkeep)
	var removenode func(i int)
	removenode = func(i int) {
		if removed[i] {
			return
		}
		removed[i] = true
		for _, j := range colornodes[a[i]] {
			if j != i {
				markkeep(j)
				break
			}
		}
		for _, j := range m[i] {
			removenode(j)
		}
	}
	res := []int{}
	for i := 2*n - 1; i >= 0; i-- {
		if removed[i] {
			continue
		}
		if mustkeep[i] {
			res = append(res, i+1)
		} else {
			removenode(i)
		}
	}
	slices.Reverse(res)
	fmt.Println(len(res))
	printSlice(res)

}

func main() {
	// ntest := readInt()
	ntest := 1
	for nt := 0; nt < ntest; nt++ {
		readInt()
		a := readSliceInt()
		edges := make([][]int, len(a)-1)
		for i := range edges {
			edges[i] = readSliceInt()
		}
		run(a, edges)
	}

}
