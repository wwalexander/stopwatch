package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

const usage = `usage: stopwatch

stopwatch displays the seconds that have passed since it began running.`

func main() {
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, usage)
	}
	flag.Parse()
	start := time.Now()
	c := time.Tick(100*time.Millisecond)
	for now := range c {
		since := now.Sub(start)
		seconds := since.Seconds()
		fmt.Fprintf(os.Stderr, "\r%.1f", seconds)
	}
}
