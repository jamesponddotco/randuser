package main

import (
	"os"

	"git.sr.ht/~jamesponddotco/randuser/internal/app"
)

func main() {
	os.Exit(app.Run(os.Args))
}
