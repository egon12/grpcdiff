package grpcdiff

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

// how to use
// echo {} | grpcdiff localhost:50051 realservice:50051 proto.DigitalDicovery GetDynamicPage

func Main() {
	hostA := os.Args[1]
	hostB := os.Args[2]
	service := os.Args[3]
	method := os.Args[4]
	input := os.Stdin

	e, err := NewExecutor(hostA, hostB, service, method, input)
	if err != nil {
		log.Println(err)
	}
	e.ExecuteAll()
}

type Executor struct {
	c  *Caller
	in *bufio.Reader
}

func NewExecutor(hostA, hostB, service, method string, input io.Reader) (*Executor, error) {
	c, err := NewCaller(hostA, hostB, service, method)
	if err != nil {
		return nil, err
	}

	return &Executor{
		c:  c,
		in: bufio.NewReader(input),
	}, nil
}

func (e *Executor) ExecuteAll() {
	for {
		err := e.Execute()
		if err != nil {
			break
		}
	}
}

func (e *Executor) Execute() error {
	input, err := e.in.ReadString('\n')
	if err != nil {
		return err
	}

	input = strings.TrimRight(input, "\n")

	r, err := e.c.Call(input)
	if err != nil {
		fmt.Printf("input %s has error: %v\n", input, err)
		return err
	}

	fmt.Println(r.Report())
	return nil
}
