package main

import (
	"flag"
	"fmt"
	"jex/progog/internal/progog"
	"os"
)

func main() {
	startHash := flag.String("start", "", "Starting block hash")
	depth := flag.Int("depth", 1, "Depth of traversal")
	retry := flag.Int("retry", 3, "Maximum number of retries for HTTP requests")
	flag.Parse()

	if *startHash == "" || *depth < 1 || *retry < 1 {
		throwExit()
	}

	progog.SetMaxRetries(*retry)
	progog.Start(*startHash, *depth)
	progog.ExportToKB("out.pl")
}

func throwExit() {
	fmt.Println("Usage: progog -start <block_hash> -depth <depth> -retry <max_retries>")
	os.Exit(1)
}
