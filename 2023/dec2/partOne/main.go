package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	gameFile, err := os.Open("../gameresults.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer gameFile.Close()

	scanner := bufio.NewScanner(gameFile)
	gameCounter := 1
	sumTotal := 0

	for scanner.Scan() {
		addGame := true
		gameSplit := strings.Split(scanner.Text(), ": ")
		println("Game Rounds: ", gameSplit[1])
		roundSplit := strings.Split(gameSplit[1], "; ")

		for _, val := range roundSplit {
			println("Rounds: ", val)
			colors := strings.Split(val, ", ")

			for _, colorVal := range colors {
				colorResults := strings.Split(colorVal, " ")
				colorCount, err := strconv.Atoi(colorResults[0])
				if err != nil {
					log.Fatal(err)
				}

				switch colorResults[1] {
				case "red":
					if colorCount > 12 {
						addGame = false
						println("[FAIL] Due to red > 12")
					}
				case "blue":
					if colorCount > 14 {
						addGame = false
						println("[FAIL] Due to blue > 14")
					}
				case "green":
					if colorCount > 13 {
						addGame = false
						println("[FAIL] Due to green > 13")
					}
				}

			}

		}

		if addGame {
			println("[SUCCESS] Adding ", gameCounter, " to total: ", sumTotal)
			sumTotal += gameCounter
		}
		gameCounter++
	}
	println("Total sum of games: ", sumTotal)
}
