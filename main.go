package main

import (
	"bufio"
	"fmt"
	"github.com/fatih/color"
	"os"
	"strconv"
)

func main() {
 	fmt.Println("il y a 3 type de police de texte.")
 	fmt.Println("il y a que 8 type de couleur de texte : \n 1 = blanc \n 2 = noir \n 3 = rouge \n 4 = bleu \n 5 = vert \n 6 = jaune \n 7 = violet \n 8 = bleu ciel  ")
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("choisie ta police de text :")
	scanner.Scan()
	police, err := strconv.Atoi(scanner.Text())

	if err != nil {
		fmt.Println("... 1, 2 ou TROIS ! Pas de(s) lettre(s) !")
		if police < 0 || police > 4 {
			fmt.Println("... Ca n'existe pas. >.>")
		}
		fmt.Println(err.Error())
		os.Exit(0)
	}

	scanner2 := bufio.NewScanner(os.Stdin)
	fmt.Println("choisie ta couleur de text :")
	scanner2.Scan()
	couleur, err := strconv.Atoi(scanner2.Text())

	if err != nil {
		fmt.Println("Pas de(s) lettre(s) !")
		if couleur < 1 || couleur > 8 {
			fmt.Println("Ca n'existe pas. >.>")
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
			a := color.New(color.FgWhite)
			a.Println(result)
			result = ""
		}

		if couleur == 2 {
			a := color.New(color.FgBlack)
			a.Println(result)
			result = ""
		}

		if couleur == 3 {
			a := color.New(color.FgRed)
			a.Println(result)
			result = ""
		}

		if couleur == 4 {
			a := color.New(color.FgBlue)
			a.Println(result)
			result = ""
		}

		if couleur == 5 {
			a := color.New(color.FgGreen)
			a.Println(result)
			result = ""
		}

		if couleur == 6 {
			a := color.New(color.FgYellow)
			a.Println(result)
			result = ""
		}

		if couleur == 7 {
			a := color.New(color.FgMagenta)
			a.Println(result)
			result = ""
		}

		if couleur == 8 {
			a := color.New(color.FgCyan)
			a.Println(result)
			result = ""
		}
	}
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
