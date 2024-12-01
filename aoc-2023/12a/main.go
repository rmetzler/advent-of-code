package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Note struct {
	Condition string
	Damages   []int
}

func NewNote(s string) Note {
	parts := strings.Split(s, " ")
	dString := strings.Split(parts[1], ",")
	damages := make([]int, len(dString))
	for i, val := range dString {
		damages[i], _ = strconv.Atoi(val)
	}
	return Note{parts[0], damages}
}

func (n Note) RegExpString() string {
	parts := make([]string, len(n.Damages))
	for i, val := range n.Damages {
		parts[i] = fmt.Sprintf("#{%v}", val)
	}
	rx := `^\.*` + strings.Join(parts, `\.+`) + `\.*$`
	return rx
}

func (n Note) GetCombinationStrings() []string {
	rxString := n.RegExpString()
	rx := regexp.MustCompile(rxString)

	var result []string
	for i := int64(0); i < (int64(1) << len(n.Condition)); i++ { // i < 2^qm
		myString := strconv.FormatInt(i, 2)
		myString = fmt.Sprintf("%0*s", len(n.Condition), myString)
		myString = strings.ReplaceAll(myString, "0", ".")
		myString = strings.ReplaceAll(myString, "1", "#")
		if n.MatchesKnownCondition(myString) {

			// fmt.Println("original: ", n.Condition)
			// fmt.Println("candidate: ", myString, "needed groups", n.Damages)
			// fmt.Println("matches:", rx.MatchString(myString))
			if rx.MatchString(myString) {
				result = append(result, myString)
			}
		}
		// if n.MatchesKnownCondition(myString) && rx.MatchString(myString) {
		// 	fmt.Println(myString)
		// 	result = append(result, myString)
		// }
	}

	return result
}

func (n Note) MatchesKnownCondition(s string) bool {
	if len(n.Condition) != len(s) {
		return false
	}
	for i, val := range n.Condition {
		if (val != '?') && (val != rune(s[i])) {
			return false
		}
	}
	return true
}

func main() {

	buf, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(buf), "\n")
	fmt.Println(lines)

	total := 0

	for _, line := range lines {
		if len(line) == 0 {
			fmt.Println("empty line")
		}
		if len(line) < 1 {
			continue
		}
		n := NewNote(line)
		fmt.Printf("%#v\t%v\n", n, n.RegExpString())
		total += len(n.GetCombinationStrings())
	}
	fmt.Println("----")
	fmt.Println("total:", total)
}
