package main

import (
	"os"
	"time"

	"github.com/rinwaowuogba/learn-go-with-tests/math/clockface"
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}
