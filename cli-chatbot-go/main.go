package main

import (
	"bufio"
	"cli_chatbot_go/ai"
	"cli_chatbot_go/responses"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

const (
	colorReset  = "\033[0m"
	colorBlue   = "\033[38;5;75m"
	colorGray   = "\033[38;5;243m"
	colorOrange = "\033[38;5;215m"
	colorPurple = "\033[38;5;141m"
	colorWhite  = "\033[38;5;255m"
	colorGreen  = "\033[38;5;150m"
	bold        = "\033[1m"
	dim         = "\033[2m"
)

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func printWelcome() {
	clearScreen()

	// ASCII art style banner
	fmt.Printf("\n")
	fmt.Printf("  %s%s━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━%s\n", bold, colorPurple, colorReset)
	fmt.Printf("\n")
	fmt.Printf("     %s%s⚡ CLI CHATBOT%s\n", bold, colorOrange, colorReset)
	fmt.Printf("     %s%sInspired by Claude Code & LazyVim%s\n", dim, colorBlue, colorReset)

	// Show AI status
	if ai.IsEnabled() {
		fmt.Printf("     %s%s🤖 AI Mode: Enabled (Gemini)%s\n", dim, colorGreen, colorReset)
	} else {
		fmt.Printf("     %s%s💡 AI Mode: Disabled (Fallback responses)%s\n", dim, colorGray, colorReset)
	}

	fmt.Printf("\n")
	fmt.Printf("  %s%s━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━%s\n", bold, colorPurple, colorReset)
	fmt.Printf("\n")

	// Commands section
	fmt.Printf("  %s%s📋 Available Commands:%s\n", dim, colorGray, colorReset)
	fmt.Printf("     help  info  time  joke  quote  clear  exit\n")
	fmt.Printf("\n")
	fmt.Printf("  %s%s💬 Start typing your message...%s\n", dim, colorGray, colorReset)
	fmt.Printf("\n")
}

func printMessage(sender, message string, isBot bool) {
	timestamp := time.Now().Format("15:04")

	if isBot {
		fmt.Printf("\n%s%s┌─ %sAssistant %s• %s%s%s\n", dim, colorGray, colorBlue, colorGray, timestamp, colorReset, colorReset)
		fmt.Printf("%s%s│%s  %s\n", dim, colorGray, colorReset, message)
		fmt.Printf("%s%s└─%s\n\n", dim, colorGray, colorReset)
	} else {
		fmt.Printf("%s%s┌─ %sYou %s• %s%s%s\n", dim, colorGray, colorGreen, colorGray, timestamp, colorReset, colorReset)
		fmt.Printf("%s%s│%s  %s\n", dim, colorGray, colorReset, message)
		fmt.Printf("%s%s└─%s\n", dim, colorGray, colorReset)
	}
}

func printPrompt() {
	fmt.Printf("%s%s❯%s ", bold, colorPurple, colorReset)
}

func showTypingIndicator() {
	frames := []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}

	fmt.Printf("\n%s%s┌─ %sAssistant %s• thinking%s\n", dim, colorGray, colorBlue, colorGray, colorReset)
	fmt.Printf("%s%s│%s  ", dim, colorGray, colorReset)

	// Animate for a short duration
	iterations := 8
	for i := 0; i < iterations; i++ {
		frame := frames[i%len(frames)]
		fmt.Printf("\r%s%s│%s  %s%s%s Thinking...", dim, colorGray, colorReset, colorPurple, frame, colorReset)
		time.Sleep(100 * time.Millisecond)
	}

	// Clear the loading lines
	fmt.Print("\r\033[K")        // Clear current line
	fmt.Print("\033[1A\r\033[K") // Move up and clear
	fmt.Print("\033[1A\r\033[K") // Move up and clear again
}

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, using default settings")
	}

	// Initialize AI
	if err := ai.Initialize(); err != nil {
		log.Printf("AI initialization: %v\n", err)
	}
	defer ai.Close()

	printWelcome()
	reader := bufio.NewReader(os.Stdin)

	for {
		printPrompt()
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("\n%s%sError reading input%s\n", colorReset, colorOrange, colorReset)
			continue
		}

		input = strings.TrimSpace(input)

		if input == "" {
			continue
		}

		// Handle exit command
		if input == "exit" || input == "/exit" {
			fmt.Printf("  %s%s━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━%s\n", bold, colorPurple, colorReset)
			fmt.Printf("%s%s%s  %sThanks for chatting! See you next time. %s✨%s           %s%s%s\n", bold, colorPurple, colorReset, colorWhite, colorOrange, colorReset, bold, colorPurple, colorReset)
			fmt.Printf("  %s%s━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━%s\n", bold, colorPurple, colorReset)
			os.Exit(0)
		}

		// Handle clear command
		if input == "clear" || input == "/clear" {
			printWelcome()
			continue
		}

		// Show user message
		printMessage("You", input, false)

		// Simulate typing
		showTypingIndicator()

		// Get response - try AI first, fallback to predefined responses
		var resp string

		// Check if it's a built-in command first
		if responses.IsCommand(input) {
			resp = responses.GetResponse(input)
		} else if ai.IsEnabled() {
			// Try to get AI response
			aiResp, err := ai.GetResponse(input)
			if err != nil {
				log.Printf("AI error: %v, falling back to predefined responses\n", err)
				resp = responses.GetResponse(input)
			} else {
				resp = aiResp
			}
		} else {
			// Fallback to predefined responses
			resp = responses.GetResponse(input)
		}

		// Display response
		printMessage("Assistant", resp, true)
	}
}
