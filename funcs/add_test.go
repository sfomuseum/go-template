package funcs

import (
	"testing"
)

func TestAdd(t *testing.T) {

	v := Add(1, 3)

	if v != 4 {
		t.Fatalf("Failed to add 1 + 3. Expected 4 but got %d", v)
	}
}
