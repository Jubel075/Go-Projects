package main

import (
	"bufio"
	"cli_chatbot_go/responses"
	"fmt"
	"os"
	"strings"
	"time"
)

const (
	colorReset  = "\033[0m"
	colorBlue   = "\033[38;5;75m"
	colorGray   = "\033[38;5;243m"
	colorGreen  = "\033[38;5;150m"
	colorOrange = "\033[38;5;215m"
	colorPurple = "\033[38;5;141m"
	colorWhite  = "\033[38;5;255m"
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
	fmt.Printf("\n%s%s┌─ %sAssistant %s• thinking%s\n", dim, colorGray, colorBlue, colorGray, colorReset)
	fmt.Printf("%s%s│%s  %s●%s %s●%s %s●%s", dim, colorGray, colorReset, colorGray, colorReset, colorGray, colorReset, colorGray, colorReset)
	time.Sleep(300 * time.Millisecond)
	fmt.Print("\r\033[K")
	fmt.Print("\033[1A\r\033[K")
	fmt.Print("\033[1A\r\033[K")
}

func main() {
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
			fmt.Printf("\n%s%s╭─────────────────────────────────────────────────────────╮%s\n", bold, colorPurple, colorReset)
			fmt.Printf("%s%s│%s  %sThanks for chatting! See you next time. %s✨%s           %s%s│%s\n", bold, colorPurple, colorReset, colorWhite, colorOrange, colorReset, bold, colorPurple, colorReset)
			fmt.Printf("%s%s╰─────────────────────────────────────────────────────────╯%s\n\n", bold, colorPurple, colorReset)
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

		// Get and display response
		resp := responses.GetResponse(input)
		printMessage("Assistant", resp, true)
	}
}
