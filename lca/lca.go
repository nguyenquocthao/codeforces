package main

import (
	"fmt"
	"math/bits"
)

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

func gcd[T int | int64](a, b T) T {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

type SparseTable struct {
	data [][]int
	l    []int
}

func NewSparseTable(l []int) *SparseTable {
	r := &SparseTable{nil, l}
	n := len(l)
	data := [][]int{make([]int, n)}
	for i := 0; i < n; i++ {
		data[0][i] = l[i]
	}
	for p := 1; (1 << p) <= n; p++ {
		prev := data[p-1]

		cur := make([]int, n-(1<<p)+1)
		for i := range cur {
			cur[i] = gcd(prev[i], prev[i+(1<<(p-1))])
		}
		data = append(data, cur)
	}
	r.data = data
	return r
}

func (r *SparseTable) Query(i, j int) int {
	if i > j {
		i, j = j, i
	}
	// if i == j {
	// 	return r.l[i]
	// }
	nb := bits.Len(uint((j - i + 1))) - 1
	return gcd(r.data[nb][i], r.data[nb][j-(1<<nb)+1])
}

func main() {
	// mchild := [][]int{
	// 	{1, 2},
	// 	{3, 4},
	// 	{},
	// 	{},
	// 	{},
	// }
	n := 20
	mchild := make([][]int, n)
	for i := 1; i < n; i++ {
		mchild[i/2] = append(mchild[i/2], i)
	}
	fmt.Println(mchild)

	lca := NewLCA(0, mchild)
	// fmt.Println(lca.query(3, 4)) // Should output the LCA of nodes 3 and 4
	// fmt.Println(lca.query(1, 2)) // Should output the LCA of nodes 1 and 2
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			fmt.Println(i, j, lca.query(i, j))
		}
	}
}
