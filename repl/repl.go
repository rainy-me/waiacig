package repl

import (
	"bufio"
	"flag"
	"fmt"
	"io"

	"waiig/evaluator"
	"waiig/lexer"
	"waiig/parser"
)

const MONKEY_FACE = `
            __,__
   .--.  .-"     "-.  .--.
  / .. \/  .-. .-.  \/ .. \
 | |  '|  /   Y   \  |'  | |
 | \   \  \ 0 | 0 /  /   / |
  \ '- ,\.-"""""""-./, -' /
   ''-' /_   ^ ^   _\ '-''
       |  \._   _./  |
       \   \ '~' /   /
        '._ '-=-' _.'
           '-----'
`

const PROMPT = ">> "

func StartREPL(in io.Reader, out io.Writer) {
	flag.Parse()
	scanner := bufio.NewScanner(in)
	io.WriteString(out, MONKEY_FACE)
	for {
		fmt.Printf(PROMPT)
		if !scanner.Scan() {
			return
		}
		line := scanner.Text()
		l := lexer.NewLexer(line)
		p := parser.NewParser(l)
		program := p.ParseProgram()
		errors := p.Errors()
		if len(errors) != 0 {
			for _, msg := range errors {
				io.WriteString(out, "\t"+msg+"\n")
			}
			continue
		}

		if evaluated := evaluator.Eval(program); evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}
