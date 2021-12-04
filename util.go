package main

import "strings"

func parseFile(f string) <-chan string {
	c := make(chan string)
	go func() {
		for _, line := range strings.Split(f, "\n") {
			c <- line
		}
		close(c)
	}()
	return c
}
