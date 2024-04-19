package utils

import (
	"fmt"
	"log"
	"os"
	"runtime/debug"
	"time"
)

// IsFileExist checks if a file exists at the specified path.
// It returns true if the file exists, false otherwise.
func IsFileExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

// ReadFileData reads the contents of a file specified by the FileURI parameter.
// It returns the file data as a byte slice and an error if the file does not exist or if there was an error reading the file.
func ReadFileData(FileURI string) ([]byte, error) {
	if _, err := os.Stat(FileURI); os.IsNotExist(err) {
		return nil, fmt.Errorf("%s does not exist", FileURI)
	}

	fileData, err := os.ReadFile(FileURI)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	return fileData, nil
}

// InArray checks if a string is present in an array of strings.
func InArray(in string, array []string) bool {
	for _, element := range array {
		if in == element {
			return true
		}
	}
	return false
}

// GoWithRecover wraps a `go func()` with recover()
func GoWithRecover(handler func(), recoverHandler func(r interface{})) {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("%s goroutine panic: %v\n%s\n", time.Now().Format("2006-01-02 15:04:05"), r, string(debug.Stack()))
				if recoverHandler != nil {
					go func() {
						defer func() {
							if p := recover(); p != nil {
								log.Printf("recover goroutine panic:%v\n%s\n", p, string(debug.Stack()))
							}
						}()
						recoverHandler(r)
					}()
				}
			}
		}()
		handler()
	}()
}
