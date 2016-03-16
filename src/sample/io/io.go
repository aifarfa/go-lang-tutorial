package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	readFile()
}

func readFile() {
	f := openFile()
	defer f.Close()
	// if f, err := os.Open("input.json"); err == nil {
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		log.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	// }
}

func openFile() *os.File {
	f, err := os.Open("input.json")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	return f
}
