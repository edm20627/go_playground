package main

import (
	"fmt"
	"strings"
)

func main() {
	counts := map[string]int{}

	str := "hoeg hoge fuga hoge fuga"
	for _, s := range strings.Split(str, " ") {
		counts[s]++
	}
	fmt.Println(counts)
}
