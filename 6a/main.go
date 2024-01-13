package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func toInt(sArr []string) []int {
	var intArr = make([]int, len(sArr))

	for idx, s := range sArr {
		i, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		intArr[idx] = i
	}

	return intArr
}

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

	var time []int
	var distance []int

	lines := strings.Split(string(buf), "\n")
	fmt.Println(lines)
	for _, line := range lines {
		if line == "" {
			continue
		}
		s := strings.Split(line, ":")
		if s[0] == "Time" {
			time = toInt(strings.Fields(s[1]))

		}
		if s[0] == "Distance" {
			distance = toInt(strings.Fields(s[1]))
		}
	}
	fmt.Println("----")
	fmt.Printf("time: %v \n", time)
	fmt.Printf("distance: %v \n", distance)

	total := 1
	for i := 0; i < len(time); i++ {
		t, d := time[i], distance[i]
		c := countPossibilities(t, d)
		fmt.Printf("%v\t%v\t%v\n", t, d, c)
		total *= c
	}
	fmt.Printf("Total: %v\n", total)
}
