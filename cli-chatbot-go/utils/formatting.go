package utils

import (
	"fmt"
	"regexp"
	"strings"
)

const (
	colorReset  = "\033[0m"
	colorBlue   = "\033[38;5;75m"
	colorGray   = "\033[38;5;243m"
	colorPurple = "\033[38;5;141m"
	colorGreen  = "\033[38;5;150m"
	bold        = "\033[1m"
	dim         = "\033[2m"
)

// CleanMarkdown converts markdown into nicely formatted terminal output.
func CleanMarkdown(text string) string {
	text = strings.TrimSpace(text)

	// Headers -> colored, bold lines
	headerRegex := regexp.MustCompile(`(?m)^#{1,6}\s+(.+)`)
	text = headerRegex.ReplaceAllStringFunc(text, func(match string) string {
		parts := strings.SplitN(match, " ", 2)
		return fmt.Sprintf("\n%s%s%s%s\n", bold, colorPurple, parts[1], colorReset)
	})

	// Bold and italic
	text = regexp.MustCompile(`\*\*([^*]+)\*\*`).ReplaceAllString(text, fmt.Sprintf("%s$1%s", bold, colorReset))
	text = regexp.MustCompile(`\*([^*]+)\*`).ReplaceAllString(text, fmt.Sprintf("%s$1%s", dim, colorReset))

	// Inline code
	text = regexp.MustCompile("`([^`]+)`").ReplaceAllString(text, fmt.Sprintf("%s$1%s", colorGray, colorReset))

	// Code blocks
	codeBlockRegex := regexp.MustCompile("(?s)```[a-z]*\n(.+?)\n```")
	text = codeBlockRegex.ReplaceAllString(text, fmt.Sprintf("\n%s$1%s\n", colorGray, colorReset))

	// Bullet lists
	bulletRegex := regexp.MustCompile(`(?m)^[\s]*[\*\-]\s+(.+)`)
	text = bulletRegex.ReplaceAllString(text, fmt.Sprintf("  %s•%s $1", colorGreen, colorReset))

	// Numbered lists
	numberedRegex := regexp.MustCompile(`(?m)^[\s]*(\d+)\.\s+(.+)`)
	text = numberedRegex.ReplaceAllString(text, fmt.Sprintf("  %s$1.%s $2", colorBlue, colorReset))

	// Collapse excess blank lines
	text = regexp.MustCompile(`\n{3,}`).ReplaceAllString(text, "\n\n")

	return text
}

// WrapText wraps long lines for consistent terminal width.
func WrapText(text string, width int) string {
	lines := strings.Split(text, "\n")
	var wrapped []string

	for _, line := range lines {
		// Don’t wrap code or bullets
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, "•") || strings.HasPrefix(trimmed, "```") {
			wrapped = append(wrapped, line)
			continue
		}

		if len(line) <= width {
			wrapped = append(wrapped, line)
			continue
		}

		words := strings.Fields(line)
		var currentLine strings.Builder
		for _, word := range words {
			if currentLine.Len() == 0 {
				if len(word) <= width {
					currentLine.WriteString(word)
				} else {
					// Long word: hyphenate if possible
					for i := 0; i < len(word); i += width - 1 { // Leave space for hyphen
						end := i + width - 1
						if end > len(word) {
							end = len(word)
						}
						chunk := word[i:end]
						if end < len(word) {
							chunk += "-"
						}
						wrapped = append(wrapped, chunk)
					}
					continue
				}
			} else if currentLine.Len()+1+len(word) <= width {
				currentLine.WriteString(" " + word)
			} else {
				wrapped = append(wrapped, currentLine.String())
				currentLine.Reset()
				currentLine.WriteString(word)
			}
		}
		if currentLine.Len() > 0 {
			wrapped = append(wrapped, currentLine.String())
		}
	}

	return strings.Join(wrapped, "\n")
}
