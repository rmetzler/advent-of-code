package main

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var (
	numberReg = regexp.MustCompile("([0-9]+),([0-9]+),([0-9]+)~([0-9]+),([0-9]+),([0-9]+)")
)

type Point [3]int

type Block struct {
	start Point
	end   Point
}

func Atoi(s string) int {
	s = strings.Trim(s, " ")
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func MaxPoint(points ...Point) Point {
	maxPoint := Point{}
	for _, p := range points {
		maxPoint[0] = max(maxPoint[0], p[0])
		maxPoint[1] = max(maxPoint[1], p[1])
		maxPoint[2] = max(maxPoint[2], p[2])
	}
	return maxPoint
}

type XYZSpace [][][]*Block

func EmptySpace(maxPoint Point) XYZSpace {
	xyzSpace := make(XYZSpace, maxPoint[0])
	for x := 1; x <= maxPoint[0]; x++ {
		yzPlane := make([][]*Block, maxPoint[1])
		xyzSpace[x-1] = yzPlane

		for y := 1; y <= maxPoint[0]; y++ {
			zColumn := make([]*Block, maxPoint[2])
			yzPlane[y-1] = zColumn
		}
	}
	return xyzSpace
}

type ByZ []Block

func (a ByZ) Len() int           { return len(a) }
func (a ByZ) Less(i, j int) bool { return a[i].start[2] < a[j].start[2] }
func (a ByZ) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func main() {

	// buf, err := os.ReadFile("sample")
	buf, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(buf), "\n")
	// fmt.Println(lines)
	blocks := []Block{}
	maxPoint := Point{}

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		// fmt.Println(line)
		if numberReg.MatchString(line) {
			match := numberReg.FindStringSubmatch(line)
			startPoint := Point{Atoi(match[1]), Atoi(match[2]), Atoi(match[3])}
			endPoint := Point{Atoi(match[4]), Atoi(match[5]), Atoi(match[6])}
			block := Block{startPoint, endPoint}
			if startPoint[2] > endPoint[2] {
				fmt.Println("startpoint z bigger:", block)
			}
			blocks = append(blocks, block)
			maxPoint = MaxPoint(maxPoint, startPoint, endPoint)
		}
	}
	fmt.Println("----")
	// fmt.Println(blocks)
	fmt.Println("MaxPoint:", maxPoint)

	// create empty space

	fmt.Println("----")
	fmt.Println("first unsorted:", blocks[0])
	fmt.Println("last unsorted:", blocks[len(blocks)-1])
	sort.Sort(ByZ(blocks))
	fmt.Println("first sorted:", blocks[0])
	fmt.Println("last sorted:", blocks[len(blocks)-1])

	fmt.Println("----")

	// _xyzSpace := EmptySpace(maxPoint)

	// for _, block := range blocks {
	// 	xyzSpace
	// }
}
