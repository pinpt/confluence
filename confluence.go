package confluence

import "github.com/pinpt/confluence/internal/html"

// ParseToHTML will parse confluence wiki format into HTML
func ParseToHTML(buf []byte) (string, error) {
	return html.Parse(string(buf))
}
