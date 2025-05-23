package repl

import (
	"bufio"
	"fmt"
	"goland-study/src/lexer"
	"goland-study/src/token"
	"io"
)

// Read-Eval-Print-Loop (读取值打印循环)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for {
		_, _ = fmt.Fprintf(out, PROMPT)
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
