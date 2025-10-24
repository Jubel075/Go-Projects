package main

import (
	"bufio"
	"cli_chatbot_go/responses"
	"fmt"
	"os"
	"strings"
)

func printWelcome() {
	fmt.Println("============================")
	fmt.Println(" Claude/LazyVim-Style CLI Chatbot ")
	fmt.Println("============================")
	fmt.Println("Type '/exit' to quit.")
}

func main() {
	printWelcome()
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("You > ")
		input, _ := reader.ReadString(' ')
		input = strings.TrimSpace(input)

		if input == "/exit" {
			fmt.Println("Bye!")
			os.Exit(0)
		}

		resp := responses.GetResponse(input)
		fmt.Printf("Bot > %s", resp)
	}
}
