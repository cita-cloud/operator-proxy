package main

import (
	"fmt"
	"os"

	"github.com/cita-cloud/operator-proxy/cli/command"
)

func main() {
	err := command.RootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
