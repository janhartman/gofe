package gofe

import (
	"fmt"
	"os"
	"strconv"
)

func GetParams() (int, int, int) {

	if len(os.Args) < 4 {
		fmt.Println("No arguments passed")
		return 0, 0, 0
	}

	if os.Args[3] == "d" {
		if len(os.Args) < 6 {
			fmt.Println("not enough arguments")
			return 0, 0, 0
		}

		n, _ := strconv.Atoi(os.Args[4])
		l, _ := strconv.Atoi(os.Args[5])
		b, _ := strconv.Atoi(os.Args[6])

		return n, l, b
	} else if os.Args[3] == "a" {
		if len(os.Args) < 3 {
			fmt.Println("not enough arguments")
			return 0, 0, 0
		}

		a, _ := strconv.Atoi(os.Args[4])

		return a, 0, 0
	} else if os.Args[3] == "s" {
		if len(os.Args) < 4 {
			fmt.Println("not enough arguments")
			return 0, 0, 0
		}

		l, _ := strconv.Atoi(os.Args[4])
		b, _ := strconv.Atoi(os.Args[5])

		return l, b, 0
	} else {
		fmt.Println("Wrong scheme: ", os.Args[3])
		return 0, 0, 0
	}

}
