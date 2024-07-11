package repl

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/Stupnikjs/interpreter/lexer"
	"github.com/Stupnikjs/interpreter/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		if string(line) == "exit" {
			os.Exit(2)
		}
		l := lexer.New(line)
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
