# HTML Anchor Tag Parser

This Go package provides a simple utility for parsing HTML anchor (`<a>`) tags and extracting their `href` attributes and text content into a structured format.

## Features

- Parse HTML documents or fragments.
- Extract `href` attributes and text content from `<a>` tags.
- Return the data in a clean structure.

## Installation

Install the package using `go get`:

```bash
go get github.com/sarathkumar17/linkparser
```

Then import it into your project:

```go
import "github.com/sarathkumar17/linkparser"
```

## Usage

Here is an example of how to use the package:

```go
package main

import (
	"fmt"
	"log"

	"github.com/yourusername/html-anchor-parser"
)

func main() {
	htmlContent := `
		<html>
			<body>
				<a href="https://example.com">Example</a>
				<a href="/about">About Us</a>
			</body>
		</html>
	`

	anchors, err := linkparser.Parse(htmlContent)
	if err != nil {
		log.Fatalf("Failed to parse HTML: %v", err)
	}

	for _, anchor := range anchors {
		fmt.Printf("Href: %s, Text: %s\n", anchor.Href, anchor.Text)
	}
}
```

## API

### `Parse(html string) ([]Anchor, error)`

Parses the provided HTML and returns a slice of `Anchor` structs or an error if parsing fails.

### Anchor Struct

```go
type Anchor struct {
	Href string // The href attribute of the <a> tag
	Text string // The text content of the <a> tag
}
```

## Error Handling

The `Parse` function returns an error if the HTML input is invalid or if parsing fails. Handle errors appropriately in your application.

---

Contributions, issues, and feature requests are welcome! Feel free to open a PR or an issue in the repository.
