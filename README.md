# go-template

Opinionated Go package for loading templates with default functions.

## Documentation

[![Go Reference](https://pkg.go.dev/badge/github.com/sfomuseum/go-template.svg)](https://pkg.go.dev/github.com/sfomuseum/go-template)

## Example

```
package main

import (
	"bufio"
	"bytes"
	"context"
	"embed"
	"github.com/sfomuseum/go-template/html"
	"strings"
)

//go:embed html_test.html
var FS embed.FS

func main() {

	ctx := context.Background()

	tpl, _ := html.LoadTemplates(ctx, FS)

	test_t := tpl.Lookup("test")

	type TestVars struct {
		PageTitle string
	}

	vars := TestVars{
		PageTitle: "This is a test",
	}

	var buf bytes.Buffer
	wr := bufio.NewWriter(&buf)

	test_t.Execute(wr, vars)
}
```

Where `html_test.html` looks like this:

```
{{ define "test"}}
<html>
	<head>
		<title>{{ if (IsAvailable "PageTitle" .) }}{{ .PageTitle }}{{ end }}</title>
	</head>
	<body>{{ if (IsAvailable "UserName" .) }}Hello {{ .UserName }}{{ else }}This is a test.{{ end }}
	</body>
</html>
{{ end }}
```