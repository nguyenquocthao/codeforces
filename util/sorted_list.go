package util

import (
	"math/bits"
	"sort"
)

type FenwickTree2 struct {
	bit  []int
	size int
}

func NewFenwickTree2(x []int) *FenwickTree2 {
	ft := &FenwickTree2{
		bit:  make([]int, len(x)),
		size: len(x),
	}
	copy(ft.bit, x)
	for i := 0; i < ft.size; i++ {
		j := i | (i + 1)
		if j < ft.size {
			ft.bit[j] += ft.bit[i]
		}
	}
	return ft
}

func (ft *FenwickTree2) Update(idx, x int) {
	for idx < ft.size {
		ft.bit[idx] += x
		idx |= idx + 1
	}
}

func (ft *FenwickTree2) Query(end int) int {
	x := 0
	for end > 0 {
		x += ft.bit[end-1]
		end &= end - 1
	}
	return x
}

func (ft *FenwickTree2) FindKth(k int) (int, int) {
	idx := -1
	for d := bits.Len(uint(ft.size)) - 1; d >= 0; d-- {
		rightIdx := idx + (1 << d)
		if rightIdx < ft.size && ft.bit[rightIdx] <= k {
			idx = rightIdx
			k -= ft.bit[idx]
		}
	}
	return idx + 1, k
}

type SortedList struct {
	blockSize int
	micros    [][]int
	macro     []int
	microSize []int
	fenwick   *FenwickTree2
	size      int
}

func NewSortedList(iterable ...int) *SortedList {
	sort.Ints(iterable)
	blockSize := 700
	micros := [][]int{}
	for i := 0; i < len(iterable); i += blockSize - 1 {
		end := i + blockSize - 1
		if end > len(iterable) {
			end = len(iterable)
		}
		micros = append(micros, iterable[i:end])
	}
	if len(micros) == 0 {
		micros = append(micros, []int{})
	}
	macro := make([]int, len(micros)-1)
	for i := 1; i < len(micros); i++ {
		macro[i-1] = micros[i][0]
	}
	microSize := make([]int, len(micros))
	for i, micro := range micros {
		microSize[i] = len(micro)
	}
	fenwick := NewFenwickTree2(microSize)
	return &SortedList{
		blockSize: blockSize,
		micros:    micros,
		macro:     macro,
		microSize: microSize,
		fenwick:   fenwick,
		size:      len(iterable),
	}
}

func lower_bound(a []int, x int) int {
	return sort.Search(len(a), func(i int) bool {
		return a[i] >= x
	})
}

func upper_bound(a []int, x int) int {
	return sort.Search(len(a), func(i int) bool {
		return a[i] > x
	})
}

func (sl *SortedList) Insert(x int) {
	i := lower_bound(sl.macro, x)
	j := upper_bound(sl.micros[i], x)
	sl.micros[i] = append(sl.micros[i][:j], append([]int{x}, sl.micros[i][j:]...)...)
	sl.size++
	sl.microSize[i]++
	sl.fenwick.Update(i, 1)
	if len(sl.micros[i]) >= sl.blockSize {
		mid := sl.blockSize >> 1
		newMicro := make([]int, len(sl.micros[i])-mid)
		copy(newMicro, sl.micros[i][mid:])
		sl.micros[i] = sl.micros[i][:mid]
		sl.micros = append(sl.micros[:i+1], append([][]int{newMicro}, sl.micros[i+1:]...)...)
		sl.microSize[i] = mid
		sl.microSize = append(sl.microSize[:i+1], append([]int{len(newMicro)}, sl.microSize[i+1:]...)...)

		sl.fenwick = NewFenwickTree2(sl.microSize)
		sl.macro = append(sl.macro[:i], append([]int{sl.micros[i+1][0]}, sl.macro[i:]...)...)
	}
}

func (sl *SortedList) Pop(k int) int {
	i, j := sl.findKth(k)
	sl.size--
	sl.microSize[i]--
	sl.fenwick.Update(i, -1)
	popped := sl.micros[i][j]
	sl.micros[i] = append(sl.micros[i][:j], sl.micros[i][j+1:]...)
	return popped
}

func (sl *SortedList) Delete(x int) {
	sl.Pop(sl.lowerBound(x))
}

func (sl *SortedList) lowerBound(x int) int {
	i := lower_bound(sl.macro, x)
	return sl.fenwick.Query(i) + lower_bound(sl.micros[i], x)
}

func (sl *SortedList) upperBound(x int) int {
	i := upper_bound(sl.macro, x)
	return sl.fenwick.Query(i) + upper_bound(sl.micros[i], x)
}

func (sl *SortedList) findKth(k int) (int, int) {
	if k < 0 {
		k += sl.size
	}
	return sl.fenwick.FindKth(k)
}

func (sl *SortedList) Len() int {
	return sl.size
}

func (sl *SortedList) At(k int) int {
	i, j := sl.findKth(k)
	return sl.micros[i][j]
}

func (sl *SortedList) ToList() []int {
	res := []int{}
	for _, l := range sl.micros {
		res = append(res, l...)
	}
	return res
}
