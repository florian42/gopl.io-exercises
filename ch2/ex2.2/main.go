package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/florian42/gopl.io-exercises/ch2/ex2.2/convert"
)

func main() {
	fmt.Println("Converter Utility...")
	if len(os.Args) > 2 {
		for _, arg := range os.Args[1:] {
			printUnits(arg)
		}
	} else {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Please enter numbers seperated by space: ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", 1)
		numbers := strings.Split(text, " ")
		for _, number := range numbers {
			printUnits(number)
		}
	}
}
func printUnits(number string) {
	fmt.Printf("number: %s\n", number)
	i, err := strconv.ParseFloat(number, 64)
	if err != nil {
		fmt.Printf("Please use a number, not %s, err: %v\n", number, err)
	} else {
		fmt.Printf("%v\n", convert.CToF(convert.Celsius(i)))
		fmt.Printf("%v\n", convert.FToC(convert.Fahrenheit(i)))
		fmt.Printf("%v\n", convert.MToF(convert.Meter(i)))
		fmt.Printf("%v\n", convert.FToM(convert.Feet(i)))
		fmt.Printf("%v\n", convert.KToP(convert.Kilogram(i)))
		fmt.Printf("%v\n", convert.PToK(convert.Pound(i)))
	}
}
