package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

var count int

func main() {
	var b bool
	flag.BoolVar(&b, "n", false, "行番号を表示")
	flag.Parse()
	fileNames := flag.Args()

	for _, fn := range fileNames {
		err := readFile(fn, b)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
		}
	}
}

func readFile(fn string, b bool) error {
	f, err := os.Open(fn)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		if b {
			count++
			fmt.Printf("%d: ", count)
		}
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}
	return nil
}
