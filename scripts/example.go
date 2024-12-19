package main

import (
    "fmt"
    "github.com/boattime/goscripter/libs/fileutils"
)

func main() {
    content, err := fileutils.ReadFile("test.txt")
    if err != nil {
        fmt.Printf("Error reading file: %v\n", err)
        return
    }

    fmt.Printf("File content: %s\n", content)
}
