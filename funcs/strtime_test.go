package funcs

import (
	"testing"
)

func TestFormatStringTime(t *testing.T) {

	tests := map[string]string{
		"2022-03-14": "March 14, 2022",
	}

	for strtime, expected := range tests {

		strtime_fmt := FormatStringTime(strtime, "2006-01-02", "January 02, 2006")

		if strtime_fmt != expected {
			t.Fatalf("Unexpected result for '%s'. Expected '%s' but got '%s'.", strtime, expected, strtime_fmt)
		}
	}
}
