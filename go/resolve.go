package main

import (
	"strings"
)

const (
	digits string = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_"
	base   int    = len(digits)
	expLen int    = 10 //expected length of short url
)

func shorten(id int) string {
	var result []byte
	for ; id != 0; id /= base {
		sym := id % base
		result = append(result, digits[sym])
	}
	for i := len(result) - expLen; i < 0; i++ {
		result = append(result, '0')
	}
	return string(reversed(result))
}

func resolve(short string) int {
	var result int
	short = string(reversed([]byte(short)))
	power := 0
	for _, c := range short {
		d := strings.IndexByte(digits, byte(c))
		if power == 0 {
			result += d
			power = base
		} else {
			result += d * power
			power *= base
		}
	}
	return result
}

func reversed(s []byte) []byte {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

/* func main() {
	println(shorten(8), resolve(shorten(8)))
	println(100, resolve("100"))
	println(digits, base)
	println(shorten(3969), resolve(shorten(3969)))
	println(shorten(63), resolve(shorten(63)))
	println(shorten(0), resolve(shorten(0)))
	println(shorten(64), resolve(shorten(64)))
	println(shorten(1), resolve(shorten(1)))
	println(shorten(62), resolve(shorten(62)))
	println(shorten(10), resolve(shorten(10)))
} */
