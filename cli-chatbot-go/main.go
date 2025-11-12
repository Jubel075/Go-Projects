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

func printMessage(message string, isBot bool) {
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

func showLoadingIndicator(done chan bool, message string) {
	frames := []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
	i := 0

	// Print header for the assistant while loading
	fmt.Printf("\n%s%s┌─ %sAssistant %s• %s%s\n", dim, colorGray, colorBlue, colorGray, message, colorReset)
	fmt.Printf("%s%s│%s  ", dim, colorGray, colorReset)

	for {
		select {
		case <-done:
			// Clear the loading line
			fmt.Print("\r\033[K")
			return
		default:
			frame := frames[i%len(frames)]
			fmt.Printf("\r%s%s %s%s %s", colorPurple, frame, colorReset, message, colorReset)
			time.Sleep(80 * time.Millisecond)
			i++
		}
	}
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
		printMessage(input, false)

		// Get response - try AI first, fallback to predefined responses
		var resp string

		// Check if it's a built-in command first
		if responses.IsCommand(input) {
			resp = responses.GetResponse(input)

			// Simple loading for commands
			done := make(chan bool)
			go showLoadingIndicator(done, "processing")
			time.Sleep(300 * time.Millisecond)
			done <- true

			// Clear the spinner header line and print the response
			fmt.Print("\033[1A\r\033[K")
			printMessage(resp, true)

			continue
		}

		// If AI enabled, stream; otherwise fallback
		if ai.IsEnabled() {
			// Show "sending prompt..." indicator briefly
			sendDone := make(chan bool)
			go showLoadingIndicator(sendDone, "sending prompt")
			time.Sleep(350 * time.Millisecond)
			sendDone <- true

			// Start "thinking" spinner
			thinkDone := make(chan bool)
			go showLoadingIndicator(thinkDone, "thinking")

			// Collect response
			var fullResponse strings.Builder
			timestamp := time.Now().Format("15:04")
			firstChunk := true

			err := ai.StreamResponse(input,
				func(chunk string) {
					if firstChunk {
						// Stop the thinking spinner
						thinkDone <- true

						// Clear spinner header
						fmt.Print("\033[1A\r\033[K")

						// Print actual assistant header (no leading \n)
						fmt.Printf("%s%s┌─ %sAssistant %s• %s%s%s\n", dim, colorGray, colorBlue, colorGray, timestamp, colorReset, colorReset)

						// Print initial content prefix and first chunk
						fmt.Printf("%s%s│%s  %s", dim, colorGray, colorReset, chunk)
						firstChunk = false
					} else {
						// Print subsequent chunks raw (streaming effect)
						fmt.Print(chunk)
					}
					fullResponse.WriteString(chunk)
				},
				func(finalText string) {
					// Close the box (no re-printing to avoid duplicates)
					fmt.Printf("\n%s%s└─%s\n\n", dim, colorGray, colorReset)
				},
			)

			if err != nil {
				// Stop spinner if not already
				if firstChunk {
					thinkDone <- true
				}
				log.Printf("AI error: %v, falling back to predefined responses\n", err)
				resp = responses.GetResponse(input) + "\n(AI error occurred; using fallback)"
				printMessage(resp, true)
			}

		} else {
			// Fallback to predefined responses
			done := make(chan bool)
			go showLoadingIndicator(done, "processing")
			time.Sleep(300 * time.Millisecond)
			done <- true

			resp = responses.GetResponse(input)
			fmt.Print("\033[1A\r\033[K")
			printMessage(resp, true)
		}
	}
}
