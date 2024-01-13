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
	hash := 0
	for _, b := range s {
		hash += int(b)
		hash = hash * 17
		hash = hash % 256
	}
	return byte(hash)
}

func HashSum(input string) int {
	sum := 0
	for _, s := range strings.Split(input, ",") {
		sum += int(Hash(s))
	}
	return sum
}

func main() {

	buf, err := os.ReadFile("input")
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
		fmt.Println("Hashsum:", HashSum(line))
	}
	fmt.Println("----")
}
