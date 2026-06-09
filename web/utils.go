package web

import (
	"regexp"
	"strings"
)
var strongRegex = regexp.MustCompile(`\*\*(.*?)\*\*`)
func strongify(markdown string)string {
	return strongRegex.ReplaceAllString(markdown, "<strong>$1</strong>")
}
// takes markdown and converts to html - but only list items and strong works
// markup now returns a slice of strings, each containing processed HTML for a single item
func markup(input string) []string {
	

	// Split by newline followed by a dash to capture multi-line blocks cleanly
	// We handle the very first dash by cleaning up the initial string
	input = strings.TrimSpace(input)
	input = strings.TrimPrefix(input, "- ")

	chunks := strings.Split(input, "\n- ")
	var items []string

	for _, chunk := range chunks {
		content := strings.TrimSpace(chunk)
		if content == "" {
			continue
		}

		// Replace any internal hard line breaks with a space
		content = regexp.MustCompile(`\s*\n\s*`).ReplaceAllString(content, " ")

		// Convert **bold** to <strong>bold</strong>
		content = strongify(content)

		items = append(items, content)
	}

	return items
}
