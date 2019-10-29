package internal

import (
    "fmt"
    "os"
    "path/filepath"
    "strings"
)

type PathCollector interface {
    Collect(rootPath string)
}

type PathCollectorImpl struct {
    pathQueue chan string
}

func (c PathCollectorImpl) Collect(rootPath string) {
    err := filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            fmt.Printf("Cannot read path %q: %v\n", path, err)
            return err
        }
        if !info.IsDir() && strings.HasSuffix(info.Name(), ".txt") {
            c.pathQueue <- path
        }
        return nil
    })
    close(c.pathQueue)
    if err != nil {
        panic(err)
    }
}

// Factory function
func NewPathCollector(pathQueue chan string) PathCollector {
    return PathCollectorImpl{ pathQueue: pathQueue }
}

