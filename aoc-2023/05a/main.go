package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	seedsReg   = regexp.MustCompile(`^seeds:(\s[0-9]+)*$`)
	numberReg  = regexp.MustCompile(`^([0-9]+)\s([0-9]+)\s([0-9]+)$`)
	mapNameReg = regexp.MustCompile(`^(.*)\smap:$`)
)

// seed-to-soil map: describes how to convert a seed number (the source) to a soil number (the destination).

// 50 98 2
// 52 50 48
// The first line has a destination range start of 50, a source range start of 98, and a range length of 2. This line means that the source range starts at 98 and contains two values: 98 and 99. The destination range is the same length, but it starts at 50, so its two values are 50 and 51. With this information, you know that seed number 98 corresponds to soil number 50 and that seed number 99 corresponds to soil number 51.

// type seeds []int

type gardenmap struct {
	Name  string
	Rules *[]convertrule
}
type convertrule struct {
	srcStart    int
	destStart   int
	rangeLength int
}

func (gm gardenmap) Convert(input int) int {
	for _, rule := range *gm.Rules {
		if (rule.srcStart <= input) &&
			(rule.srcStart+rule.rangeLength > input) {
			return (rule.destStart + (input - rule.srcStart))
		}
	}
	return input
}

// seeds
// seed-to-soil
// soil-to-fertilizer
// fertilizer-to-water
// water-to-light
// light-to-temperature
// temperature-to-humidity
// humidity-to-location

func Atoi(s string) int {
	s = strings.Trim(s, " ")
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func main() {

	buf, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(buf), "\n")
	// fmt.Println(lines)

	var seeds []int
	var gm gardenmap
	rules := make(map[string]gardenmap)

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		if seedsReg.MatchString(line) {
			fmt.Println("seeds:", line)

			split := strings.Split(line, " ")
			seeds = make([]int, len(split)-1)
			for i := 1; i < len(split); i++ {
				seeds[i-1] = Atoi(split[i])
			}
			fmt.Println("seeds:", seeds)
		}
		if mapNameReg.MatchString(line) {
			fmt.Println("map:", line)
			match := mapNameReg.FindStringSubmatch(line)
			name := match[1]
			gm = gardenmap{
				Name:  name,
				Rules: &[]convertrule{},
			}
			rules[name] = gm
		}
		if numberReg.MatchString(line) {
			fmt.Println("num:", line)
			match := numberReg.FindStringSubmatch(line)
			convertrule := convertrule{
				destStart:   Atoi(match[1]),
				srcStart:    Atoi(match[2]),
				rangeLength: Atoi(match[3]),
			}

			*gm.Rules = append(*gm.Rules, convertrule)
		}
	}
	fmt.Println("----")
	fmt.Println(seeds)
	fmt.Println(rules)
	fmt.Println(rules["fertilizer-to-water"].Rules)

	fmt.Println("----")

	minLocation := math.MaxInt
	for _, seed := range seeds {
		fmt.Println("----")
		fmt.Println("seed:", seed)
		soil := rules["seed-to-soil"].Convert(seed)
		fmt.Println("soil:", soil)
		fertilizer := rules["soil-to-fertilizer"].Convert(soil)
		fmt.Println("fertilizer:", fertilizer)
		water := rules["fertilizer-to-water"].Convert(fertilizer)
		fmt.Println("water:", water)
		light := rules["water-to-light"].Convert(water)
		fmt.Println("light:", light)
		temperature := rules["light-to-temperature"].Convert(light)
		fmt.Println("temperature:", temperature)
		humidity := rules["temperature-to-humidity"].Convert(temperature)
		fmt.Println("humidity:", humidity)
		location := rules["humidity-to-location"].Convert(humidity)
		fmt.Println("location:", location)
		minLocation = min(minLocation, location)
	}

	fmt.Println("----")
	fmt.Println("minLocation:", minLocation)
}
