package search

import (
	"context"
	"testing"
)

func TestAll_success(t *testing.T) {
	ch := All(context.Background(), "file", []string{"../../data/test.txt"})
	_, ok := <-ch
	if !ok {
		t.Errorf("error: %v", ok)
	}
}