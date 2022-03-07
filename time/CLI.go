package poker

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"
)

const PlayerPrompt = "Please enter the number of players: "
const BadPlayerInputErrMsg = "Bad value received for number of players, please try again with a number"
const BadGameWinnerInputErrMsg = "Bad value received for game winner, please enter game winner in the format provided"

type Game interface {
	Start(numberOfPlayers int)
	Finish(winner string)
}

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
	c.game.Start(numberOfPlayers)

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

type TexasHoldem struct {
	store   PlayerStore
	alerter BlindAlerter
}

func (t *TexasHoldem) Start(numberOfPlayers int) {
	blindIncrement := time.Duration(5+numberOfPlayers) * time.Minute

	blinds := []int{
		100, 200, 300, 400, 500,
		600, 800, 1000, 2000, 4000,
		8000,
	}
	blindTime := 0 * time.Second
	for _, blind := range blinds {
		t.alerter.ScheduleAlert(blindTime, blind)
		blindTime = blindTime + blindIncrement
	}
}

func (t *TexasHoldem) Finish(name string) {
	t.store.RecordWin(name)
}

func NewTexasHoldemGame(alerter BlindAlerter, store PlayerStore) Game {
	return &TexasHoldem{store, alerter}
}
