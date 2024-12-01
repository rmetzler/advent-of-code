package main

import (
	"fmt"
	"os"
	"strings"
)

type Pipe struct {
	Symbol  rune
	Connect [2]Direction
}

type Direction int

const (
	NOTHING Direction = iota
	UNKNOWN
	NORTH
	EAST
	SOUTH
	WEST
)

type Position [2]int

var (
	VERTICAL_NS   Pipe = Pipe{'|', [2]Direction{NORTH, SOUTH}}
	HORIZONTAL_EW Pipe = Pipe{'-', [2]Direction{EAST, WEST}}
	BEND_NE       Pipe = Pipe{'L', [2]Direction{NORTH, EAST}}
	BEND_NW       Pipe = Pipe{'J', [2]Direction{NORTH, WEST}}
	BEND_SW       Pipe = Pipe{'7', [2]Direction{SOUTH, WEST}}
	BEND_SE       Pipe = Pipe{'F', [2]Direction{SOUTH, EAST}}
	GROUND        Pipe = Pipe{'.', [2]Direction{NOTHING, NOTHING}}
	START         Pipe = Pipe{'S', [2]Direction{UNKNOWN, UNKNOWN}}

	MOVE map[Direction]Position = map[Direction]Position{
		NORTH: {0, -1},
		SOUTH: {0, +1},
		EAST:  {+1, 0},
		WEST:  {-1, 0},
	}

	OPPOSITE map[Direction]Direction = map[Direction]Direction{
		NORTH: SOUTH,
		SOUTH: NORTH,
		EAST:  WEST,
		WEST:  EAST,
	}
)

var (
	pipeMap map[rune]Pipe = map[rune]Pipe{
		VERTICAL_NS.Symbol:   VERTICAL_NS,
		HORIZONTAL_EW.Symbol: HORIZONTAL_EW,
		BEND_NE.Symbol:       BEND_NE,
		BEND_NW.Symbol:       BEND_NW,
		BEND_SW.Symbol:       BEND_SW,
		BEND_SE.Symbol:       BEND_SE,
		GROUND.Symbol:        GROUND,
		START.Symbol:         START,
	}
)

type Field struct {
	Pipes [][]Pipe
	Start Position
}

// func (f Field) String() string {

// 	for x, line := range f {
// 		for y, r := range line {
// 			string(line)
// 		}
// 	}

// 	return
// }

func main() {

	buf, err := os.ReadFile("sample")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(buf), "\n")
	fmt.Println(lines)
	width := len(lines[0])
	field := [][]Pipe{}

	empty := make([]Pipe, width)
	for i := range empty {
		empty[i] = GROUND
	}

	field = append(field, empty)
	for _, line := range lines {
		if line == "" {
			fmt.Println("empty line")
		}
		if len(line) < 1 {
			continue
		}
	}
	fmt.Println("----")

	field = append(field, empty)

	fmt.Println(field)
}
