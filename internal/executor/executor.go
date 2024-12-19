package executor

import (
    "fmt"
    "os/exec"
    "path/filepath"

    "github.com/boattime/goscripter/internal/loader"
)

func Execute(script *loader.Script) error {
    dir := filepath.Dir(script.Path)
    
    // Run the script using 'go run'
    cmd := exec.Command("go", "run", script.Path)
    cmd.Dir = dir
    
    output, err := cmd.CombinedOutput()
    if err != nil {
        return fmt.Errorf("error executing script: %v\nOutput: %s", err, output)
    }

    fmt.Printf("Script output:\n%s\n", output)
    return nil
}
