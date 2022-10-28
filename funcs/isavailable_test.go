package funcs

import (
	"testing"
)

func TestIsAvailable(t *testing.T) {

	type Vars struct {
		Hello string
	}

	vars := Vars{
		Hello: "world",
	}

	if !IsAvailable("Hello", vars) {
		t.Fatalf("Expected to find 'Hello' in vars")
	}

	if IsAvailable("Foo", vars) {
		t.Fatalf("Did not expect to find 'Foo' in vars")
	}
}
