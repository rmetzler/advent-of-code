package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"golang.org/x/exp/maps"
)

var (
	stepsReg = regexp.MustCompile("^([LR]+)$")
	nodeReg  = regexp.MustCompile(`^(...) = \((...), (...)\)$`)
)

type Node struct {
	Name  string
	Left  string
	Right string
}

func NewNode(s string) *Node {
	match := nodeReg.FindStringSubmatch(s)
	if match == nil {
		return nil
	}
	fmt.Println(match)
	return &Node{
		Name:  match[1],
		Left:  match[2],
		Right: match[3],
	}
}

type Steps struct {
	Instructions string
	Step         int64
}

func (s *Steps) Next() string {
	x := s.Instructions[s.Step%int64(len(s.Instructions))]
	s.Step = s.Step + 1
	return string(x)
}

func StartNodes(nodes []string) []string {
	var result []string
	for _, node := range nodes {
		if string(node[2]) == "A" {
			result = append(result, node)
		}
	}
	return result
}

func AllEndNodes(nodes []string) bool {
	for _, node := range nodes {
		if string(node[2]) != "Z" {
			return false
		}
	}
	return true
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int64) int64 {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int64, integers ...int64) int64 {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func main() {

	buf, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}

	var steps Steps
	nodes := make(map[string]Node)
	lines := strings.Split(string(buf), "\n")
	fmt.Println(lines)
	for i, line := range lines {
		if i == 0 {
			steps = Steps{Instructions: stepsReg.FindString(line)}
			continue
		}
		if line == "" {
			fmt.Println("----")
			continue
		}

		node := NewNode(line)
		nodes[node.Name] = *node
	}
	fmt.Println("----")

	fmt.Printf("%#v\n", steps)
	fmt.Printf("%#v\n", nodes)
	fmt.Println("----")

	start := StartNodes(maps.Keys(nodes))
	end := make([]int64, len(start))
	for i, node := range start {
		current := node
		steps.Step = 0
		for string(current[2]) != "Z" {
			switch steps.Next() {
			case "L":
				current = nodes[current].Left
			case "R":
				current = nodes[current].Right
			default:
				panic("this should never happen")
			}
		}
		end[i] = steps.Step
		fmt.Printf("Start: %v\tSteps: %#v\n", start[i], end[i])
	}

	// lcm := LCM(end[0], end[1], end[2], end[3], end[4], end[5])
	lcm := LCM(end[0], end[1], end[2:]...)

	// current := StartNodes(maps.Keys(nodes))
	// for !AllEndNodes(current) {
	// 	fmt.Printf("Current : %v\n", current)
	// 	switch steps.Next() {
	// 	case "L":
	// 		for i, cNode := range current {
	// 			current[i] = nodes[cNode].Left
	// 		}
	// 	case "R":
	// 		for i, cNode := range current {
	// 			current[i] = nodes[cNode].Right
	// 		}
	// 	default:
	// 		panic("this should never happen")
	// 	}
	// }

	fmt.Printf("LCM: %v\n", lcm)
}
