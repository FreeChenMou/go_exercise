package main

import (
	"fmt"
	"go_code/go_exercise/go_programming_2/tempconv"
	"os"
	"strconv"
)

//exercise2.2 converts type
func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		// converts celsius to fahrenheit and fahrenheit to celsius
		celsius := tempconv.Celsius(t)
		fahrenheit := tempconv.Fahrenheit(t)
		fmt.Printf("%s = %s, %s = %s\n",
			celsius, tempconv.CToF(celsius), fahrenheit, tempconv.FToC(fahrenheit))

		// converts meter to inch and inch to meter
		meter := tempconv.Meter(t)
		inch := tempconv.Inch(t)
		fmt.Printf("%s = %s, %s = %s\n",
			meter, tempconv.MToI(meter), inch, tempconv.IToM(inch))

		// converts kilogram to pound and pound to kilogram
		kilogram := tempconv.Kilogram(t)
		pound := tempconv.Pound(t)
		fmt.Printf("%s = %s, %s = %s\n",
			kilogram, tempconv.KToP(kilogram), pound, tempconv.PToK(pound))
	}
}
