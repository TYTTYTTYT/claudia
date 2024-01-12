package tools_test

import (
	"testing"

	"github.com/TYTTYTTYT/claudia/tools"
)

func TestLRU(t *testing.T) {
	cache := tools.NewLRU[int, int](500)
	for i := 0; i < 1000; i++ {
		cache.Put(i, i)
	}
	for i := 0; i < 500; i++ {
		_, ok := cache.Get(i)
		if ok {
			t.Fatalf("Key %v should not exists in the cache but it did.", i)
		}
	}
	for i := 500; i < 1000; i++ {
		r, ok := cache.Get(i)
		if !ok {
			t.Fatalf("Key %v should exists in the cache but it did not.", i)
		}
		if r != i {
			t.Fatalf("Key %v should map to value %v but found %v instead.", i, i, r)
		}
	}
}
