package main

import (
	"bufio"
	"fmt"
	"os"
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

func run(data []string) bool {
	m, n := len(data), len(data[0])
	rowhas := func(i int, ch byte) bool {
		return strings.Contains(data[i], string(ch))
	}
	colhas := func(j int, ch byte) bool {
		for i := 0; i < m; i++ {
			if data[i][j] == ch {
				return true
			}
		}
		return false
	}
	check := func(ch byte) bool {
		return rowhas(0, ch) && rowhas(m-1, ch) && colhas(0, ch) && colhas(n-1, ch)
	}
	return check('B') || check('W')
}

func main() {
	// ntest := 1
	ntest := readInt()
	// debug := ntest == 10000 && false
	// startat := 18
	for nt := 0; nt < ntest; nt++ {
		l := readSliceInt()
		data := make([]string, l[0])
		for i := 0; i < l[0]; i++ {
			data[i] = readString()
		}
		if run(data) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}

	}

}
