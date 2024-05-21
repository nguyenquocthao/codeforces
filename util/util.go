package util

import (
	"bufio"
	"fmt"
	"math"
	"math/big"
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

func readString() string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func readSliceInt() []int {
	data := strings.Split(readString(), " ")
	res := make([]int, len(data))
	for i, v := range data {
		res[i], _ = strconv.Atoi(v)
	}
	return res
}

func readSliceInt64() []int64 {
	data := strings.Split(readString(), " ")
	res := make([]int64, len(data))
	for i, v := range data {
		res[i], _ = strconv.ParseInt(v, 10, 64)
	}
	return res
}

func printSlice[T any](l []T) {
	output := make([]string, len(l))
	for i, v := range l {
		output[i] = fmt.Sprint(v)
	}
	fmt.Println(strings.Join(output, " "))
}

func Max[T int | float32 | string | int64](args ...T) T {
	res := args[0]
	for i := 1; i < len(args); i++ {
		if args[i] > res {
			res = args[i]
		}
	}
	return res
}

func Min[T int | float32 | string | int64](args ...T) T {
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

func abs(v int) int {
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

func StackPop[T any](st []T) ([]T, T) {
	w := st[len(st)-1]
	st = st[:len(st)-1]
	return st, w
}

func Tarjan[T comparable](nodes []T, getEdges func(T) []T) map[T]int {
	lowLinks := map[T]int{}
	onstack := map[T]bool{}
	stack := []T{}

	var strongconnect func(T)
	strongconnect = func(v T) {
		if _, ok := lowLinks[v]; ok {
			return
		}
		index := len(lowLinks)
		lowLinks[v] = index
		stack = append(stack, v)
		onstack[v] = true
		for _, w := range getEdges(v) {
			strongconnect(w)
			if onstack[w] {
				lowLinks[v] = Min(lowLinks[v], lowLinks[w])
			}
		}
		if lowLinks[v] == index {
			for {
				var w T
				stack, w = StackPop(stack)
				delete(onstack, w)
				if w == v {
					break
				}
			}
		}
	}
	for _, v := range nodes {
		strongconnect(v)
	}

	return lowLinks
}

func GetDAGFromSCS[T comparable](nodes []T, getNb func(T) []T, sccs map[T]int) map[int][]int {
	marked, connected := map[T]bool{}, map[[2]int]bool{}
	var dp func(T)
	dp = func(i T) {
		if marked[i] {
			return
		}
		marked[i] = true
		for _, j := range getNb(i) {
			key := [2]int{sccs[i], sccs[j]}
			if key[0] != key[1] {
				connected[key] = true
			}

			dp(j)
		}
	}
	for _, i := range nodes {
		dp(i)
	}
	res := map[int][]int{}
	for pair := range connected {
		res[pair[0]] = append(res[pair[0]], pair[1])
	}
	return res
}

const inf int64 = 1 << 62

var biinf = big.NewInt(inf)

const MOD = 998244353

// const MOD = 1000000007
const maxn = 1000000

var FAC = make([]int64, maxn+1)
var IFAC = make([]int64, maxn+1)

func init() {
	FAC[0], FAC[1] = 1, 1
	IFAC[0], IFAC[1] = 1, 1
	for i := int64(2); i < maxn+1; i++ {
		FAC[i] = (i * FAC[i-1]) % MOD
		IFAC[i] = mod_inverse(FAC[i])
	}
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

func comb(n, k int64) int64 {
	if n < 0 || k > n {
		return 0
	}
	inv := (IFAC[k] * IFAC[n-k]) % MOD
	return (FAC[n] * inv) % MOD
}

func mod(v int64) int64 {
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

// copied from assert.Len
func GetLen(x interface{}) (length int, ok bool) {
	v := reflect.ValueOf(x)
	defer func() {
		if e := recover(); e != nil {
			ok = false
		}
	}()
	return v.Len(), true
}

func IsTrue(variable any) bool {
	if variable == nil {
		return false
	}

	v := reflect.ValueOf(variable)
	if v.IsZero() {
		return false
	}

	k := v.Kind()
	if k == reflect.Pointer {
		if v.IsNil() {
			return false
		}
		variable = v.Elem().Interface()
	}

	length, ok := GetLen(variable)
	if ok {
		return length > 0
	}

	switch z := variable.(type) {
	case bool:
		return z
	case string:
		return len(z) > 0
	case int:
		return z != 0

	}
	return true
}

func CountTrue(l []any) int {
	res := 0
	for _, v := range l {
		if IsTrue(v) {
			res += 1
		}
	}
	return res

}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

// func Sum[T bool | int | int64 | float64](l []T) T {
// 	var res T
// 	for _, v := range l {
// 		res += v
// 	}
// }

func toInt(v string) int {
	res, _ := strconv.Atoi(v)
	return res
}

func makeArray[T any](v T, n int) []T {
	res := make([]T, n)
	for i := 0; i < n; i++ {
		res[i] = v
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

func factorize(v int) []int {
	lo, hi := []int{}, []int{}
	for i := 1; i*i <= v; i++ {
		if i*i == v {
			lo = append(lo, i)
			break
		} else if v%i == 0 {
			lo = append(lo, i)
			hi = append(hi, v/i)
		}
	}
	for j := len(hi) - 1; j >= 0; j-- {
		lo = append(lo, hi[j])
	}
	return lo
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

func Keys[K comparable, V any](m map[K]V) []K {
	res := []K{}
	for k := range m {
		res = append(res, k)
	}
	return res
}

func divceil(a, b int) int {
	res := a / b
	if a%b > 0 {
		res += 1
	}
	return res
}

func cal2(a, b int) int {
	// 	def cal(a,b):
	//     c=0
	//     while b>0:
	//         b-=a
	//         if b<=0: break
	//         c+=b
	//     return c
	// def cal2(a,b):
	//     n = b // a  # This is floor(b/a)
	//     total = (n * (2*b - a * (n + 1))) // 2
	//     return total
	n := b / a
	// b-a => b-a*n
	return (n * (2*b - a*(n+1))) / 2
}

func next_greater_index(arr []int) []int {
	n := len(arr)
	res := make([]int, n)
	st := []int{}
	for i, x := range arr {
		res[i] = n
		for len(st) > 0 && arr[st[len(st)-1]] < x {
			res[st[len(st)-1]] = i
			st = st[:len(st)-1]
		}
		st = append(st, i)
	}
	return res
}

func mainwithtestcase() {
	ntest := readInt()

	debug := ntest == 10000 && false
	logtest := 495

	// ntest := 1

	for nt := 0; nt < ntest; nt++ {
		readInt()
		a := readSliceInt()
		// l := readSliceInt()
		// res := run(a)
		res := 1
		// fmt.Println()
		if !debug {
			fmt.Println(res)
		} else {
			if nt >= logtest {
				fmt.Println(nt, a)
				fmt.Println(res)
			}

		}

	}

}

func Unique[T comparable](l []T) []T {
	m := map[T]bool{}
	for _, v := range l {
		m[v] = true
	}
	res := []T{}
	for v := range m {
		res = append(res, v)
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
