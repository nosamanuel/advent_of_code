package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	frequency := 0
	history := make(map[int]int)
	history[0] = 1
	for true {
		file, err := os.Open("./input.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			// fmt.Println(line)
			number, err := strconv.Atoi(line)
			if err != nil {
				log.Fatal(err)
			}

			frequency += number
			history[frequency] += 1
			fmt.Println(frequency, number, history[frequency])
			if history[frequency] >= 2 {
				fmt.Println(frequency)
				return
			}
		}

		// fmt.Println(frequency)
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}
}
