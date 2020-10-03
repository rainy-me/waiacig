package repl

import (
	"bufio"
	"flag"
	"fmt"
	"io"

	"waiacig/compiler"
	"waiacig/evaluator"
	"waiacig/lexer"
	"waiacig/object"
	"waiacig/parser"
	"waiacig/vm"
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

var vmFlag = flag.Bool("vm", false, "enable vm")

func StartREPL(in io.Reader, out io.Writer) {
	flag.Parse()
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()
	macroEnv := object.NewEnvironment()
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

		if *vmFlag {
			comp := compiler.NewCompiler()
			err := comp.Compile(program)
			if err != nil {
				fmt.Fprintf(out, "Woops! Compilation failed:\n %s\n", err)
				continue
			}
			machine := vm.NewVM(comp.Bytecode())
			err = machine.Run()
			if err != nil {
				fmt.Fprintf(out, "Woops! Executing bytecode failed:\n %s\n", err)
				continue
			}
			lastPopped := machine.LastPoppedStackElem()
			io.WriteString(out, lastPopped.Inspect())
			io.WriteString(out, "\n")
		}

		evaluator.DefineMacros(program, macroEnv)
		expanded := evaluator.ExpandMacros(program, macroEnv)
		if evaluated := evaluator.Eval(expanded, env); evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}
