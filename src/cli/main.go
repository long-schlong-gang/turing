package main

import (
	"os"

	"github.com/teris-io/cli"

	"turing/commands"
)

func main() {

	app := cli.New("Cypher CLI").
		WithCommand(commands.Decypher)

	os.Exit(app.Run(os.Args, os.Stdout))
}

// TODO: Replace all constant vars with functions so they don't get diddled with by shitty devs
