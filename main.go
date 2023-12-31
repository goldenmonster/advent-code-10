package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var sum int
var lines []string
var listener chan int
var cycleCnt int
var registerValue int

func operationAdd(val int) {
	listener <- 0
	listener <- val
}

func operationNoop() {
	listener <- 0
}

func calcCycle(ch chan int, done chan int) {
	for v := range ch {
		cycleCnt++
		currentPosition := (cycleCnt-1)%40 + 1

		if currentPosition >= registerValue && currentPosition < registerValue+3 {
			fmt.Printf("#")
		} else {
			fmt.Printf(".")
		}

		if cycleCnt%40 == 0 {
			fmt.Printf("\n")
		}
		registerValue += v
	}
}

func init() {
	registerValue = 1
	listener = make(chan int)
}

func main() {

	file, err := os.Open("./in.txt")
	var lines []string
	done := make(chan int)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	go calcCycle(listener, done)

	for _, line := range lines {

		operation := strings.Fields(line)

		if operation[0] == "noop" {
			operationNoop()
		}

		if operation[0] == "addx" {
			val, _ := strconv.Atoi(operation[1])
			operationAdd(val)
		}
	}

	close(listener)
}
