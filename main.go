package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"togo-list/cmd"
)

func main() {
	fmt.Println("toGo List!")

	add := flag.String("add", "", "Add a new task!")
	flag.Parse()

	var command cmd.Command
	if *add != "" {
		fmt.Println("add is a chosen option!", *add)
		command = &cmd.AddCommand{Task: *add}
	}

	if err := command.Execute(); err != nil {
		log.Println("Error during command execute:", err)
		os.Exit(1)
	}
}
