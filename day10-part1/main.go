// Package main provides ...
package main

import (
	"bytes"
	"fmt"
)

func main() {
	//fmt.Println(lookAndSay("1"))
	//fmt.Println(lookAndSay("11"))
	//fmt.Println(lookAndSay("21"))
	//fmt.Println(lookAndSay("1211"))
	//fmt.Println(lookAndSay("111221"))

	in := "1113122113"
	times := 50
	for i := 0; i < times; i++ {
		in = lookAndSay(in)
	}

	fmt.Println(len(in))
}

func lookAndSay(in string) string {
	out := bytes.NewBufferString("")

	for i := 0; i < len(in); {
		count := 1
		cur := in[i]
		for {
			if i >= len(in)-1 || in[i+1] != cur {
				break
			}
			count++
			i++
		}

		out.WriteString(fmt.Sprintf("%d", count))
		out.WriteString(string(cur))

		i++
	}

	return out.String()
}
