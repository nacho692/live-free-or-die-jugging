package app

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type writer struct {
	output io.Writer
}

func (w writer) Write(message string) error {
	_, err := fmt.Fprint(w.output, message)
	return err
}

func (w writer) WriteLn(message string) error {
	_, err := fmt.Fprintln(w.output, message)
	return err
}

type reader struct {
	input *bufio.Reader
}

func (r reader) Read() (string, error) {
	input, err := r.input.ReadString('\n')
	if err != nil {
		return "", err
	}
	// unix delimiter
	input = strings.TrimRight(input, "\n")
	// windows delimiter
	input = strings.TrimRight(input, "\r\n")
	return input, nil
}
