package main

import (
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type card struct {
	Number  int
	Winners []int
	Having  []int
	Copies  int
}

func toNumber(s string) int {
	i, err := strconv.Atoi(s)

	if err != nil {
		fmt.Printf("string: %v\n", s)
		panic(err)
	}
	return i
}

func toNumbers(s string) []int {
	fmt.Println("toNumbers: ", s)
	sArr := strings.Split(s, " ")
	numArr := []int{}
	for _, sNum := range sArr {
		if "" == sNum {
			continue
		}
		numArr = append(numArr, toNumber(sNum))
	}
	return numArr
}

func toCard(line string) card {
	s := strings.Split(line, ":")
	number, _ := strconv.Atoi(cardReg.FindStringSubmatch(s[0])[1])
	numStrings := strings.Split(s[1], "|")
	return card{
		Number:  number,
		Copies:  1,
		Winners: toNumbers(strings.Trim(numStrings[0], " ")),
		Having:  toNumbers(strings.Trim(numStrings[1], " ")),
	}
}

func (c card) countWinners() int {
	count := 0
	for _, number := range c.Winners {
		if slices.Contains(c.Having, number) {
			count += 1
		}
	}
	return count
}

var (
	cardReg = regexp.MustCompile("Card\\s+([0-9]+)")
)

func main() {

	buf, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}

	sum := 0
	lines := strings.Split(string(buf), "\n")
	// fmt.Println(lines)
	var cards []card
	for _, line := range lines {
		fmt.Println(line)
		if len(line) < 2 {
			continue
		}
		cards = append(cards, toCard(line))
	}

	for i, c := range cards {
		fmt.Println(c)
		count := c.countWinners()
		for j := 1; j <= count; j++ {
			cards[i+j].Copies += c.Copies
		}
		fmt.Printf("count: %v\tcopies:%v\n", count, c.Copies)
		sum += c.Copies
	}
	fmt.Println("----")
	fmt.Printf("sum: %v\n", sum)
}
