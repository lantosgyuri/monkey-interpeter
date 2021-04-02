package main

import (
	"fmt"
	"github.com/lantosgyuri/monkey-interpreter/repl"
	"os"
	"os/user"
)

func main() {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hi %s! Type some code for me!\n", usr.Name)
	repl.Start(os.Stdin, os.Stdout)
}
