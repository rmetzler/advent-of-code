package main

import (
	"reflect"
	"testing"
)

func TestNewNode(t *testing.T) {
	tests := []struct {
		line string
		want *Node
	}{
		{
			line: "AAA = (BBB, CCC)",
			want: &Node{"AAA", "BBB", "CCC"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.line, func(t *testing.T) {
			if got := NewNode(tt.line); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
