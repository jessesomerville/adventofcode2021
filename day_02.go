package main

import (
	_ "embed"
	"strconv"
	"strings"
)

//go:embed inputs/day_02.txt
var commandsFile string

func dive() int {
	var x, y int
	for _, line := range strings.Split(commandsFile, "\n") {
		cmd := strings.Split(line, " ")
		units, _ := strconv.Atoi(cmd[1])

		switch cmd[0] {
		case "forward":
			x += units
		case "down":
			y += units
		case "up":
			y -= units
		}
	}
	return x * y
}

func diveWithAim() int {
	var x, y, aim int
	for _, line := range strings.Split(commandsFile, "\n") {
		cmd := strings.Split(line, " ")
		units, _ := strconv.Atoi(cmd[1])

		switch cmd[0] {
		case "forward":
			x += units
			y += units * aim
		case "down":
			aim += units
		case "up":
			aim -= units
		}
	}
	return x * y
}
