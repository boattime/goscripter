package loader

import (
    "fmt"
    "go/parser"
    "go/token"
    "io/ioutil"
)

type Script struct {
    Path    string
    Content string
}

func LoadScript(path string) (*Script, error) {
    content, err := ioutil.ReadFile(path)
    if err != nil {
        return nil, fmt.Errorf("error reading script: %v", err)
    }

    // Validate that it's valid Go code
    fset := token.NewFileSet()
    if _, err := parser.ParseFile(fset, path, content, parser.AllErrors); err != nil {
        return nil, fmt.Errorf("invalid Go script: %v", err)
    }

    return &Script{
        Path:    path,
        Content: string(content),
    }, nil
}
