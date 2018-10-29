package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {

	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please input some arguments!")
		os.Exit(1)
	}

	var n float64
	var err error = errors.New("Not a numberic argument!")
	for i := 1; i < len(arguments); i++ {
		if arguments[i] == "STOP" {
			break
		}
		n, err = strconv.ParseFloat(arguments[i], 64)
		if err == nil {
			n += n
		} else {
			fmt.Println(err.Error())
		}
	}

	average := n / float64(len(arguments))
	if err == nil {
		fmt.Println("Sum is: ", n)
		fmt.Println("Average is: ", average)
	} else {
		fmt.Println("Some arguments are not numberic")
	}

}
