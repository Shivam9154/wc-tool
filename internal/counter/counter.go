package counter

import (
	"bufio"
	"io"
	"unicode"
)

// Counter tracks lines, words, bytes, and characters in a text stream.
type Counter struct {
	lines      int
	words      int
	bytes      int
	characters int
}

// New initializes and returns a new Counter instance.
func New() *Counter {
	return &Counter{}
}

func (c *Counter) Process(r io.Reader) error {
	reader := bufio.NewReader(r)

	inWord := false
	for {
		rune, size, err := reader.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		c.bytes += size
		c.characters++

		if rune == '\n' {
			c.lines++
		}

		if unicode.IsSpace(rune) {
			inWord = false
		} else {
			if !inWord {
				c.words++
				inWord = true
			}
		}
	}

	return nil
}

func (c *Counter) Lines() int {
	return c.lines
}

func (c *Counter) Words() int {
	return c.words
}

func (c *Counter) Bytes() int {
	return c.bytes
}

func (c *Counter) Characters() int {
	return c.characters
}
