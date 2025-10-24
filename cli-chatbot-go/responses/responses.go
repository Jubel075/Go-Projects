package responses

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const (
	colorReset  = "\033[0m"
	colorBlue   = "\033[38;5;75m"
	colorGreen  = "\033[38;5;150m"
	colorOrange = "\033[38;5;215m"
	colorPurple = "\033[38;5;141m"
	bold        = "\033[1m"
)

var replyMap = map[string]string{
	"hello":     "Hi there! 👋 How can I assist you today?",
	"hi":        "Hello! Great to see you. What can I help you with?",
	"hey":       "Hey! What's on your mind?",
	"help":      fmt.Sprintf("Available commands:\n  %s%shelp%s - Show this help message\n  %s%sinfo%s - Learn about this chatbot\n  %s%stime%s - Get current time\n  %s%sjoke%s - Hear a programming joke\n  %s%squote%s - Get an inspiring quote\n  %s%sclear%s - Clear the screen\n  %s%sexit%s - Exit the chatbot", bold, colorGreen, colorReset, bold, colorGreen, colorReset, bold, colorGreen, colorReset, bold, colorGreen, colorReset, bold, colorGreen, colorReset, bold, colorGreen, colorReset, bold, colorGreen, colorReset),
	"info":      fmt.Sprintf("✨ %sCLI Chatbot v1.0%s\n\nA modern command-line chatbot inspired by Claude Code and LazyVim aesthetics.\nBuilt with Go for speed and simplicity.\n\nFeatures:\n  • Clean, colorful interface\n  • Real-time responses\n  • Extensible command system\n  • Future AI integration planned!", bold, colorReset),
	"time":      getCurrentTime(),
	"joke":      getRandomJoke(),
	"quote":     getRandomQuote(),
	"thanks":    "You're very welcome! 😊",
	"thank you": "Happy to help! Let me know if you need anything else.",
	"bye":       "Goodbye! Have a wonderful day! 👋",
	"goodbye":   "Take care! Come back anytime!",
}

var jokes = []string{
	"Why do programmers prefer dark mode? Because light attracts bugs! 🐛",
	"Why don't programmers like nature? It has too many bugs! 🌳",
	"How many programmers does it take to change a light bulb? None, that's a hardware problem! 💡",
	"What's a programmer's favorite hangout place? The Foo Bar! 🍺",
	"Why did the developer go broke? Because he used up all his cache! 💸",
	"What do you call a programmer from Finland? Nerdic! 🇫🇮",
	"A SQL query walks into a bar, walks up to two tables and asks, 'Can I join you?' 🗄️",
}

var quotes = []string{
	"\"Code is like humor. When you have to explain it, it's bad.\" - Cory House",
	"\"First, solve the problem. Then, write the code.\" - John Johnson",
	"\"Experience is the name everyone gives to their mistakes.\" - Oscar Wilde",
	"\"The best error message is the one that never shows up.\" - Thomas Fuchs",
	"\"Simplicity is the soul of efficiency.\" - Austin Freeman",
	"\"Make it work, make it right, make it fast.\" - Kent Beck",
	"\"Programming isn't about what you know; it's about what you can figure out.\" - Chris Pine",
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func getCurrentTime() string {
	now := time.Now()
	return fmt.Sprintf("🕐 Current time: %s%s%s", bold, now.Format("Monday, January 2, 2006 at 3:04 PM"), colorReset)
}

func getRandomJoke() string {
	return jokes[rand.Intn(len(jokes))]
}

func getRandomQuote() string {
	return quotes[rand.Intn(len(quotes))]
}

func GetResponse(input string) string {
	input = strings.ToLower(strings.TrimSpace(input))

	// Direct command matches
	if resp, ok := replyMap[input]; ok {
		// Refresh time if time command
		if input == "time" {
			return getCurrentTime()
		}
		// Get new joke/quote each time
		if input == "joke" {
			return getRandomJoke()
		}
		if input == "quote" {
			return getRandomQuote()
		}
		return resp
	}

	// Pattern matching for more natural conversation
	if containsAny(input, []string{"hello", "hi", "hey", "greetings"}) {
		return replyMap["hello"]
	}

	if containsAny(input, []string{"help", "commands", "what can you do"}) {
		return replyMap["help"]
	}

	if containsAny(input, []string{"thank", "thanks", "appreciate"}) {
		return replyMap["thanks"]
	}

	if containsAny(input, []string{"bye", "goodbye", "see you", "later"}) {
		return replyMap["bye"]
	}

	if containsAny(input, []string{"what time", "current time", "time is it"}) {
		return getCurrentTime()
	}

	if containsAny(input, []string{"joke", "funny", "make me laugh"}) {
		return getRandomJoke()
	}

	if containsAny(input, []string{"quote", "inspire", "motivation"}) {
		return getRandomQuote()
	}

	if containsAny(input, []string{"how are you", "how's it going", "what's up"}) {
		return "I'm doing great, thanks for asking! How can I help you today? 😊"
	}

	if containsAny(input, []string{"who are you", "what are you", "your name"}) {
		return fmt.Sprintf("I'm a CLI chatbot built with Go! Think of me as a helpful assistant in your terminal. Type %s%shelp%s to see what I can do!", bold, colorGreen, colorReset)
	}

	// Default response
	responses := []string{
		"I'm not sure how to respond to that yet. Try typing 'help' to see what I can do! 🤔",
		"Hmm, I don't have a response for that. Type 'help' to see available commands.",
		"That's interesting, but I'm still learning! Try 'help' to see what I know.",
	}
	return responses[rand.Intn(len(responses))]
}

func containsAny(input string, keywords []string) bool {
	for _, keyword := range keywords {
		if strings.Contains(input, keyword) {
			return true
		}
	}
	return false
}
