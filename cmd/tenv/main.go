package main

import (
	"log"
	"os"
	"path"

	"github.com/kbence/tenv/internal/commands"
	"github.com/kbence/tenv/internal/tenv"
)

func main() {
	binaryName := path.Base(os.Args[0])

	switch binaryName {
	case "teleport", "tsh", "tctl", "tbot":
		exitCode, err := tenv.Execute(binaryName, os.Args[1:]...)
		if err != nil {
			log.Printf("error executing '%s': %s", binaryName, err)
			os.Exit(exitCode + 128)
		}
	default:
		err := commands.NewRootCommand().Execute()
		if err != nil {
			log.Fatalf("error: %s", err)
		}
	}
}
