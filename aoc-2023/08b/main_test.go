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

func TestAllEndNodes(t *testing.T) {
	tests := []struct {
		name  string
		nodes []string
		want  bool
	}{
		{
			name:  "alle enden auf Z",
			nodes: []string{"11Z", "22Z", "33Z"},
			want:  true,
		},
		{
			name:  "einer endet nicht auf Z",
			nodes: []string{"11Z", "22Z", "33Z", "00A"},
			want:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AllEndNodes(tt.nodes); got != tt.want {
				t.Errorf("AllEndNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}
