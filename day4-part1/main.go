package main

import (
	"crypto/md5"
	"fmt"
	"strings"
)

func main() {
	for x := 0; ; x++ {
		s := hashIt("iwrupvqb", x)
		if strings.HasPrefix(s, "000000") {
			fmt.Println(x)
			break
		}
	}
}

func hashIt(input string, i int) string {
	d := []byte(fmt.Sprintf("%s%d", input, i))
	return fmt.Sprintf("%x", md5.Sum(d))
}
