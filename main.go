package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		return
	}
	s := os.Args[1]
	for _, v := range os.Args[2:] {
		s += " " + v
	}
	result := ""

	for i := 0; i < 8; i++ {

		for _, v := range s {

			result += ScanLine(1 + int(v-32)*9 + i)

		}
		fmt.Println(result)
		result = ""
	}
}

func ScanLine2(num int) string {

	fileName, err := os.Open("standard.txt")
	scanner := bufio.NewScanner(fileName)
	checkLine := 0
	printLine := ""

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}

	for scanner.Scan() {

		if checkLine == num {
			printLine = scanner.Text()
		}
		checkLine++

	}

	return printLine
}