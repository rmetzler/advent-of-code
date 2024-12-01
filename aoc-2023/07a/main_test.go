package main

import "testing"

func TestRankOfHand(t *testing.T) {
	tests := []struct {
		hand string
		want Rank
	}{
		{hand: "AAAAA", want: RANK_FIVE_OF_A_KIND},
		{hand: "AA8AA", want: RANK_FOUR_OF_A_KIND},
		{hand: "23332", want: RANK_FULL_HOUSE},
		{hand: "TTT98", want: RANK_THREE_OF_A_KIND},
		{hand: "23432", want: RANK_TWO_PAIR},
		{hand: "A23A4", want: RANK_ONE_PAIR},
		{hand: "23456", want: RANK_HIGH_CARD},
	}
	for _, tt := range tests {
		t.Run(tt.hand, func(t *testing.T) {
			if got := RankOfHand(NewHand(tt.hand)); got != tt.want {
				t.Errorf("RankOfHand() = %v, want %v", got, tt.want)
			}
		})
	}
}
