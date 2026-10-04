package repl

import (
	"bufio"
	"fmt"
	"interpreter-go/01/lexer"
	"interpreter-go/01/token"
	"io"
)

const PROMPT = ">>"

func Start(in io.Reader, out io.Writer) {
	Scanner := bufio.NewScanner(in)

	for {
		fmt.Print(PROMPT)
		Scanned := Scanner.Scan()
		if !Scanned {
			return
		}

		line := Scanner.Text()
		l := lexer.New(line)
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
