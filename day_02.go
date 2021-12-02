package main

import (
	_ "embed"
	"strconv"
	"strings"
)

//go:embed inputs/day_02.txt
var commandsFile string

type cmd struct {
	direction string
	magnitude int
}

func parseCommands() <-chan *cmd {
	c := make(chan *cmd)
	go func() {
		for _, line := range strings.Split(commandsFile, "\n") {
			s := strings.Split(line, " ")
			m, _ := strconv.Atoi(s[1])
			c <- &cmd{s[0], m}
		}
		close(c)
	}()
	return c
}

func dive() int {
	var x, y int
	for cmd := range parseCommands() {
		switch cmd.direction {
		case "forward":
			x += cmd.magnitude
		case "down":
			y += cmd.magnitude
		case "up":
			y -= cmd.magnitude
		}
	}
	return x * y
}

func diveWithAim() int {
	var x, y, aim int
	for cmd := range parseCommands() {
		switch cmd.direction {
		case "forward":
			x += cmd.magnitude
			y += cmd.magnitude * aim
		case "down":
			aim += cmd.magnitude
		case "up":
			aim -= cmd.magnitude
		}
	}
	return x * y
}
