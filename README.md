# GoScripter

GoScripter is a flexible scripting system for Go that allows you to write, manage, and execute Go scripts while leveraging shared libraries and utilities. It provides a structured way to create and run scripts with access to common functionality.

## Features

- Write scripts in pure Go
- Access shared utility libraries
- Easy script execution
- Modular and extensible design
- Built-in common utilities
- Script management tools

## Installation

To install GoScripter, use the following command:

```bash
git clone https://github.com/boattime/goscripter
```

## Quick Start

Create a script in the scripts directory:
```go
// scripts/hello.go
package main

import (
    "fmt"
    "github.com/boattime/goscripter/libs/fileutils"
)

func main() {
    content, err := fileutils.ReadFile("example.txt")
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }
    fmt.Printf("File content: %s\n", content)
}

    Run your script:

go run main.go -script scripts/hello.go
```

## Project Structure
```
goscripter/
├── internal/
│   ├── executor/        # Script execution logic
│   └── loader/          # Script loading and validation
├── libs/                # Shared libraries
│   ├── fileutils/       # File operations
│   └── stringutils/     # String manipulation
├── scripts/             # Your scripts go here
└── main.go/

Available Libraries
FileUtils

fileutils.ReadFile(path string) (string, error)
fileutils.WriteFile(path string, content string) error
fileutils.EnsureDir(path string) error

StringUtils

// Coming soon
```

Writing Scripts

Scripts are regular Go programs that can import and use GoScripter's utility libraries. Create your scripts in the scripts/ directory:

```bash
package main

import (
    "fmt"
    "github.com/boattime/goscripter/libs/fileutils"
)

func main() {
    // Your script logic here
}
```

Usage Examples

    Basic script execution:

go run main.go -script scripts/myscript.go

    Using with arguments (coming soon):

```bash
go run main.go -script scripts/process.go -args="file1.txt,file2.txt"
```

Contributing

Contributions are welcome! Here's how you can help:

    Fork the repository
    Create your feature branch (git checkout -b feature/amazing-feature)
    Commit your changes (git commit -m 'Add some amazing feature')
    Push to the branch (git push origin feature/amazing-feature)
    Open a Pull Request

Development

To set up the development environment:

# Clone the repository
git clone https://github.com/boattime/goscripter.git

# Navigate to the project
cd goscripter

# Install dependencies
go mod tidy

# Build the project
go build main.go

# Run tests
go test ./...

Roadmap

    Script argument support
    Script templates
    Concurrent script execution
    Script dependency management
    Additional utility libraries
    Script validation and linting
    Configuration management

## License

This project is licensed under the MIT License - see the LICENSE file for details.
Support

If you encounter any issues or have questions, please file an issue on the GitHub repository.
