package text

import (
	"bufio"
	"bytes"
	"context"
	"embed"
	"strings"
	"testing"
)

//go:embed text_test.txt
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
		PageTitle string
	}

	vars := TestVars{
		PageTitle: "This is a test",
	}

	var buf bytes.Buffer
	wr := bufio.NewWriter(&buf)

	err = test_t.Execute(wr, vars)

	if err != nil {
		t.Fatalf("Failed to render template, %v", err)
	}

	wr.Flush()

	output := strings.TrimSpace(buf.String())

	expected := `This is a test
This is a test.`

	if output != expected {
		t.Fatalf("Unexpected output '%s'", output)
	}
}
