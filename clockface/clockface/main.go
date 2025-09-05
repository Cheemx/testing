package main

import (
	"os"
	"time"

	"github.com/Cheemx/testing/clockface"
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}
