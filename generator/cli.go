package main

import (
	"flag"
	"fmt"
	"log"

	"os"
)

const (
	USAGE = `Usage:
	fsg generate -n n
		Generate n json Files in json Folder
	fsg post -n n
		Post n json Files to server`
)

type CLI struct{}

func (c *CLI) printUsage() {
	fmt.Println(USAGE)
}

func (c *CLI) checkArgs() {
	if len(os.Args) < 2 {
		c.printUsage()
		os.Exit(1)
	}
}

func check(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func (c *CLI) Run() {
	c.checkArgs()

	generateCmd := flag.NewFlagSet("generate", flag.ExitOnError)
	postCmd := flag.NewFlagSet("post", flag.ExitOnError)
	generateNumber := generateCmd.Int("n", 0, "Type in Generate number:")
	postNumber := postCmd.Int("n", 0, "Type in Post Number:")

	switch os.Args[1] {
	case "generate":
		err := generateCmd.Parse(os.Args[2:])
		check(err)
		buildJson(*generateNumber)
	case "post":
		err := postCmd.Parse(os.Args[2:])
		check(err)
		postJson(*postNumber)
	default:
		c.printUsage()
		os.Exit(1)
	}
}
