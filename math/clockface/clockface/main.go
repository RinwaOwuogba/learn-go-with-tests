package main

import (
	"os"
	"time"

	"github.com/rinwaowuogba/learn-go-with-tests/math/clockface/svg"
)

func main() {
	t := time.Now()
	svg.Write(os.Stdout, t)
}
