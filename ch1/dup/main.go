package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func countLines(f *os.File, counts map[string]int, line2file map[string]string) {
	reader := bufio.NewScanner(f)
	for reader.Scan() {
		counts[reader.Text()]++
		if !strings.Contains(line2file[reader.Text()], f.Name()) {
			line2file[reader.Text()] += f.Name() + " "
		}
	}
}

func main() {
	counts := make(map[string]int)
	line2file := make(map[string]string)

	if len(os.Args) == 1 {
		var fileName = "input.txt"
		f, err := os.Create(fileName)
		if err != nil {
			log.Fatal("Wrong in open file.")
		}
		fmt.Println("Input file content.\nEnter \"#\" to exit")
		reader := bufio.NewScanner(os.Stdin)
		for reader.Scan() {
			if reader.Text() == "#" {
				break
			}
			f.WriteString(reader.Text() + "\n")
		}
		f.Seek(0, 0)
		countLines(f, counts, line2file)
		f.Close()
	} else {
		for _, arg := range os.Args[1:] {
			f, err := os.Open(arg)
			if err != nil {
				log.Fatal("Wrong in open file.")
			}
			countLines(f, counts, line2file)
			f.Close()
		}
	}

	for text, n := range counts {
		if n > 1 {
			fmt.Println(text, n, strings.TrimSuffix(line2file[text], " "))
		}
	}
}
