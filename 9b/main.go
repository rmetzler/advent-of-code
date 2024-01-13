package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

func AllZero(numbers []int) bool {
	return lo.Every([]int{0}, numbers)
}

func Deduction(numbers []int) []int {
	var res []int = make([]int, len(numbers)-1)

	last := numbers[0]
	for i := 1; i < len(numbers); i++ {
		res[i-1] = numbers[i] - last
		last = numbers[i]
	}

	return res
}

func Extrapolate(numbers []int) int {
	if AllZero(numbers) {
		return 0
	}

	return numbers[(len(numbers)-1)] + Extrapolate(Deduction(numbers))
}

func ExtrapolateLeft(numbers []int) int {
	if AllZero(numbers) {
		return 0
	}

	return numbers[0] - ExtrapolateLeft(Deduction(numbers))
}

func main() {

	buf, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(buf), "\n")

	sum := 0
	for _, line := range lines {
		if len(line) == 0 {
			fmt.Println("empty line")
		}
		if len(line) < 1 {
			continue
		}

		// numbers := lo.Map(strings.Split(line, " "),
		// 	func(s string, index int) int {
		// 		i, _ := strconv.Atoi(s)
		// 		return i
		// 	})

		nstrings := strings.Split(line, " ")
		numbers := make([]int, len(nstrings))
		for i, nstring := range nstrings {
			num, _ := strconv.Atoi(nstring)
			numbers[i] = num
		}

		x := ExtrapolateLeft(numbers)
		sum += x
		fmt.Println(numbers, x)
	}
	fmt.Println("----")
	fmt.Println("sum: ", sum)
}
