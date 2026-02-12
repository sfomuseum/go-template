package funcs

import (
	"testing"
)

func TestURLQueryEscape(t *testing.T) {

	tests := map[string]string{
		"cl:///search/advanced/?q=cup&exhibition=1159160315&airline=1159284127": "cl%3A%2F%2F%2Fsearch%2Fadvanced%2F%3Fq%3Dcup%26exhibition%3D1159160315%26airline%3D1159284127",
	}

	for u, expected := range tests {

		u_esc := URLQueryEscape(u)

		if u_esc != expected {
			t.Fatalf("Unexpected result for '%s'. Expected '%s' but got '%s'.", u, expected, u_esc)
		}
	}
}
