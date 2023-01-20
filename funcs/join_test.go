package funcs

import (
	"testing"
)

func TestJoinPath(t *testing.T) {

	type TestJoin struct {
		Prefix   string
		Path     string
		Expected string
	}

	joins := []TestJoin{
		TestJoin{
			Prefix:   "",
			Path:     "/javascript/test.js",
			Expected: "/javascript/test.js",
		},
		TestJoin{
			Prefix:   "/testing",
			Path:     "/javascript/test.js",
			Expected: "/testing/javascript/test.js",
		},
		TestJoin{
			Prefix:   "/testing/",
			Path:     "/javascript/test.js",
			Expected: "/testing/javascript/test.js",
		},
		TestJoin{
			Prefix:   "/testing/../",
			Path:     "/javascript/test.js",
			Expected: "/javascript/test.js",
		},
	}

	for _, j := range joins {

		v := JoinPath(j.Prefix, j.Path)

		if v != j.Expected {
			t.Fatalf("Join failed. Expected '%s', but got '%s'", j.Expected, v)
		}
	}
}
