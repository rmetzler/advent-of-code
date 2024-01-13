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

type seedrange struct {
	start int
	count int
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

	var seeds []seedrange
	var gm gardenmap
	rules := make(map[string]gardenmap)

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		if seedsReg.MatchString(line) {
			fmt.Println("seeds:", line)

			split := strings.Split(line, " ")
			seeds = make([]seedrange, (len(split)-1)/2)
			for i := 1; i <= (len(split)-1)/2; i++ {
				seeds[i-1] = seedrange{
					start: Atoi(split[i*2-1]),
					count: Atoi(split[i*2]),
				}
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
	for _, sr := range seeds {
		fmt.Printf("seedrange: %#v\n", sr)
		for seed := sr.start; seed < sr.start+sr.count; seed++ {
			// fmt.Println("----")
			// fmt.Println("seed:", seed)
			soil := rules["seed-to-soil"].Convert(seed)
			// fmt.Println("soil:", soil)
			fertilizer := rules["soil-to-fertilizer"].Convert(soil)
			// fmt.Println("fertilizer:", fertilizer)
			water := rules["fertilizer-to-water"].Convert(fertilizer)
			// fmt.Println("water:", water)
			light := rules["water-to-light"].Convert(water)
			// fmt.Println("light:", light)
			temperature := rules["light-to-temperature"].Convert(light)
			// fmt.Println("temperature:", temperature)
			humidity := rules["temperature-to-humidity"].Convert(temperature)
			// fmt.Println("humidity:", humidity)
			location := rules["humidity-to-location"].Convert(humidity)
			// fmt.Println("location:", location)
			minLocation = min(minLocation, location)
		}
	}

	fmt.Println("----")
	fmt.Println("minLocation:", minLocation)
}
