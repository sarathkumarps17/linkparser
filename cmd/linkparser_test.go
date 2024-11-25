// Extracts all anchor nodes from a simple HTML document
package linkparser

import (
	"strings"
	"testing"
)

// Parse valid HTML and return a list of links
func TestParseValidHTML(t *testing.T) {
	htmlData := `<html><body><a href="http://example.com">Example</a></body></html>`
	reader := strings.NewReader(htmlData)

	links, err := Parse(reader)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(links) != 1 {
		t.Fatalf("Expected 1 link, got %d", len(links))
	}

	expectedLink := Link{Href: "http://example.com", Text: "Example"}
	if links[0] != expectedLink {
		t.Errorf("Expected link %v, got %v", expectedLink, links[0])
	}
}

// Handle empty HTML documents gracefully
func TestParseEmptyHTML(t *testing.T) {
	htmlData := ``
	reader := strings.NewReader(htmlData)

	links, err := Parse(reader)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(links) != 0 {
		t.Fatalf("Expected 0 links, got %d", len(links))
	}
}

// Correctly identify and process all link nodes in the document
func TestParseMultipleLinks(t *testing.T) {
	htmlData := `<html><body>
						<a href="http://example.com">Example</a>
						<a href="http://test.com">Test</a>
					</body></html>`
	reader := strings.NewReader(htmlData)
	links, err := Parse(reader)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(links) != 2 {
		t.Fatalf("expected 2 links, got %d", len(links))
	}
	expectedHrefs := []string{"http://example.com", "http://test.com"}
	for i, link := range links {
		if link.Href != expectedHrefs[i] {
			t.Errorf("expected href '%s', got %s", expectedHrefs[i], link.Href)
		}
	}
}

func TestDocumentWithCommentedLinkTexts(t *testing.T) {
	htmlData := `<html>
    <body>
        <a href="/dog-cat">dog cat <!-- commented text SHOULD NOT be included! --></a>
    </body>
</html>`
	reader := strings.NewReader(htmlData)
	links, err := Parse(reader)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(links) != 1 {
		t.Fatalf("expected 1 links, got %d", len(links))
	}
	expectedHrefs := []string{"/dog-cat"}
	for i, link := range links {
		if link.Href != expectedHrefs[i] {
			t.Errorf("expected href '%s', got %s", expectedHrefs[i], link.Href)
		}
	}
}
