package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("inputs/4.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	validCount := 0
	validatedCount := 0

	validFields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	validEyeColors := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
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

			// Solve for part 1
			allFieldsPresent := true
			for _, validField := range validFields {
				if _, ok := passport[validField]; !ok {
					allFieldsPresent = false
				}
			}

			if allFieldsPresent {
				validCount++

				// Solve for part 2 (only relevant if part 1 is true)
				allFieldsValid := true
				for key, value := range passport {
					if !allFieldsValid {
						break
					}

					switch key {
					case "byr":
						year, _ := strconv.Atoi(value)
						allFieldsValid = allFieldsValid && year >= 1920 && year <= 2002
					case "iyr":
						year, _ := strconv.Atoi(value)
						allFieldsValid = allFieldsValid && year >= 2010 && year <= 2020
					case "eyr":
						year, _ := strconv.Atoi(value)
						allFieldsValid = allFieldsValid && year >= 2020 && year <= 2030
					case "hgt":
						// Check suffix
						if strings.HasSuffix(value, "in") {
							measurement, err := strconv.Atoi(value[:len(value)-2])
							if err != nil {
								allFieldsValid = false
							} else {
								allFieldsValid = allFieldsValid && measurement >= 59 && measurement <= 76
							}
						} else if strings.HasSuffix(value, "cm") {
							measurement, err := strconv.Atoi(value[:len(value)-2])
							if err != nil {
								allFieldsValid = false
							} else {
								allFieldsValid = allFieldsValid && measurement >= 150 && measurement <= 193
							}
						} else {
							allFieldsValid = false
						}
					case "hcl":
						found, err := regexp.MatchString("^#[0-9a-f]{6}$", value)
						if err != nil {
							allFieldsValid = false
						}
						allFieldsValid = allFieldsValid && found
					case "ecl":
						foundEyeColor := false
						for _, eyeColor := range validEyeColors {
							if value == eyeColor {
								foundEyeColor = true
							}
						}
						allFieldsValid = allFieldsValid && foundEyeColor
					case "pid":
						found, err := regexp.MatchString("^[0-9]{9}$", value)
						if err != nil {
							allFieldsValid = false
						}
						allFieldsValid = allFieldsValid && found

					}

				}
				if allFieldsValid {
					validatedCount++
				}
			}

			passport = make(map[string]string)
		}

	}

	println(validCount)
	println(validatedCount)
}
