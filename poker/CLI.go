package poker

import (
	"bufio"
	"strings"
	"io"
)

type CLI struct {
	Store Collection
	In	*bufio.Scanner
}

func NewCLI(store Collection, in io.Reader) *CLI {
	return &CLI{store, bufio.NewScanner(in)}
}

func(c *CLI) PlayPoker() {
	userInput := c.readLine()
	c.Store.RecordWin(extractWinner(userInput))
}

func extractWinner(userInput string) string {
	return strings.Replace(userInput, " wins", "", 1)
}

func(c *CLI) readLine() string {
	c.In.Scan()
	return c.In.Text()
}

