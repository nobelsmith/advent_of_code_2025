package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func AbsInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	// Define the starting position of the dial
	dial := 50
	times_pointing_at_zero := 0
	dial_start_position := 50
	fmt.Println(dial)
	// Read in the file
	file, err := os.Open("../data/input.txt")
	// file, err := os.Open("../data/sample_input.txt")
	// file, err := os.Open("../data/sample_input2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		instruction := line[0]
		total_number_of_clicks, err := strconv.Atoi(line[1:])
		functional_number_of_clicks := 0
		if total_number_of_clicks <= 99 {
			functional_number_of_clicks = total_number_of_clicks
		} else {
			fmt.Printf("Adding multiple times pointing at zero %d\n", total_number_of_clicks/100)
			times_pointing_at_zero += total_number_of_clicks / 100
			functional_number_of_clicks = total_number_of_clicks % 100
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("line: %s, total number of clicks: %d, functional number of clicks: %d\n", line, total_number_of_clicks, functional_number_of_clicks)
		var dial_after_instruction int
		if instruction == 'L' {
			fmt.Println("LEFT")
			dial_after_instruction = dial - functional_number_of_clicks
		}
		if instruction == 'R' {
			fmt.Println("RIGHT")
			dial_after_instruction = dial + functional_number_of_clicks
		}
		fmt.Printf("dial after instruction is applied: %d\n", dial_after_instruction)
		if dial_after_instruction == 100 {
			dial = 0
		} else if dial_after_instruction < 0 {
			dial = 100 - AbsInt(dial_after_instruction)
			if dial_start_position != 0 {
				times_pointing_at_zero += 1
			}
		} else if dial_after_instruction > 100 {
			functional_dial_turns := dial_after_instruction % 100
			dial = 0 + functional_dial_turns
			if dial_start_position != 0 {
				times_pointing_at_zero += 1
			}
		} else {
			dial = dial_after_instruction
		}
		fmt.Printf("dial after enforcing circle: %d\n", dial)
		fmt.Printf("Dial: %d\n", dial)
		if dial == 0 {
			times_pointing_at_zero += 1
			fmt.Printf("dial ended on zero: %d\n", times_pointing_at_zero)
		}
		dial_start_position = dial
		fmt.Printf("Times Every Pointing at zero %d\n", times_pointing_at_zero)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file, ", err)
	}
	fmt.Printf("dial pointed at zero %d times total", times_pointing_at_zero)
}
