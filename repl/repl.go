package repl

import (
	"bufio"
	"fmt"
	evaluator "github.com/lantosgyuri/monkey-interpreter/evaulator"
	"github.com/lantosgyuri/monkey-interpreter/lexer"
	"github.com/lantosgyuri/monkey-interpreter/object"
	"github.com/lantosgyuri/monkey-interpreter/parser"
	"io"
)

const PROMPT = "LETS TYPE >> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	for {
		_, err := fmt.Fprint(out,PROMPT)
		if err != nil {
			return
		}

		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()

		if len(p.Errors()) != 0 {
			printError(out, p.Errors())
		}

		eval := evaluator.Eval(program, env)

		if eval != nil {
			_, err := io.WriteString(out, eval.Inspect())
			if err != nil {
				fmt.Printf("Error with writing to the console: %v \n", err)
			}
		}
	}
}

func printError(out io.Writer, errors []string) {
	for _, v := range errors {
		_, err := io.WriteString(out, fmt.Sprintf("%v \n", v))
		if err != nil {
		fmt.Printf("Error with writing to the console: %v \n", err)
	}
	}
}