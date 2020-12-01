package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("inputs/1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var values []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		//fmt.Println(scanner.Text())
		// Let's just assume these are all ints and we ignore the err value
		i, _ := strconv.Atoi(scanner.Text())
		values = append(values, i)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(values); i++ {
		for j := i+1; j < len(values); j++ {
			//fmt.Println(values[i]+values[j])
			if v:= values[i]+values[j]; v == 2020 {
				fmt.Println(values[i]*values[j])
			}
		}
	}
}
