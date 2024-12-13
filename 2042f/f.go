package main

import "slices"

func Max[T int | float32 | string | int64 | byte](args ...T) T {
	res := args[0]
	for i := 1; i < len(args); i++ {
		if args[i] > res {
			res = args[i]
		}
	}
	return res
}

// half open => 1 => 2 half open => 1.5 half open => 2
type Dt [7]int64

type SegNode struct {
	Lo, Hi      int
	Left, Right *SegNode
	A, B        int64 // in case lo!=hi, A is sum of a of the segment
	Data        Dt
}

const INF = 1 << 60

func CreateNode(a, b []int64) *SegNode {
	var create func(lo, hi int) *SegNode
	create = func(lo, hi int) *SegNode {
		if lo == hi {
			res := &SegNode{lo, hi, nil, nil, a[lo], b[lo], Dt{}}
			res.Fix()
			return res
		}
		mid := (lo + hi) / 2
		res := &SegNode{lo, hi, create(lo, mid), create(mid+1, hi), 0, 0, Dt{}}
		res.Fix()
		return res
	}
	return create(0, len(a)-1)
}

func (node *SegNode) Fix() {
	if node.Lo == node.Hi {
		node.Data = Dt{node.A + node.B, node.A + node.B, node.A + 2*node.B, -INF, -INF, -INF, -INF}
	} else {
		node.A = node.Left.A + node.Right.A
		x, y, z := node.Data, node.Left.Data, node.Right.Data
		x[0] = Max(y[0]+node.Right.A, z[0])
		x[1] = Max(y[1], node.Left.A+z[1])
		x[2] = Max(y[2], z[2], y[0]+z[1])
		x[3] = Max(y[3]+node.Right.A, y[1]+z[0], node.Left.A+z[3])
		x[4] = Max(y[4]+node.Right.A, y[2]+z[0], y[0]+z[3], z[4])
		x[5] = Max(node.Left.A+z[5], y[1]+z[2], y[3]+z[1], y[5])
		x[6] = Max(y[6], z[6], y[2]+z[2], y[0]+z[5], y[4]+z[1])
		node.Data = x
	}
	// fmt.Println(447, node)
}

func (node *SegNode) Update(i int, v int64, isa bool) {
	if node.Lo == i && node.Hi == i {
		if isa {
			node.A = v
		} else {
			node.B = v
		}
		// node.A, node.B = a, b
		node.Fix()
		return
	}
	if i <= node.Left.Hi {
		node.Left.Update(i, v, isa)
	} else {
		node.Right.Update(i, v, isa)
	}
	node.Fix()
}

func (node *SegNode) CopyData() *SegNode {
	return &SegNode{node.Lo, node.Hi, nil, nil, node.A, 0, node.Data}
}

func (node *SegNode) Query(i, j int) *SegNode {
	if i <= node.Lo && node.Hi <= j {
		return node.CopyData()
	}
	if j <= node.Left.Hi {
		return node.Left.Query(i, j)
	} else if i >= node.Right.Lo {
		return node.Right.Query(i, j)
	}
	res := &SegNode{Lo: node.Left.Lo, Hi: node.Right.Hi, Left: node.Left.Query(i, j), Right: node.Right.Query(i, j)}
	res.Fix()
	res.Left, res.Right = nil, nil
	return res
}

type Matrix [5][5]int64

func NewMatrix(a, b int) Matrix {
	x, y := int64(a), int64(b)
	res := Matrix{}
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			res[i][j] = -INF
		}
	}
	res[0][0], res[1][1], res[2][2], res[3][3], res[4][4] = 0, x, 0, x, 0
	res[0][1], res[1][2], res[2][3], res[3][4] = x+y, x+y, x+y, x+y
	res[0][2], res[2][4] = x+2*y, x+2*y
	return res
}

func CombineMatrix(a, b Matrix) Matrix {
	res := Matrix{}
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			res[i][j] = -INF
			for k := i; k <= j; k++ {
				res[i][j] = Max(res[i][j], a[i][k]+b[k][j])
			}
		}
	}
	return res
}

type SegTree struct {
	N    int
	A, B []int
	Data []Matrix
}

func NewSegTree(a, b []int) *SegTree {
	n := len(a)
	data := make([]Matrix, 2*n)
	for i := 0; i < n; i++ {
		data[i+n] = NewMatrix(a[i], b[i])
	}
	for i := n - 1; i >= 1; i-- {
		data[i] = CombineMatrix(data[2*i], data[2*i+1])
	}
	// fmt.Println(data)
	// for i, v := range data {
	// 	fmt.Println(i, v)
	// }
	return &SegTree{n, a, b, data}
}
func (s *SegTree) Update(i, v int, isa bool) {
	if isa {
		s.A[i] = v
	} else {
		s.B[i] = v
	}
	s.Data[i+s.N] = NewMatrix(s.A[i], s.B[i])
	i += s.N
	for i > 1 {
		i >>= 1
		s.Data[i] = CombineMatrix(s.Data[2*i], s.Data[2*i+1])
	}
}

func (s *SegTree) Query(i, j int) int64 {
	i, j = i+s.N, j+s.N
	chain, rvChain := []Matrix{}, []Matrix{}
	for i < j {
		if i&1 > 0 {
			chain = append(chain, s.Data[i])
			i += 1
		}
		if j&1 > 0 {
			j -= 1
			rvChain = append(rvChain, s.Data[j])
		}
		i, j = i>>1, j>>1
	}
	slices.Reverse(rvChain)
	chain = append(chain, rvChain...)
	res := chain[0]
	for i := 1; i < len(chain); i++ {
		res = CombineMatrix(res, chain[i])
	}
	return res[0][4]
}
