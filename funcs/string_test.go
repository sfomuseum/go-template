package funcs

import (
	"testing"
)

func TestStringHasPrefix(t *testing.T) {

	ok := map[string]string{
		"hello world":       "hello",
		"hello world again": "he",
	}

	not_ok := map[string]string{
		"hello world": "bob",
	}

	for input, prefix := range ok {

		if !StringHasPrefix(input, prefix) {
			t.Fatalf("Expected '%s' to have prefix '%s'", input, prefix)
		}
	}

	for input, prefix := range not_ok {

		if StringHasPrefix(input, prefix) {
			t.Fatalf("Did not expected '%s' to have prefix '%s'", input, prefix)
		}
	}

}
