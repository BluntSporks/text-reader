// read_text prints text at a human-readable speed with line wrapping.
package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"time"
)

var wordRegExp = regexp.MustCompile(`\pL+('\pL+)*|.`)

func main() {
	// Parse flags.
	wpm := flag.Int("wpm", 300, "Words per minute")
	lineWrap := flag.Int("wrap", 72, "Line wrap")
	flag.Parse()

	// Get filename.
	if len(flag.Args()) < 1 {
		log.Fatal("Missing filename argument")
	}
	file := flag.Arg(0)
	hdl, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer hdl.Close()

	// Calculate delay in milliseconds.
	delay := time.Duration(int(1000.0 / (float32(*wpm) / 60.0)))

	// Process each line.
	scanner := bufio.NewScanner(hdl)
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
		line := scanner.Text()
		words := wordRegExp.FindAllString(line, -1)
		curLen := 0
		for _, word := range words {
			curLen += len(word)

			// Only print a newline if there is more than one word, the line would be longer than the wrap,
			// and the word length is greater than 1, which excludes spaces and punctuation marks.
			if curLen > len(word) && curLen > *lineWrap && len(word) > 1 {
				fmt.Println()
				curLen = len(word)
			}
			fmt.Printf("%s", word)
			if len(word) > 1 {
				time.Sleep(delay * time.Millisecond)
			}
		}
		fmt.Println()
	}
}
