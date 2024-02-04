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

func TestFormatUnixTime(t *testing.T) {

	tests := map[int64]string{
		1707090186: "February 04, 2024",
	}

	for ts, expected := range tests {

		ts_fmt := FormatUnixTime(ts, "January 02, 2006")

		if ts_fmt != expected {
			t.Fatalf("Unexpected result for '%d'. Expected '%s' but got '%s'.", ts, expected, ts_fmt)
		}
	}
}
