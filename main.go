package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
)

type Command struct {
    Parameter int
    Type      CommandType
}

type Program []Command

func IsWhitespaceChar(char byte) bool {
    return char == Space || char == Tab || char == Newline
}

func ExtractWhitespaceFromFile(filepath string) ([]byte, error) {

    const InitialBufferSize int = 256

    file, err := os.Open(filepath)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    reader := bufio.NewReader(file)
    whitespaceChars := make([]byte, 0, InitialBufferSize)

    for {
        currentByte, err := reader.ReadByte()

        if err != nil {
            if errors.Is(err, io.EOF) {
                break
            }

            return nil, fmt.Errorf("[Error] Failed to read file `%s` - %w", filepath, err)
        }

        if IsWhitespaceChar(currentByte) {
            whitespaceChars = append(whitespaceChars, currentByte)
        }
    }

    return whitespaceChars, nil
}

func ParseWhitespaceChars(whitespaceChars []byte) (Program, error) {
    return nil, nil
}

func InterpretProgram(program Program) (exitCode int) {
    return 0
}

func main() {

    if len(os.Args) < 2 {
        fmt.Fprintln(os.Stderr, "[Error] No Whitespace file provided. Aborting")
        os.Exit(1)
    }

    whitespaceChars, err := ExtractWhitespaceFromFile(os.Args[1])
    if err != nil {
        fmt.Fprintln(os.Stderr, err.Error())
        os.Exit(1)
    }

    program, err := ParseWhitespaceChars(whitespaceChars)
    if err != nil {
        fmt.Fprintln(os.Stderr, err.Error())
        os.Exit(1)
    }

    os.Exit(InterpretProgram(program))
}
