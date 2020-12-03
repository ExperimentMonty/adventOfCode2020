package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("inputs/2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	validCount := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		//fmt.Println(scanner.Text())
		args := strings.Split(scanner.Text(), ":")
		policy := args[0]
		password := args[1]

		if validPassword(policy, password) {
			validCount++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	println(validCount)
}

func validPassword (policy string, password string) bool {
	components := strings.Split(policy, " ")
	numberRange := strings.Split(components[0], "-")
	minimum, _ := strconv.Atoi(numberRange[0])
	maximum, _ := strconv.Atoi(numberRange[1])

	character := components[1]

	count := strings.Count(password, character)

	return count >= minimum && count <= maximum



}
