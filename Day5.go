package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("inputs/5.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var maxId uint64 = 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		boardingPass := scanner.Text()
		//println(strings.ReplaceAll(strings.ReplaceAll(boardingPass[:7], "F", "0"), "B", "1"))
		//println(strings.ReplaceAll(strings.ReplaceAll(boardingPass[7:], "L", "0"), "R", "1"))
		rowNumber, _ := strconv.ParseUint(strings.ReplaceAll(strings.ReplaceAll(boardingPass[:7], "F", "0"), "B", "1"), 2, 8)
		columnNumber, _:= strconv.ParseUint(strings.ReplaceAll(strings.ReplaceAll(boardingPass[7:], "L", "0"), "R", "1"), 2, 8)
		//println(rowNumber, columnNumber)
		seatId := rowNumber*8 + columnNumber
		//println(seatId)
		if seatId > maxId {
			maxId = seatId
		}
	}
	println(maxId)
}
