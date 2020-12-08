package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("choisie ta police de text :")
	scanner.Scan()
	police, err := strconv.Atoi(scanner.Text())

	if err != nil {
		fmt.Println("... 1, 2 ou TROIS ! Pas de(s) lettre(s) !")
		if police < 1 || police > 4 {
			fmt.Println("... Ca n'existe pas. >.>")
		}
		fmt.Println(err.Error())
		os.Exit(0)
	}

	if len(os.Args) == 1 {
		return
	}
	s := os.Args[1]
	for _, v := range os.Args[2:] {
		s += " " + v
	}
	result := ""

	for i := 0; i < 8; i++ {

		if police == 1 {

			for _, v := range s {

				result += ScanLine(1 + int(v-32)*9 + i)

			}
		}

		if police == 2 {

			for _, v := range s {

				result += ScanLine2(1 + int(v-32)*9 + i)

			}
		}

		if police == 3 {

			for _, v := range s {

				result += ScanLine3(1 + int(v-32)*9 + i)

			}
		}

		fmt.Println(result)
		result = ""
	}
}

func ScanLine55(num int) string {

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

func ScanLine255(num int) string {

	fileName, err := os.Open("shadow.txt")
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

func ScanLine355(num int) string {

	fileName, err := os.Open("thinkertoy.txt")
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
