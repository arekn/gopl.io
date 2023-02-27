// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 6.
//!+

// Echo2 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	//echo1
	start := time.Now()
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
	sec1 := time.Since(start).Seconds()

	//echo2
	start2 := time.Now()
	s2, sep2 := "", ""
	for _, arg := range os.Args[1:] {
		s2 += sep2 + arg
		sep2 = " "
	}
	fmt.Println(s2)
	sec2 := time.Since(start2).Seconds()

	//echo3
	start3 := time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	sec3 := time.Since(start3).Seconds()

	fmt.Printf("echo1: %fs, echo2 %fs, echo3 %fs", sec1, sec2, sec3)
}

//!-
