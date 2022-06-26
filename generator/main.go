package main

import "log"

func main() {
	log.Printf("Hi, Welcome to FakeSchemaGenerator. 🎉 \n")

	cli := new(CLI)
	cli.Run()
}
