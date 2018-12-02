package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	var ids []string
	statistics := make(map[int]int)

	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		ids = append(ids, line)
		// fmt.Println(line)

		counts := make(map[rune]int)
		for _, rune := range line {
			counts[rune]++
		}

		// fmt.Println(counts)
		for _, value := range counts {
			if value == 2 {
				statistics[2]++
				break
			}
		}

		for _, value := range counts {
			if value == 3 {
				statistics[3]++
				break
			}
		}

		// fmt.Println(statistics)
	}

	fmt.Println(statistics[2] * statistics[3])
	fmt.Println(locate(ids))
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
