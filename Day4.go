package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("inputs/4.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	validCount := 0

	validFields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	var passport = make(map[string]string)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		inputLine := scanner.Text()
		if inputLine != "" {
			fields := strings.Fields(inputLine)
			for _, field := range fields {
				// I'm going to try this another way so I don't keep learning the same strings.Split syntax
				// Find delimiter
				delimiter := strings.Index(field, ":")
				passport[field[:delimiter]] = field[delimiter+1:]
			}
		} else {

			allFieldsPresent := true
			for _, validField := range validFields {
				if _, ok := passport[validField]; !ok {
					allFieldsPresent = false
				}
			}

			if allFieldsPresent {
				validCount++
			}

			passport = make(map[string]string)
		}

	}

	println(validCount)
}
