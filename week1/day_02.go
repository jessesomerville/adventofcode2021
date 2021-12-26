package week1

import (
	"strconv"
	"strings"
)

type cmd struct {
	direction string
	magnitude int
}

func Dive(f string) int {
	var x, y int
	for cmd := range ParseCommands(f) {
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

func DiveWithAim(f string) int {
	var x, y, aim int
	for cmd := range ParseCommands(f) {
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

func ParseCommands(f string) <-chan *cmd {
	c := make(chan *cmd)
	go func() {
		for _, line := range strings.Split(f, "\n") {
			s := strings.Split(line, " ")
			m, _ := strconv.Atoi(s[1])
			c <- &cmd{s[0], m}
		}
		close(c)
	}()
	return c
}
