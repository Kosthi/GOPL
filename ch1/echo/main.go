package main

import (
	"fmt"
	"os"
	"strings"
	time2 "time"
)

func main() {
	var s, sep string

	begin1 := time2.Now()
	for _, arg := range os.Args {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
	end1 := time2.Now()
	fmt.Println(end1.Sub(begin1))

	s = ""
	begin2 := time2.Now()
	s += os.Args[0]
	for _, arg := range os.Args[1:] {
		s += " " + arg
	}
	fmt.Println(s)
	end2 := time2.Now()
	fmt.Println(end2.Sub(begin2))

	begin3 := time2.Now()
	fmt.Println(strings.Join(os.Args, " "))
	end3 := time2.Now()
	fmt.Println(end3.Sub(begin3))
}
