package main

import (
	"fmt"
	"os"
	"regexp"
	"slices"
	"sort"
	"strconv"
	"strings"
)

var (
	numberReg = regexp.MustCompile("([0-9]+),([0-9]+),([0-9]+)~([0-9]+),([0-9]+),([0-9]+)")
)

type Point [3]int

const (
	X int = 0
	Y int = 1
	Z int = 2
)

type Points []Point

type Block struct {
	start Point
	end   Point
	// settled *Points
	// BlocksBelow *[]*Block
	// BlocksAbove *[]*Block
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
		maxPoint[X] = max(maxPoint[X], p[X])
		maxPoint[Y] = max(maxPoint[Y], p[Y])
		maxPoint[Z] = max(maxPoint[Z], p[Z])
	}
	return maxPoint
}

type XYZSpace [][][]*Block

func EmptySpace(maxPoint Point) XYZSpace {
	xyzSpace := make(XYZSpace, maxPoint[X]+1)
	for x := 0; x <= maxPoint[X]; x++ {
		yzPlane := make([][]*Block, maxPoint[Y]+1)
		xyzSpace[x] = yzPlane

		for y := 0; y <= maxPoint[Y]; y++ {
			zColumn := make([]*Block, maxPoint[Z]+1)
			yzPlane[y] = zColumn
		}
	}
	return xyzSpace
}

type XYCoordinates [2]int
type XYCoordList []XYCoordinates

func (b Block) Direction(d int) int {
	if b.start[d] == b.end[d] {
		return 0
	}
	if b.start[d] < b.end[d] {
		return 1
	}
	return -1
}

func Add(p1, p2 Point) Point {
	return Point{
		p1[X] + p2[X],
		p1[Y] + p2[Y],
		p1[Z] + p2[Z],
	}
}

func Equals(p1, p2 Point) bool {
	return (p1[X] == p2[X]) && (p1[Y] == p2[Y]) && (p1[Z] == p2[Z])
}

func Dimensions(b Block) int {
	return ((b.end[X]-b.start[X])*b.Direction(X) + 1) * ((b.end[Y]-b.start[Y])*b.Direction(Y) + 1) * ((b.end[Z]-b.start[Z])*b.Direction(Z) + 1)
}

func (b Block) Orientation() (byte, error) {
	direction := Point{
		b.Direction(X),
		b.Direction(Y),
		b.Direction(Z),
	}
	switch direction {
	case Point{1, 0, 0}, Point{-1, 0, 0}:
		return 'x', nil
	case Point{0, 1, 0}, Point{0, -1, 0}:
		return 'y', nil
	default:
		return 'z', nil
	}
}

func GetPointList(b Block) []Point {
	direction := Point{
		b.Direction(X),
		b.Direction(Y),
		b.Direction(Z),
	}

	re := []Point{}
	for p := b.start; !Equals(p, b.end); p = Add(p, direction) {
		re = append(re, p)
	}
	re = append(re, b.end)

	return re
}

func (xyz *XYZSpace) Insert(b Block) {
	// 	simple: add it in the exact place

	pointList := GetPointList(b)
	zColumns := make([][]*Block, len(pointList))
	maxNonNil := make([]int, len(pointList))

	for i, p := range pointList {
		zColumn := (*xyz)[p[X]][p[Y]]
		zColumns[i] = zColumn
		maxNonNil[i] = -1
		for z, bPointer := range zColumn {
			if bPointer != nil {
				maxNonNil[i] = z
			}
		}
	}
	newZ := slices.Max(maxNonNil) + 1
	o, err := b.Orientation()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Block: %v, Orientation %c, newZ: %v\n", b, o, newZ)

	switch o {
	case 'z':
		fmt.Println("case z")
		pointList := GetPointList(b)

		// we assume that the pointList is ordered from lowest Z to highest Z
		zDiff := pointList[0][Z] - newZ

		for _, p := range pointList {
			x, y, z := p[X], p[Y], p[Z]
			shiftedZ := z - zDiff
			fmt.Printf("shifted z: %v\n", shiftedZ)
			zColumn := (*xyz)[x][y]

			if zColumn[shiftedZ] != nil {
				fmt.Printf("Problem with Block %#v, Point %#v, blocked by Block %#v\n", b, p, (*xyz)[x][y][shiftedZ])
			} else {
				zColumn[shiftedZ] = &b
			}
		}

	default:
		for _, p := range GetPointList(b) {
			x, y := p[X], p[Y]
			zColumn := (*xyz)[x][y]
			if zColumn[newZ] != nil {
				fmt.Printf("Problem with Block %#v, Point %#v, blocked by Block %#v\n", b, p, (*xyz)[x][y][newZ])
			} else {
				zColumn[newZ] = &b
			}
		}
	}

}

type ByZ []Block

func (a ByZ) Len() int           { return len(a) }
func (a ByZ) Less(i, j int) bool { return a[i].start[Z] < a[j].start[Z] }
func (a ByZ) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

type BySize []Block

func (a BySize) Len() int           { return len(a) }
func (a BySize) Less(i, j int) bool { return Dimensions(a[i]) < Dimensions(a[j]) }
func (a BySize) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func main() {

	buf, err := os.ReadFile("sample")
	// buf, err := os.ReadFile("input")
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

	lastBlock := blocks[len(blocks)-1]
	fmt.Println(GetPointList(lastBlock))
	zeroBlock := Block{Point{0, 0, 0}, Point{0, 0, 0}}
	fmt.Println(GetPointList(zeroBlock))

	fmt.Println("----")

	fmt.Println(Dimensions(lastBlock))
	fmt.Println(Dimensions(zeroBlock))

	fmt.Println("----")

	// sort.Sort(BySize(blocks))
	// lastBlock = blocks[len(blocks)-1]

	// fmt.Println(lastBlock, Dimensions(lastBlock), GetPointList(lastBlock))

	xyzSpace := EmptySpace(maxPoint)

	for _, block := range blocks {
		xyzSpace.Insert(block)
	}

	fmt.Println("----")

	fmt.Println(xyzSpace)
}
