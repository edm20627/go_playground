package main

import (
	"fmt"
)

func main() {
	fs := make([]func(), 3)
	for i := range fs {
		// これがないと結果は[2, 2, 2]
		// iのポインタが同じなため
		// {}のスコープで再定義する
		// クロージャ
		i := i
		fs[i] = func() { fmt.Println(i) }
	}
	for _, f := range fs {
		f()
	}
}
