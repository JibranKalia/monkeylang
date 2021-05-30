package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/jibrankalia/monkeylang/repl"
)

func main() {
	user, err := user.Current()

	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is Monkey Lang\n", user.Username)
	fmt.Printf("Free free to type\n")
	repl.Start(os.Stdin, os.Stdout)
}
