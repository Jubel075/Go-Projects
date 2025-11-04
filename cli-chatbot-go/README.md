# 🤖 CLI Chatbot

A modern, beautiful command-line chatbot built with Go, inspired by Claude Code and LazyVim aesthetics.

![Go Version](https://img.shields.io/badge/Go-1.24-00ADD8?style=flat&logo=go)
![License](https://img.shields.io/badge/license-MIT-blue.svg)

## ✨ Features

- 🎨 **Beautiful Interface** - Modern UI with colors, borders, and smooth animations
- ⚡ **Fast & Lightweight** - Built with Go for optimal performance
- 💬 **Natural Conversations** - Smart pattern matching for contextual responses
- 🎯 **Multiple Commands** - Help, info, time, jokes, quotes, and more
- 🔄 **Real-time Feedback** - Typing indicators and timestamps
- 🎭 **Extensible** - Easy to add new commands and responses

## 📦 Installation

### Prerequisites
- Go 1.24 or higher

### Setup
```bash
# Clone the repository
git clone <your-repo-url>
cd cli_chatbot_go

# Initialize Go modules (if needed)
go mod tidy

# Run the chatbot
go run main.go
```

### Build Binary
```bash
# Build executable
go build -o chatbot

# Run the binary
./chatbot
```

## 🚀 Usage

### Starting the Chatbot
```bash
go run main.go
```

### Available Commands

| Command | Description |
|---------|-------------|
| `help` | Display all available commands |
| `info` | Learn about the chatbot |
| `time` | Show current date and time |
| `joke` | Get a random programming joke |
| `quote` | Get an inspiring developer quote |
| `clear` | Clear the terminal screen |
| `exit` | Exit the chatbot |

### Natural Language

The chatbot understands natural language patterns:

- **Greetings**: "hello", "hi", "hey", "what's up"
- **Questions**: "how are you", "who are you", "what can you do"
- **Gratitude**: "thanks", "thank you", "appreciate it"
- **Farewells**: "bye", "goodbye", "see you later"
- **Time queries**: "what time is it", "current time"
- **Entertainment**: "tell me a joke", "inspire me"

## 📁 Project Structure

```
cli_chatbot_go/
├── main.go              # Main application entry point
├── responses/
│   └── responses.go     # Response logic and pattern matching
├── go.mod               # Go module definition
├── config.yaml          # Configuration file
└── README.md            # This file
```

## 🎨 Customization

### Adding New Responses

Edit `responses/responses.go` to add new commands:

```go
var replyMap = map[string]string{
    "your_command": "Your response here",
}
```

### Adding Pattern Matching

Add contextual responses in the `GetResponse` function:

```go
if containsAny(input, []string{"keyword1", "keyword2"}) {
    return "Your contextual response"
}
```

### Customizing Colors

Modify color constants in `main.go`:

```go
const (
    colorBlue   = "\033[38;5;75m"   // Change color codes
    colorGreen  = "\033[38;5;150m"
    colorPurple = "\033[38;5;141m"
)
```

## 🔮 Future Enhancements

- [ ] AI integration (OpenAI, Anthropic Claude API)
- [ ] Conversation history
- [ ] YAML-based configuration loading
- [ ] Plugin system for extensions
- [ ] Multi-language support
- [ ] Session persistence
- [ ] Slash commands (e.g., `/help`, `/clear`)
- [ ] User preferences

## 🤝 Contributing

Contributions are welcome! Feel free to:

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📝 License

This project is open source and available under the MIT License.

## 🙏 Acknowledgments

- Inspired by [Claude Code](https://www.anthropic.com/) aesthetics
- UI patterns from [LazyVim](https://www.lazyvim.org/)
- Built with ❤️ and Go

## 📧 Contact

For questions or feedback, feel free to open an issue or reach out!

---

**Made with Go • Inspired by modern CLI tools • Built for developers**
