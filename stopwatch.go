package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

const usage = `usage: stopwatch

stopwatch displays the seconds and milliseconds since it began running.`

func main() {
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, usage)
	}
	flag.Parse()
	start := time.Now()
	c := time.Tick(25 * time.Millisecond)
	for now := range c {
		since := now.Sub(start)
		fmt.Fprintf(os.Stderr, "\r%.3f", since.Seconds())
	}
}
