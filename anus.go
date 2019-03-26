// Command anus.io emits excrement on standard output.
package main

import (
	"os"
	"time"
)

func main() {
	for {
		os.Stdout.Write([]byte("\U0001F4A9"))
		time.Sleep(100 * time.Millisecond)
	}
}
