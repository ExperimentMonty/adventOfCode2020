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
	validCount2 := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		//fmt.Println(scanner.Text())
		args := strings.Split(scanner.Text(), ":")
		policy := args[0]
		password := args[1]

		if validPassword(policy, password) {
			validCount++
		}

		if validPassword2(policy, password) {
			validCount2++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	println(validCount)
	println(validCount2)
}

func validPassword(policy string, password string) bool {
	components := strings.Split(policy, " ")
	numberRange := strings.Split(components[0], "-")
	minimum, _ := strconv.Atoi(numberRange[0])
	maximum, _ := strconv.Atoi(numberRange[1])

	character := components[1]

	count := strings.Count(password, character)

	return count >= minimum && count <= maximum

}

func validPassword2(policy string, password string) bool {
	components := strings.Split(policy, " ")
	slots := strings.Split(components[0], "-")
	slot1, _ := strconv.Atoi(slots[0])
	slot2, _ := strconv.Atoi(slots[1])

	// Ordinarily, we'd need to do this zero-index nonsense, but our separation left a preceding space
	// No zero index
	//slot1--
	//slot2--

	// Get this as a character instead of a string
	character := components[1][0]

	return (password[slot1] == character && password[slot2] != character) ||
		(password[slot1] != character && password[slot2] == character)
}
