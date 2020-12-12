package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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
	var seatList []uint64

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		boardingPass := scanner.Text()
		//println(strings.ReplaceAll(strings.ReplaceAll(boardingPass[:7], "F", "0"), "B", "1"))
		//println(strings.ReplaceAll(strings.ReplaceAll(boardingPass[7:], "L", "0"), "R", "1"))
		rowNumber, _ := strconv.ParseUint(strings.ReplaceAll(strings.ReplaceAll(boardingPass[:7], "F", "0"), "B", "1"), 2, 8)
		columnNumber, _:= strconv.ParseUint(strings.ReplaceAll(strings.ReplaceAll(boardingPass[7:], "L", "0"), "R", "1"), 2, 8)
		totalNumber, _ := strconv.ParseUint(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(boardingPass, "F", "0"), "B", "1"), "L", "0"), "R", "1"), 2, 16)
		println(totalNumber)
		seatList = append(seatList, totalNumber)
		//println(rowNumber, columnNumber)
		seatId := rowNumber*8 + columnNumber
		//println(seatId)
		if seatId > maxId {
			maxId = seatId
		}
	}
	println(maxId)

	sort.Slice(seatList, func(i, j int) bool { return seatList[i] < seatList[j]})
	println(seatList)
	fmt.Printf("%v\n", seatList)
	idx := 1
	for i:= seatList[0]+1; i < maxId; i++ {
		if seatList[idx] != i {
			println(i)
			break
		} else {
			idx++
		}
	}
}
