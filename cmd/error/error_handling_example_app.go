package main

import (
	"fmt"
	"io"
	"os"
)

type Runner interface {
	Run()
}

type MarathonRunner struct {
	writer io.Writer
	err    error
}

func (m *MarathonRunner) Run() {
	_, err := m.writer.Write([]byte("Running 42km in one take.\n"))
	m.err = err
}

func doBusyThing(runner Runner) {
	//just run for now
	runner.Run()
}

func main() {
	runner := &MarathonRunner{writer: os.Stdout}
	doBusyThing(runner)
	if runner.err != nil {
		fmt.Printf("There was an error while running: %v\n", runner.err)
	} else {
		fmt.Println("Running complete successfully.")
	}
}
