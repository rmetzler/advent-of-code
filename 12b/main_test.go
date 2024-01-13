package main

import (
	"regexp"
	"testing"
)

func TestNote_RegExpString(t *testing.T) {
	note := NewNote("?###???????? 3,2,1")
	rxString := note.RegExpString()
	rx := regexp.MustCompile(rxString)
	// fmt.Println(rxString)
	tests := []struct {
		field string
		want  bool
	}{
		{".###.##.#...", true},
		{".###.##..#..", true},
		{".###.##...#.", true},
		{".###.##....#", true},
		{".###..##.#..", true},
		{".###..##..#.", true},
		{".###..##...#", true},
		{".###...##.#.", true},
		{".###...##..#", true},
		{".###....##.#", true},
		{".##.....##.#", false},
	}
	for _, tt := range tests {
		t.Run("all match", func(t *testing.T) {
			if got := rx.MatchString(tt.field); got != tt.want {
				t.Errorf("Note.RegExpString() = %v, want %v (%v)", got, tt.want, tt.field)
			}
		})
	}
}
