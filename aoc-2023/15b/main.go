package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

var (
	deleteReg = regexp.MustCompile(`^(\w+)-$`)
	assignReg = regexp.MustCompile(`^(\w+)=([0-9]+)$`)
)

type Lens struct {
	Name        string
	FocalLength int
}

func (l Lens) String() string { return fmt.Sprintf("[%v %v]", l.Name, l.FocalLength) }

type Box []*Lens

func (b Box) Empty() bool { return len(b) == 0 }

type Hashmap struct {
	Boxes []*Box
}

func NewHashmap() Hashmap {
	boxes := make([]*Box, 256)
	for i := range boxes {
		boxes[i] = &Box{}
	}
	return Hashmap{Boxes: boxes}
}

func Hash(s string) byte {
	hash := 0
	for _, b := range s {
		hash += int(b)
		hash = hash * 17
		hash = hash % 256
	}
	return byte(hash)
}

func (hm *Hashmap) Insert(name string, val int) {
	hash := Hash(name)
	box := hm.Boxes[hash]
	alreadyInBox := false
	// if box == nil {
	// 	box = &Box{}
	// }
	for _, lens := range *box {
		if lens.Name == name {
			lens.FocalLength = val
			alreadyInBox = true
		}
	}

	if !alreadyInBox {
		newBox := append(*box, &Lens{name, val})
		hm.Boxes[hash] = &newBox
	}
}

func (hm *Hashmap) Delete(name string) {
	hash := Hash(name)
	box := *hm.Boxes[hash]
	newBox := Box{}
	for _, lens := range box {
		if lens.Name != name {
			newBox = append(newBox, lens)
		}
	}
	hm.Boxes[hash] = &newBox
}

func (hm Hashmap) FocusingPower() int {
	sum := 0
	for i, box := range hm.Boxes {
		// if box == nil {
		// 	continue
		// }
		for j, lens := range *box {
			sum += (i + 1) * (j + 1) * lens.FocalLength
		}
	}
	return sum
}

func (hm Hashmap) String() string {
	result := ""
	for i, box := range hm.Boxes {
		// if box == nil {
		// 	continue
		// }
		if !box.Empty() {
			result += fmt.Sprintf("Box %v: %v\n", i,
				strings.Join(
					lo.Map(*box, func(l *Lens, index int) string { return l.String() }),
					" ",
				))
		}
	}
	return result
}

func main() {

	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		// Return nothing if at end of file and no data passed
		if atEOF && len(data) == 0 {
			return 0, nil, nil
		}

		if i := strings.IndexAny(string(data), ",\n"); i >= 0 {
			return i + 1, data[0:i], nil
		}

		// If at end of file with data return the data
		if atEOF {
			return len(data), data, nil
		}

		return
	})

	hashmap := NewHashmap()

	for scanner.Scan() {
		operation := scanner.Text()

		if deleteReg.MatchString(operation) {
			fmt.Println("DELETE\t", operation)
			match := deleteReg.FindStringSubmatch(operation)
			hashmap.Delete(match[1])
		} else {
			if assignReg.MatchString(operation) {
				fmt.Println("ASSIGN\t", operation)
				match := assignReg.FindStringSubmatch(operation)
				name := match[1]
				val, err := strconv.Atoi(match[2])
				if err != nil {
					panic(err)
				}
				hashmap.Insert(name, val)
			} else {
				fmt.Println("PROBLEM\t", operation)
			}
		}
	}
	fmt.Println("----")
	fmt.Println(hashmap)
	fmt.Println(hashmap.FocusingPower())
}
