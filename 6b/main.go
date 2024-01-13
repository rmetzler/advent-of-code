package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func countPossibilities(time, distance int) int {
	count := 0
	for i := 1; i <= time-1; i++ {
		remaining := time - i
		speed := i
		move := remaining * speed
		if move > distance {
			count++
		}
	}
	return count
}

func main() {

	buf, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}

	var t int
	var d int

	lines := strings.Split(string(buf), "\n")
	fmt.Println(lines)
	for _, line := range lines {
		if line == "" {
			continue
		}
		s := strings.Split(line, ":")
		if s[0] == "Time" {
			nString := strings.Join(strings.Fields(s[1]), "")
			t, _ = strconv.Atoi(nString)
		}
		if s[0] == "Distance" {
			nString := strings.Join(strings.Fields(s[1]), "")
			d, _ = strconv.Atoi(nString)
		}
	}
	fmt.Println("----")
	fmt.Printf("time: %v \n", t)
	fmt.Printf("distance: %v \n", d)

	c := countPossibilities(t, d)
	fmt.Printf("%v\t%v\t%v\n", t, d, c)
}
