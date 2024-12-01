package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Hand []string
type Bid int

type Player struct {
	Hand
	Bid
}

type Round struct {
	Players []Player
}

type Rank int

const (
	RANK_HIGH_CARD Rank = iota
	RANK_ONE_PAIR
	RANK_TWO_PAIR
	RANK_THREE_OF_A_KIND
	RANK_FULL_HOUSE
	RANK_FOUR_OF_A_KIND
	RANK_FIVE_OF_A_KIND
)

var cardValueMap = map[string]int{
	"A": 14,
	"K": 13,
	"Q": 12,
	// "J": 11,
	"T": 10,
	"9": 9,
	"8": 8,
	"7": 7,
	"6": 6,
	"5": 5,
	"4": 4,
	"3": 3,
	"2": 2,
	"J": 1, // J cards are now the weakest individual cards
}

type ByCardValue Hand

func (a ByCardValue) Len() int           { return len(a) }
func (a ByCardValue) Less(i, j int) bool { return cardValueMap[a[i]] < cardValueMap[a[j]] }
func (a ByCardValue) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

type ByHandValue []Player

func (a ByHandValue) Len() int { return len(a) }
func (a ByHandValue) Less(i, j int) bool {
	rankA := RankOfHand(a[i].Hand)
	rankB := RankOfHand(a[j].Hand)

	// only if both have the same rank, compare the cards
	if rankA == rankB {
		for k := range a[i].Hand {
			cvA := cardValueMap[a[i].Hand[k]]
			cvB := cardValueMap[a[j].Hand[k]]
			if cvA != cvB {
				return cvA < cvB
			}
		}
	}
	return rankA < rankB
}

func (a ByHandValue) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func NewHand(s string) Hand {
	return strings.Split(s, "")
}

func NewPlayer(line string) Player {
	s := strings.Split(line, " ")

	hand := NewHand(s[0])
	bid, err := strconv.Atoi(s[1])
	if err != nil {
		panic(err)
	}

	return Player{Hand(hand), Bid(bid)}
}

func RankOfHand(hand Hand) Rank {
	// fmt.Println(hand)
	count := make(map[string]int)
	jokers := 0
	for _, card := range hand {
		if card == "J" {
			// jokers are counted individually
			jokers += 1
			continue
		}
		count[card] = count[card] + 1
	}
	var counts []int
	for _, v := range count {
		counts = append(counts, v)
	}
	// order from highes to lowest count
	sort.Sort(sort.Reverse(sort.IntSlice(counts)))

	// we need this rule to accommodate to the hand "JJJJJ"
	if jokers >= 5 {
		return RANK_FIVE_OF_A_KIND
	}

	// J cards can pretend to be whatever card is best
	//  for the purpose of determining hand type

	counts[0] += jokers

	switch counts[0] {
	case 5:
		return RANK_FIVE_OF_A_KIND
	case 4:
		return RANK_FOUR_OF_A_KIND
	case 3:
		{
			if counts[1] == 2 {
				return RANK_FULL_HOUSE
			} else {
				return RANK_THREE_OF_A_KIND
			}
		}
	case 2:
		{
			if counts[1] == 2 {
				return RANK_TWO_PAIR
			} else {
				return RANK_ONE_PAIR
			}
		}
	default:
		return RANK_HIGH_CARD
	}
}

func main() {

	buf, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(buf), "\n")
	fmt.Println(lines)
	round := Round{}
	for _, line := range lines {
		if len(line) == 0 {
			fmt.Println("empty line")
		}
		if len(line) < 1 {
			continue
		}
		round.Players = append(round.Players, NewPlayer(line))
	}

	fmt.Printf("%#v\n", round)
	fmt.Println("----")

	// hands are sorted from lowest to highest
	sort.Sort(ByHandValue(round.Players))
	fmt.Printf("%#v\n", round)
	fmt.Println("----")

	total := 0
	for i, p := range round.Players {
		total += (i + 1) * int(p.Bid)
	}
	fmt.Printf("total: %v\n", total)
}
