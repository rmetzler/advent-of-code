package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type round struct {
	Red   int
	Blue  int
	Green int
}

type game struct {
	Number int
	Rounds []round
}

func getNumber(r *regexp.Regexp, s string) int {
	submatches := r.FindStringSubmatch(s)
	if len(submatches) <= 1 {
		return 0
	}

	i, err := strconv.Atoi(submatches[1])

	if err != nil {
		fmt.Printf("string: %v\n", s)
		panic(err)
	}

	return i
}

func toRound(s string) round {
	// get the string, make red blue green
	red := getNumber(redReg, s)
	blue := getNumber(blueReg, s)
	green := getNumber(greenReg, s)
	return round{red, blue, green}
}

func toRounds(s string) []round {
	// get the string, make red blue green

	sArr := strings.Split(s, ";")
	rounds := make([]round, len(sArr))
	for i, sRound := range sArr {
		rounds[i] = toRound(sRound)
	}
	return rounds
}

func toGame(line string) game {
	s := strings.Split(line, ":")
	number, _ := strconv.Atoi(gameReg.FindStringSubmatch(s[0])[1])

	return game{Number: number, Rounds: toRounds(s[1])}
}

func (r round) valid() bool {
	valid := r.Blue <= threshold.Blue &&
		r.Red <= threshold.Red &&
		r.Green <= threshold.Green
	return valid
}

func (g game) valid() bool {
	for _, r := range g.Rounds {
		if !r.valid() {
			return false
		}
	}
	return true
}

var (
	threshold = round{Red: 12, Green: 13, Blue: 14}
	gameReg   = regexp.MustCompile("Game ([0-9]+)")
	redReg    = regexp.MustCompile("([0-9]+) red")
	blueReg   = regexp.MustCompile("([0-9]+) blue")
	greenReg  = regexp.MustCompile("([0-9]+) green")
)

func main() {

	buf, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}

	sum := 0
	lines := strings.Split(string(buf), "\n")
	fmt.Println(lines)
	for _, line := range lines {
		if len(line) < 2 {
			continue
		}
		g := toGame(line)
		fmt.Println(g)
		fmt.Printf("valid: %v\n", g.valid())
		if g.valid() {
			sum += g.Number
		}
	}
	fmt.Println("----")
	fmt.Printf("sum: %v\n", sum)
}
