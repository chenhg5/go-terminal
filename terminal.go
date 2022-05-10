package terminal

import (
	"io"
	"os"

	"golang.org/x/term"
)

type shell struct {
	r io.Reader
	w io.Writer
}

func (sh *shell) Read(data []byte) (n int, err error) {
	return sh.r.Read(data)
}

func (sh *shell) Write(data []byte) (n int, err error) {
	return sh.w.Write(data)
}

type Terminal struct {
	oldState *term.State
	term     *term.Terminal
	fd       int
}

func NewTerminal(prompt string) (*Terminal, error) {
	sh := &shell{r: os.Stdin, w: os.Stdout}
	te := term.NewTerminal(sh, prompt)
	fd := int(os.Stdin.Fd())
	oldState, err := term.MakeRaw(fd)
	return &Terminal{
		oldState: oldState,
		term:     te,
		fd:       fd,
	}, err
}

func (t *Terminal) ReadLine() (line string, err error) {
	return t.term.ReadLine()
}

func (t *Terminal) Write(b []byte) (n int, err error) {
	return t.term.Write(b)
}

func (t *Terminal) Close() {
	term.Restore(t.fd, t.oldState)
}
