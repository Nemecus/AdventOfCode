package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var calibrationSum int = 0

func main() {

	calibrationFile, err := os.Open("../calibration.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer calibrationFile.Close()

	scanner := bufio.NewScanner(calibrationFile)

	for scanner.Scan() {
		calibrationLine := scanner.Text()

		println("Line is: ", calibrationLine)

		re, err := regexp.Compile(`[-]?\d[\d,]*[\.]?[\d{2}]*`)
		if err != nil {
			log.Fatal(err)
		}

		numStrings := re.FindAllString(calibrationLine, -1)
		calibrationNumbers := strings.Join(numStrings, "")
		firstNum := string(calibrationNumbers[0])
		secondNum := string(calibrationNumbers[len(calibrationNumbers)-1])
		calibrationNumbers = firstNum + secondNum

		println("Adding numbers: ", calibrationNumbers, "to sum: ", calibrationSum)

		tempNumbers, err := strconv.Atoi(calibrationNumbers)
		if err != nil {
			log.Fatal(err)
		}
		calibrationSum += tempNumbers

	}
	println("Sum of Calibrations: ", calibrationSum)
}
