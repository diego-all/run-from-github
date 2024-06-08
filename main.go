package main

import (
	"fmt"
	"os"

	"github.com/diego-all/run-from-github/cli"
	"github.com/diego-all/run-from-github/codexgen"
)

func main() {
	args := os.Args[1:]

	fmt.Print("ARGS from main: ", args)

	//cli.Start(os.Stdout, args, codexgen.Welcome)
	cli.Start(os.Stdout, args, codexgen.Generate)
}
