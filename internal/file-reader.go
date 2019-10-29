package internal

import (
    "bufio"
    "os"
)

type FileReader interface {
    Read(path string)
}

type FileReaderImpl struct {
    tokenizer Tokenizer
}

func (r FileReaderImpl) Read(path string) {
    file, err := os.Open(path)
    if err != nil {
        panic(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        r.tokenizer.Tokenize(scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        panic(err)
    }
}

// Factory function
func NewFileReader(tokenizer Tokenizer) FileReader {
    return FileReaderImpl{ tokenizer: tokenizer }
}
