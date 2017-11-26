package utils

import (
	"strconv"
)

var (
	Token  string
	Length int
)

func init() {
	for i := 0; i < 10; i++ {
		Token += strconv.Itoa(i)
	}

	for i := 0; i < 26; i++ {
		Token += string('a' + i)
	}

	for i := 0; i < 26; i++ {
		Token += string('A' + i)
	}

	Length = len(Token)
}

func IdToString(id int) (res string) {
	res = ""
	for id > 0 {
		d := id % Length
		id = id / Length
		res = string(Token[d]) + res
	}
	return res
}