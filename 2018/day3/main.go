package main

import (
	"fmt"
	"regexp"
	"strconv"
)

var /* const */ CLAIM_PATTERN = regexp.MustCompile(`#(\d+)\s@\s(\d+),(\d+):\s(\d+)x(\d+)`)

type Claim struct {
	id     int
	left   int
	top    int
	width  int
	height int
}

func (claim Claim) Draw(fabric *[1000][1000]byte) {
	for x := claim.left; x < claim.left+claim.width; x++ {
		for y := claim.top; y < claim.top+claim.height; y++ {
			if fabric[x][y] != 'C' && fabric[x][y] != 'X' {
				fabric[x][y] = 'C'
			} else {
				fabric[x][y] = 'X'
			}
		}
	}
}

func (claim Claim) HasConflict(fabric *[1000][1000]byte) bool {
	for x := claim.left; x < claim.left+claim.width; x++ {
		for y := claim.top; y < claim.top+claim.height; y++ {
			if fabric[x][y] == 'X' {
				return true
			}
		}
	}

	return false
}

func parseClaim(line string) Claim {
	match := CLAIM_PATTERN.FindStringSubmatch(line)[1:]

	id, _ := strconv.Atoi(match[0])
	left, _ := strconv.Atoi(match[1])
	top, _ := strconv.Atoi(match[2])
	width, _ := strconv.Atoi(match[3])
	height, _ := strconv.Atoi(match[4])

	return Claim{id, left, top, width, height}
}

func main() {
	var claims []Claim
	var fabric [1000][1000]byte
	for _, line := range readLines() {
		claim := parseClaim(line)
		claims = append(claims, claim)
		claim.Draw(&fabric)
	}

	var conflicts int
	for _, row := range fabric {
		for _, value := range row {
			if value == 'X' {
				conflicts++
			}
		}
	}

	fmt.Println("Square inches of fabric are within two or more claims: ", conflicts)
	for _, claim := range claims {
		hasConflict := claim.HasConflict(&fabric)
		if !hasConflict {
			fmt.Printf("Claim #%d has no conflict!\n", claim.id)
		}
	}
}
