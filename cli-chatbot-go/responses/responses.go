package responses

import "strings"

var replyMap = map[string]string{
	"hello": "Hi there! How can I assist you today?",
	"help": "Available commands: hello, help, info, exit.",
	"info": "This is a lightweight CLI chatbot inspired by Claude and LazyVim. Future updates will include AI integration!",
	"exit": "Exiting the chat...",
}

func GetResponse(input string) string {
	input = strings.ToLower(strings.TrimSpace(input))
	if resp, ok := replyMap[input]; ok {
		return resp
	}
	return "Sorry, I haven't been programmed to answer that yet."
}
