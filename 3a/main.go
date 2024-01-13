package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

var (
	numberReg = regexp.MustCompile("([0-9]+)")
)

func Hash(s string) byte {
	hash := byte(0)
	for i, b := range s {
		fmt.Println(i, b)
	}
	return hash
}

func main() {

	buf, err := os.ReadFile("sample")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(buf), "\n")
	fmt.Println(lines)
	for _, line := range lines {
		if len(line) == 0 {
			fmt.Println("empty line")
		}
		if len(line) < 1 {
			continue
		}
	}
	fmt.Println("----")
}
