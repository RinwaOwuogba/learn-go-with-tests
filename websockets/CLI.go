package poker

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

const PlayerPrompt = "Please enter the number of players: "
const BadPlayerInputErrMsg = "Bad value received for number of players, please try again with a number"
const BadGameWinnerInputErrMsg = "Bad value received for game winner, please enter game winner in the format provided"

type CLI struct {
	in   *bufio.Scanner
	out  io.Writer
	game Game
}

func (c *CLI) PlayPoker() {
	fmt.Fprint(c.out, PlayerPrompt)

	numberOfPlayers, err := strconv.Atoi(c.readLine())
	if err != nil {
		fmt.Fprint(c.out, BadPlayerInputErrMsg)
		return
	}
	c.game.Start(numberOfPlayers, os.Stdout)

	winnerInput := c.readLine()
	winner, err := extractWinner(winnerInput)
	if err != nil {
		fmt.Fprint(c.out, BadGameWinnerInputErrMsg)
		return
	}

	c.game.Finish(winner)
}

func extractWinner(userInput string) (string, error) {
	if !strings.Contains(userInput, " wins") {
		return "", fmt.Errorf("invalid winner message format")
	}

	return strings.Replace(userInput, " wins", "", 1), nil
}

func (c *CLI) readLine() string {
	c.in.Scan()
	return c.in.Text()
}

func NewCLI(in io.Reader, out io.Writer, game Game) *CLI {
	return &CLI{
		in:   bufio.NewScanner(in),
		out:  out,
		game: game,
	}
}
