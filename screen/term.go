package screen

import (
	"bufio"
	"fmt"
	"os"
)

type Screen interface {
	Clear()
	Flush()
	DrawAt(c string, x, y int)
}

var screen = bufio.NewWriter(os.Stdout)

func NewScreen() Screen {
	return &term{}
}

type term struct{}

func (t *term) DrawAt(c string, x, y int) {
	t.moveCursor(x, y)
	t.draw(c)
}

// clear screen
func (t *term) Clear() {
	fmt.Fprint(screen, "\033[2J")
}

// move cursor to x, y
func (t *term) moveCursor(x, y int) {
	fmt.Fprintf(screen, "\033[%d;%dH", y, x)

}

// draw rune 'c' on the terminal
func (t *term) draw(c string) {
	fmt.Fprint(screen, c)
}

// writes on the screen
func (t *term) Flush() {
	screen.Flush()
}
