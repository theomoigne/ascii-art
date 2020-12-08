package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	Info = Teal
	Warn = Yellow
	Fata = Red
)

var (
	Black   = Color("\033[1;30m%s\033[0m")
	Red     = Color("\033[1;31m%s\033[0m")
	Green   = Color("\033[1;32m%s\033[0m")
	Yellow  = Color("\033[1;33m%s\033[0m")
	Purple  = Color("\033[1;34m%s\033[0m")
	Magenta = Color("\033[1;35m%s\033[0m")
	Teal    = Color("\033[1;36m%s\033[0m")
	White   = Color("\033[1;37m%s\033[0m")
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("choisie ta police de text :")
	scanner.Scan()
	police, err := strconv.Atoi(scanner.Text())

	scanner2 := bufio.NewScanner(os.Stdin)
	fmt.Println("choisie ta couleur de text :")
	scanner2.Scan()
	couleur, err := strconv.Atoi(scanner.Text())

	if err != nil {
		fmt.Println("Pas de(s) lettre(s) !")
		if couleur < 0 || couleur > 16 {
			fmt.Println("Ca n'existe pas. >.>")
		}
		fmt.Println(err.Error())
		os.Exit(0)
	}

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

		if couleur == 1 {
			fmt.Println(Red(result))
			result = ""
		}

		if couleur == 2 {
			fmt.Println(Green(result))
			result = ""
		}
	}
}

func Color(colorString string) func(...interface{}) string {
	sprint := func(args ...interface{}) string {
		return fmt.Sprintf(colorString,
			fmt.Sprint(args...))
	}
	return sprint
}


func ScanLine(num int) string {

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

func ScanLine2(num int) string {

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

func ScanLine3(num int) string {

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
