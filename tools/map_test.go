package tools_test

import (
	"testing"

	"github.com/TYTTYTTYT/claudia/tools"
)

func double(v int) int {
	return v * 2
}

func TestMap(t *testing.T) {
	data := make([]int, 1000000)
	for i := 0; i < 1000000; i++ {
		data[i] = i
	}
	tools.Map(data, double)

	for i := 0; i < 1000000; i++ {
		if data[i] != 2*i {
			t.Fatalf("Error while testing Map, want %v at position %v but get %v", 2*i, i, data[i])
		}
	}
}
