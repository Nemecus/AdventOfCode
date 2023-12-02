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
var numberWords = [9][2]string{
	{"1", "one"},
	{"2", "two"},
	{"3", "three"},
	{"4", "four"},
	{"5", "five"},
	{"6", "six"},
	{"7", "seven"},
	{"8", "eight"},
	{"9", "nine"},
}

func main() {

	calibrationFile, err := os.Open("calibration.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer calibrationFile.Close()

	scanner := bufio.NewScanner(calibrationFile)

	for scanner.Scan() {
		calibrationLine := scanner.Text()

		println("Line is: ", calibrationLine)

		for _, val := range numberWords {
			if strings.Contains(calibrationLine, val[1]) {
				calibrationLine = strings.ReplaceAll(calibrationLine, val[1], val[0])
				println("Calibration line changed: ", calibrationLine)
			}
		}
		println("New Calibration Line: ", calibrationLine)
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
