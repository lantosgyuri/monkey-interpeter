package repl

import (
	"bufio"
	"fmt"
	"github.com/lantosgyuri/monkey-interpreter/lexer"
	"github.com/lantosgyuri/monkey-interpreter/token"
	"io"
)

const PROMPT = "LETS TYPE >> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

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

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			_, _ = fmt.Fprintf(out, "%+v\n", tok)
		}
	}
}