package funcs

import (
	"testing"
)

func TestGjsonGet(t *testing.T) {

	body := `{"name": "Bob"}`

	v := GjsonGet(body, "name")

	switch v.(type) {
	case string:
		// okay
	default:
		t.Fatalf("Invalid type, expected string")
	}

	if v.(string) != "Bob" {
		t.Fatalf("Expected Bob")
	}

}
