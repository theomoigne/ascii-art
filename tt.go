package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	normal            = "\033[0m"
	grave             = "\033[1m"
	soulignerGrave    = "\033[21m"
	fondBlanc         = "\033[7m"
	colorBlack        = "\033[30m"
	colorRedSombre    = "\033[31m"
	colorGreenSombre  = "\033[32m"
	colorYellowSombre = "\033[33m"
	colorBlueSombre   = "\033[34m"
	colorPurpleSombre = "\033[35m"
	colorCyanSombre   = "\033[36m"
	colorRed          = "\033[91m"
	colorGreen        = "\033[92m"
	colorYellow       = "\033[93m"
	colorBlue         = "\033[94m"
	colorRose         = "\033[95m"
	colorCyan         = "\033[96m"
	colorWhite        = "\033[97m"
	fondNoir          = "\033[40m"
	fondRouge         = "\033[41m"
	fondVert          = "\033[42m"
	fondJaune         = "\033[43m"
	fondBleu          = "\033[44m"
	fondViolet        = "\033[45m"
	fondCyan          = "\033[46m"
	fondGris          = "\033[100m"
	fondRoseRouge     = "\033[101m"
	fondRose          = "\033[105m"
	fondBleuCiel      = "\033[106m"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("choisie ta police de text et couleur :")
	scanner.Scan()
	choix, err := strconv.Atoi(scanner.Text())
	if err != nil {
		if choix < 1 || choix > 4 {
			fmt.Println("... Ca n'existe pas. >.>")
	if choix == 1 {
		fmt.Println("vous avez choise entre la polise stantard, shadow et thinkertoy : ")
		scanner.Scan()
		police, err := strconv.Atoi(scanner.Text())
		if police == 1 {
			s := os.Args[1]
			for _, v := range os.Args[2:] {
				s += " " + v
			}
			result := ""

			for i := 0; i < 8; i++ {

				for _, v := range s {

					result += ScanLine(1 + int(v-32)*9 + i)

				}
			}
		}
		if police == 2 {
			s := os.Args[1]
			for _, v := range os.Args[2:] {
				s += " " + v
			}
			result := ""

			for i := 0; i < 8; i++ {

				for _, v := range s {

					result += ScanLine2(1 + int(v-32)*9 + i)

				}
			}
		}
		if police == 3 {
			s := os.Args[1]
			for _, v := range os.Args[2:] {
				s += " " + v
			}
			result := ""

			for i := 0; i < 8; i++ {

				for _, v := range s {

					result += ScanLine3(1 + int(v-32)*9 + i)

				}
			}
		}


		if err != nil {
			if police < 1 || police > 4 {
				fmt.Println("... Ca n'existe pas. >.>")
			}
			fmt.Println(err.Error())
			os.Exit(0)
		}


		if choix == 2 {
			fmt.Println("vous avez choise entre la polise stantard, shadow et thinkertoy : ")
			scanner.Scan()
			police, err := strconv.Atoi(scanner.Text())
			if couleur == 1 {
				result := ""

				for i := 0; i < 8; i++ {

					for _, v := range s {

						result += ScanLine(1 + int(v-32)*9 + i)

					}
				}
			}
			if couleur == 2 {
				s := os.Args[1]
				for _, v := range os.Args[2:] {
					s += " " + v
				}
				result := ""

				for i := 0; i < 8; i++ {

					for _, v := range s {

						result += ScanLine2(1 + int(v-32)*9 + i)

					}
				}
			}
			if police == 3 {
				s := os.Args[1]
				for _, v := range os.Args[2:] {
					s += " " + v
				}
				result := ""

				for i := 0; i < 8; i++ {

					for _, v := range s {

						result += ScanLine3(1 + int(v-32)*9 + i)

					}
				}
			}


			if err != nil {
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

fmt.Println(result)
result = ""

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

func color(result string) {
	fmt.Println("choisie ta couleur :")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	couleur, err := strconv.Atoi(scanner.Text())
	fmt.Println("Vous allez choisir la couleur")
	if couleur == 1 {
		fmt.Println(colorBlue + result)
	}
	if couleur == 2 {
		fmt.Println(colorBlack)
	}
	if err != nil {
		fmt.Println("... 1, 2 ou TROIS ! Pas de(s) lettre(s) !")
		if couleur < 1 || couleur > 4 {
			fmt.Println("... Ca n'existe pas. >.>")
		}
		fmt.Println(err.Error())
		os.Exit(0)
	}
}
