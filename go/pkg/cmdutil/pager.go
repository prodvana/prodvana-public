package cmdutil

import (
	"io"
	"os"
	"os/exec"

	"github.com/pkg/errors"
	"golang.org/x/term"
)

func MaybePage(input io.Reader) error {
	stdout := os.Stdout
	pager := os.Getenv("PAGER")
	if term.IsTerminal(int(stdout.Fd())) && pager != "" {
		// interactive, and pager is defined. try to page
		cmd := exec.Command(pager)
		cmd.Stdin = input
		cmd.Stdout = stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		return errors.Wrap(err, "failed to page")
	} else {
		// not attached to terminal, just pipe to stdout directly
		_, err := io.Copy(stdout, input)
		return errors.Wrap(err, "failed to write")
	}
}
