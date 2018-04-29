package main

import (
	"fmt"
	"os"
)

var (
	version   string
	buildtime string
)

func main() {
	cliApp := NewCLI(version + " " + buildtime)

	err := cliApp.Run(os.Args)

	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

}
