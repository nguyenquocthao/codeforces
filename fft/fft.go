package fft

import (
	"fmt"
	"math"
	"math/bits"
	"math/cmplx"
)

const MODNTT = 998244353
const INV2 = 499122177

var cachedNTT = map[int]int{}

func modpos[T int | int64](x T, mod int) int {
	res := x % T(mod)
	if res < 0 {
		res += T(mod)
	}
	return int(res)
}

func powntt(x, n int) int {
	x = modpos(x, MODNTT)
	n = modpos(n, MODNTT-1)

	if x == 3 {
		if res, ok := cachedNTT[n]; ok {
			return res
		}
	}

	res := int64(1)
	for n > 0 {
		if n%2 == 1 {
			res = (res * int64(x)) % MODNTT
		}
		x = (x * x) % MODNTT
		n = n / 2
	}
	if x == 3 {
		cachedNTT[n] = int(res)
	}
	return int(res)
}

func modn(n, i int) int {
	p := ((MODNTT - 1) / n) * i
	if p < 0 {
		p += MODNTT - 1
	}
	return powntt(3, p)
}

func mod2[T int | int64](v T) int {
	return modpos(v*INV2, MODNTT)
}

func ntt(x []int) []int {
	n := len(x)
	if n == 1 {
		return x
	} else if n == 2 {
		return []int{modpos(x[0]+x[1], MODNTT), modpos(x[0]-x[1], MODNTT)}
	}
	even, odd := make([]int, n/2), make([]int, n/2)
	for i, v := range x {
		if i%2 == 0 {
			even[i/2] = v
		} else {
			odd[i/2] = v
		}
	}
	even, odd = ntt(even), ntt(odd)
	for i, v := range odd {
		x := modpos(int64(v)*int64(modn(n, i)), MODNTT)
		odd[i] = int(x)
	}
	res := make([]int, n)
	for i, _ := range res {
		if i < n/2 {
			res[i] = modpos(even[i]+odd[i], MODNTT)
		} else {
			res[i] = modpos(even[i-n/2]-odd[i-n/2], MODNTT)
		}
	}
	return res
}

func intt(x []int) []int {
	n := len(x)
	if n == 1 {
		return x
	} else if n == 2 {
		return []int{mod2(x[0] + x[1]), mod2(x[0] - x[1])}
	}
	even, odd := make([]int, n/2), make([]int, n/2)
	for i, v := range x {
		if i%2 == 0 {
			even[i/2] = v
		} else {
			odd[i/2] = v
		}
	}
	even, odd = intt(even), intt(odd)
	for i, v := range odd {
		x := modpos(int64(v)*int64(modn(n, -i)), MODNTT)
		odd[i] = int(x)
	}
	res := make([]int, n)
	for i, _ := range res {
		if i < n/2 {
			res[i] = mod2(even[i] + odd[i])
		} else {
			res[i] = mod2(even[i-n/2] - odd[i-n/2])
		}
	}
	return res
}

// a and b are different slices
func convolventt(a []int, b []int) []int {
	n := len(a) + len(b) - 1
	pow2 := 1 << (bits.Len(uint(n)) - 1)
	if pow2 < n {
		n = 2 * pow2
	}
	a, b = append(a, make([]int, n-len(a))...), append(b, make([]int, n-len(b))...)

	x, y := a, b
	x, y = ntt(x), ntt(y)
	for i, v := range x {
		x[i] = modpos(int64(v)*int64(y[i]), MODNTT)
	}
	return intt(x)
}

/// https://github.com/argusdusty/gofft/blob/v1.2.1/convolve.go#L10

// Convolve computes the discrete convolution of x and y using FFT.
// Pads x and y to the next power of 2 from len(x)+len(y)-1
func Convolve(x, y []complex128) ([]complex128, error) {
	if len(x) == 0 && len(y) == 0 {
		return nil, nil
	}
	n := len(x) + len(y) - 1
	N := NextPow2(n)
	x = ZeroPad(x, N)
	y = ZeroPad(y, N)
	err := FastConvolve(x, y)
	return x[:n], err
}

// FastConvolve computes the discrete convolution of x and y using FFT
// and stores the result in x, while erasing y (setting it to 0s).
// Since this does no allocations, x and y are assumed to already be 0-padded
// for at least half their length.
func FastConvolve(x, y []complex128) error {
	if len(x) == 0 && len(y) == 0 {
		return nil
	}
	if err := checkZero("difference in FastConvolve input vectors length", len(x)-len(y)); err != nil {
		return err
	}
	if err := checkLength("FastConvolve input vector length", len(x)); err != nil {
		return err
	}
	convolve(x, y)
	return nil
}

// convolve does the actual work of convolutions.
func convolve(x, y []complex128) {
	fft(x)
	fft(y)
	for i := 0; i < len(x); i++ {
		x[i] *= y[i]
		y[i] = 0
	}
	ifft(x)
}

// IsPow2 returns true if N is a perfect power of 2 (1, 2, 4, 8, ...) and false otherwise.
// Algorithm from: https://graphics.stanford.edu/~seander/bithacks.html#DetermineIfPowerOf2
func IsPow2(N int) bool {
	if N == 0 {
		return false
	}
	return (uint64(N) & uint64(N-1)) == 0
}

// NextPow2 returns the smallest power of 2 >= N.
func NextPow2(N int) int {
	if N == 0 {
		return 1
	}
	return 1 << uint64(bits.Len64(uint64(N-1)))
}

// ZeroPad pads x with 0s at the end into a new array of length N.
// This does not alter x, and creates an entirely new array.
// This should only be used as a convience function, and isn't meant for performance.
// You should call this as few times as possible since it does potentially large allocations.
func ZeroPad(x []complex128, N int) []complex128 {
	y := make([]complex128, N)
	copy(y, x)
	return y
}

// InputSizeError represents an error when an input vector's size is not a power of 2.
type InputSizeError struct {
	Context     string
	Requirement string
	Size        int
}

func (e *InputSizeError) Error() string {
	return fmt.Sprintf("Size of %s must be %s, is: %d", e.Context, e.Requirement, e.Size)
}

// checkLength checks that the length of x is a valid power of 2
func checkLength(Context string, N int) error {
	if !IsPow2(N) {
		return &InputSizeError{Context: Context, Requirement: "power of 2", Size: N}
	}
	return nil
}

// checkLength checks that the length of x is a valid power of 2
func checkZero(Context string, N int) error {
	if N != 0 {
		return &InputSizeError{Context: Context, Requirement: "zero", Size: N}
	}
	return nil
}

// fft does the actual work for FFT
func fft(x []complex128) {
	N := len(x)
	// Handle small N quickly
	switch N {
	case 1:
		return
	case 2:
		x[0], x[1] = x[0]+x[1], x[0]-x[1]
		return
	case 4:
		f := complex(imag(x[1])-imag(x[3]), real(x[3])-real(x[1]))
		x[0], x[1], x[2], x[3] = x[0]+x[1]+x[2]+x[3], x[0]-x[2]+f, x[0]-x[1]+x[2]-x[3], x[0]-x[2]-f
		return
	}
	// Reorder the input array.
	permute(x)
	// Butterfly
	// First 2 steps
	for i := 0; i < N; i += 4 {
		f := complex(imag(x[i+2])-imag(x[i+3]), real(x[i+3])-real(x[i+2]))
		x[i], x[i+1], x[i+2], x[i+3] = x[i]+x[i+1]+x[i+2]+x[i+3], x[i]-x[i+1]+f, x[i]-x[i+2]+x[i+1]-x[i+3], x[i]-x[i+1]-f
	}
	// Remaining steps
	w := complex(0, -1)
	for n := 4; n < N; n <<= 1 {
		w = cmplx.Sqrt(w)
		for o := 0; o < N; o += (n << 1) {
			wj := complex(1, 0)
			for k := 0; k < n; k++ {
				i := k + o
				f := wj * x[i+n]
				x[i], x[i+n] = x[i]+f, x[i]-f
				wj *= w
			}
		}
	}
}

// ifft does the actual work for IFFT
func ifft(x []complex128) {
	N := len(x)
	// Reverse the input vector
	for i := 1; i < N/2; i++ {
		j := N - i
		x[i], x[j] = x[j], x[i]
	}

	// Do the transform.
	fft(x)

	// Scale the output by 1/N
	invN := complex(1.0/float64(N), 0)
	for i := 0; i < N; i++ {
		x[i] *= invN
	}
}

// permutate permutes the input vector using bit reversal.
// Uses an in-place algorithm that runs in O(N) time and O(1) additional space.
func permute(x []complex128) {
	N := len(x)
	// Handle small N quickly
	switch N {
	case 1, 2:
		return
	case 4:
		x[1], x[2] = x[2], x[1]
		return
	case 8:
		x[1], x[3], x[4], x[6] = x[4], x[6], x[1], x[3]
		return
	}
	shift := 64 - uint64(bits.Len64(uint64(N-1)))
	N2 := N >> 1
	for i := 0; i < N; i += 2 {
		ind := int(bits.Reverse64(uint64(i)) >> shift)
		// Skip cases where low bit isn't set while high bit is
		// This eliminates 25% of iterations
		if i < N2 {
			if ind > i {
				x[i], x[ind] = x[ind], x[i]
			}
		}
		ind |= N2 // Fast way to get int(bits.Reverse64(uint64(i+1)) >> shift) here
		if ind > i+1 {
			x[i+1], x[ind] = x[ind], x[i+1]
		}
	}
}

func ConvolveInt(a []int, b []int) []int {
	tocomplex := func(a []int) []complex128 {
		y := make([]complex128, len(a))
		for i, v := range a {
			y[i] = complex(float64(v), 0)
		}
		return y
	}
	toint := func(a []complex128) []int {
		y := make([]int, len(a))
		for i, v := range a {
			y[i] = int(math.Round(real(v)))
		}
		return y
	}

	res, err := Convolve(tocomplex(a), tocomplex(b))
	if err != nil {
		panic(err)
	}
	return toint(res)
}

// . is any character
func SearchStr(a, b string) []int {
	if len(a) < len(b) {
		return []int{}
	}
	chars := map[rune]bool{}
	for _, v := range a {
		chars[v] = true
	}
	n, m := len(a), len(b)
	indexes := make([]bool, n-m+1)
	for i := 0; i < len(indexes); i++ {
		indexes[i] = true
	}
	x, y := make([]int, n), make([]int, m)
	// fmt.Println(256, chars)
	for curch := range chars {
		if curch == '.' {
			continue
		}
		// fmt.Println(string(curch))

		for i, v := range a {
			if v == curch {
				x[i] = 1
			} else {
				x[i] = 0
			}
		}
		for i, v := range b {
			if v != '.' && v != curch {
				y[m-1-i] = 1
			} else {
				y[m-1-i] = 0
			}
		}
		z := ConvolveInt(x, y)
		// z := convolventt(x, y)
		for i, v := range indexes {
			if v && z[i+m-1] != 0 {
				indexes[i] = false
			}
		}
	}
	res := []int{}
	for i, v := range indexes {
		if v {
			res = append(res, i)
		}
	}
	// fmt.Println("searchstr", a, b, indexes)
	return res

}
