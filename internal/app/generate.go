package app

import (
	"fmt"
	"os"

	"git.sr.ht/~jamesponddotco/randuser/internal/username"
	"git.sr.ht/~jamesponddotco/randuser/internal/validate"
	"github.com/urfave/cli/v2"
	"golang.org/x/term"
)

// GenerateAction is the main action for the application.
func GenerateAction(ctx *cli.Context) error {
	var (
		title      = ctx.Bool("title-case")
		format     = ctx.String("format")
		isTerminal = term.IsTerminal(int(os.Stdout.Fd()))
	)

	if !validate.Format(format) {
		return username.ErrInvalidFormat
	}

	var (
		name string
		err  error
	)

	if title {
		name, err = username.GenerateWithTitleCase(format)
		if err != nil {
			return fmt.Errorf("failed to generate username: %w", err)
		}
	} else {
		name, err = username.Generate(format)
		if err != nil {
			return fmt.Errorf("failed to generate username: %w", err)
		}
	}

	if isTerminal {
		if _, err = fmt.Fprintf(ctx.App.Writer, "%s\n", name); err != nil {
			return fmt.Errorf("failed to write output: %w", err)
		}

		return nil
	}

	if _, err = fmt.Fprintf(ctx.App.Writer, "%s", name); err != nil {
		return fmt.Errorf("failed to write output: %w", err)
	}

	return nil
}
