package css

import (
	"bufio"
	"bytes"
	"context"
	"embed"
	"strings"
	"testing"
)

//go:embed css_test.css
var FS embed.FS

func TestLoadTemplates(t *testing.T) {

	ctx := context.Background()

	tpl, err := LoadTemplates(ctx, FS)

	if err != nil {
		t.Fatalf("Failed to load templates, %v", err)
	}

	test_t := tpl.Lookup("test")

	if test_t == nil {
		t.Fatalf("Unable to find 'test' template")
	}

	type TestVars struct {
		BackgroundColor string
	}

	vars := TestVars{
		BackgroundColor: "#ccc999",
	}

	var buf bytes.Buffer
	wr := bufio.NewWriter(&buf)

	err = test_t.Execute(wr, vars)

	if err != nil {
		t.Fatalf("Failed to render template, %v", err)
	}

	wr.Flush()

	output := strings.TrimSpace(buf.String())

	expected := `body { background-color: #ccc999; }`

	if output != expected {
		t.Fatalf("Unexpected output '%s'", output)
	}
}
