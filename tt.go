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

func ScanLine(num int) string {

	police1, err := os.Open("standard.txt")
	police2, err := os.Open("shadow.txt")
	scanner1 := bufio.NewScanner(police1)
	scanner2 := bufio.NewScanner(police2)
	checkLine := 0
	printLine := ""

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}

	for scanner1.Scan() {

		if checkLine == num {
			printLine = scanner1.Text()
		}
		checkLine++

	}

	for scanner2.Scan() {

		if checkLine == num {
			printLine = scanner2.Text()
		}
		checkLine++

	}

	return printLine

}
