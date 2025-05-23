package main

import (
	"fmt"
	"goland-study/src/repl"
	"os"
	"os/user"
)

func main() {
	curUser, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello and welcome, %s!\n", curUser.Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
