package utils

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

var rdr *bufio.Reader

func Readline() string {
	buf := make([]byte, 0, 16)
	for {
		l, p, e := rdr.ReadLine()
		if e != nil {
			fmt.Println(e.Error())
			panic(e)
		}
		buf = append(buf, l...)
		if !p {
			break
		}
	}
	return string(buf)
}

func ReadInt() int {
	return String2Int(Readline())
}

func ReadIntSlice() []int {
	slice := make([]int, 0)
	lines := strings.Split(Readline(), " ")
	for _, line := range lines {
		slice = append(slice, String2Int(line))
	}
	return slice
}

func String2Int(s string) int {
	v, ok := strconv.Atoi(s)
	if ok != nil {
		panic(ok)
	}
	return v
}

func Int2String(i int) string {
	return strconv.Itoa(i)
}
