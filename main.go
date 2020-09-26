package main

import (
	"fmt"
	"os"
	"os/user"

	"waiacig/repl"
)

func main() {
	u, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the waiacig repl\n", u.Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.StartREPL(os.Stdin, os.Stdout)
}
