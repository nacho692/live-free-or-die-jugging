package app

import (
	"bufio"
	"fmt"
	"io"
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
	input io.Reader
}

func (r reader) Read() (string, error) {
	input, err := bufio.NewReader(r.input).ReadString('\n')
	if err != nil {
		return "", err
	}
	return input[:len(input)-1], nil
}
