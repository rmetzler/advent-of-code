package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

var (
	stepsReg = regexp.MustCompile("^([LR]+)$")
	nodeReg  = regexp.MustCompile(`([A-Z]{3}) = \(([A-Z]{3}), ([A-Z]{3})\)`)
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
	Step         int
}

func (s *Steps) Next() string {
	x := s.Instructions[s.Step%len(s.Instructions)]
	s.Step = s.Step + 1
	return string(x)
}

const (
	START_NODE = "AAA"
	END_NODE   = "ZZZ"
)

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

	current := START_NODE
	for current != END_NODE {
		fmt.Printf("%#v\n", nodes[current])
		switch steps.Next() {
		case "L":
			fmt.Println("L")
			current = nodes[current].Left
		case "R":
			fmt.Println("R")
			current = nodes[current].Right
		default:
			panic("this should never happen")
		}
	}
	fmt.Printf("Number of Steps: %v\n", steps.Step)
}
