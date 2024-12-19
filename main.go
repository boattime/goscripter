package main

import (
    "flag"
    "log"
    "path/filepath"

    "github.com/boattime/goscripter/internal/executor"
    "github.com/boattime/goscripter/internal/loader"
)

func main() {
    scriptPath := flag.String("script", "", "Path to the script to execute")
    flag.Parse()

    if *scriptPath == "" {
        log.Fatal("Please provide a script path using -script flag")
    }

    absPath, err := filepath.Abs(*scriptPath)
    if err != nil {
        log.Fatalf("Error resolving script path: %v", err)
    }

    // Load and execute the script
    script, err := loader.LoadScript(absPath)
    if err != nil {
        log.Fatalf("Error loading script: %v", err)
    }

    if err := executor.Execute(script); err != nil {
        log.Fatalf("Error executing script: %v", err)
    }
}
