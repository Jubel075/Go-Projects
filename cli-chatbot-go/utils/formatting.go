package utils

import (
	"regexp"
	"strings"
)

const (
	colorReset  = "\033[0m"
	colorBlue   = "\033[38;5;75m"
	colorGreen  = "\033[38;5;150m"
	colorOrange = "\033[38;5;215m"
	colorPurple = "\033[38;5;141m"
	colorYellow = "\033[38;5;228m"
	colorGray   = "\033[38;5;243m"
	bold        = "\033[1m"
	dim         = "\033[2m"
	italic      = "\033[3m"
)

// CleanMarkdown removes or converts markdown syntax to terminal-friendly format
func CleanMarkdown(text string) string {
	// Remove leading/trailing whitespace
	text = strings.TrimSpace(text)

	// Convert **bold** to colored bold
	boldRegex := regexp.MustCompile(`\*\*([^*]+)\*\*`)
	text = boldRegex.ReplaceAllString(text, bold+colorPurple+"$1"+colorReset)

	// Convert *italic* to dim text
	italicRegex := regexp.MustCompile(`\*([^*]+)\*`)
	text = italicRegex.ReplaceAllString(text, italic+"$1"+colorReset)

	// Convert # Headers to colored bold text
	headerRegex := regexp.MustCompile(`(?m)^#{1,6}\s+(.+)$`)
	text = headerRegex.ReplaceAllString(text, "\n"+bold+colorOrange+"$1"+colorReset+"\n")

	// Convert `code` to highlighted code
	codeRegex := regexp.MustCompile("`([^`]+)`")
	text = codeRegex.ReplaceAllString(text, colorGreen+"$1"+colorReset)

	// Convert bullet points to clean format
	// * item or - item → • item
	bulletRegex := regexp.MustCompile(`(?m)^[\s]*[\*\-]\s+(.+)$`)
	text = bulletRegex.ReplaceAllString(text, "  "+colorBlue+"•"+colorReset+" $1")

	// Convert numbered lists to colored numbers
	numberedRegex := regexp.MustCompile(`(?m)^[\s]*(\d+)\.\s+(.+)$`)
	text = numberedRegex.ReplaceAllString(text, "  "+colorPurple+"$1."+colorReset+" $2")

	// Remove excessive blank lines (more than 2 consecutive)
	excessiveNewlines := regexp.MustCompile(`\n{3,}`)
	text = excessiveNewlines.ReplaceAllString(text, "\n\n")

	// Clean up code blocks
	codeBlockRegex := regexp.MustCompile("(?s)```[a-z]*\n(.+?)\n```")
	text = codeBlockRegex.ReplaceAllStringFunc(text, func(match string) string {
		// Extract code content
		content := codeBlockRegex.FindStringSubmatch(match)
		if len(content) > 1 {
			lines := strings.Split(content[1], "\n")
			var formatted []string
			formatted = append(formatted, "\n"+dim+colorGray+"┌─ Code"+colorReset)
			for _, line := range lines {
				formatted = append(formatted, dim+colorGray+"│"+colorReset+"  "+colorGreen+line+colorReset)
			}
			formatted = append(formatted, dim+colorGray+"└─"+colorReset+"\n")
			return strings.Join(formatted, "\n")
		}
		return match
	})

	return text
}

// WrapText wraps text to specified width while preserving formatting
func WrapText(text string, width int) string {
	lines := strings.Split(text, "\n")
	var wrapped []string

	for _, line := range lines {
		// Skip wrapping for lines with special formatting
		if strings.HasPrefix(strings.TrimSpace(line), "•") ||
			strings.HasPrefix(strings.TrimSpace(line), "┌") ||
			strings.HasPrefix(strings.TrimSpace(line), "│") ||
			strings.HasPrefix(strings.TrimSpace(line), "└") {
			wrapped = append(wrapped, line)
			continue
		}

		// Simple word wrapping
		if len(line) <= width {
			wrapped = append(wrapped, line)
			continue
		}

		words := strings.Fields(line)
		var currentLine string
		for _, word := range words {
			if len(currentLine)+len(word)+1 <= width {
				if currentLine != "" {
					currentLine += " "
				}
				currentLine += word
			} else {
				if currentLine != "" {
					wrapped = append(wrapped, currentLine)
				}
				currentLine = word
			}
		}
		if currentLine != "" {
			wrapped = append(wrapped, currentLine)
		}
	}

	return strings.Join(wrapped, "\n")
}
