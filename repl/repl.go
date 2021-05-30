package repl

import (
	"fmt"
	"io"

	"github.com/jibrankalia/monkeylang/lexer"
	"github.com/jibrankalia/monkeylang/token"
)

// "fmt"
// "io"

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := buffer.NewScanner(in)

	for {
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Fprintf(out, "%+v\n", tok)
		}
	}
}
