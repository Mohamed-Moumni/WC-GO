package main

import (
	"fmt"
	"os"
)

// list of options
var option_chars = []rune{'c', 'm', 'l', 'w'}

// storing available option
var options = map[string]bool {
	"c":false,
	"m":false,
	"l":false,
	"w":false,
}

func main() {
	// parse arguments
	parseArguments()
	fmt.Println(len(os.Args))
	// fmt.Println(options)
}

func parseArguments() {
	if len(os.Args) == 1 {
		return
	}

	for i := 1; i < len(os.Args); i++ {
		// check argument
		if !checkArg(os.Args[i]) {
			os.Exit(1)
		}
	}
}

func check_file_exist(filePath string) bool{
	if _, err := os.Stat(filePath); err == nil {
		return true
    } else if os.IsNotExist(err) {
        fmt.Println("File does not exist.")
    } else {
        fmt.Println("Error:", err)
    }
	return false
}

func checkArg (arg string) bool{
	// if the lenght of the arg is less or equal to 2
	// then you have to check if it is a file or not
	// for now you have to throw an exception
	// otherwise you have to check if it is a valid option

	argLen := len(arg)

	switch {
	case argLen == 3:
		if arg[0] == '-' && arg[1] == '-' && (arg[2] == 'c' || arg[2] == 'm' || arg[2] == 'l' || arg[2] == 'w') {
			return true
		} else {
			fmt.Println(arg, " Is not A Good Option")
		}
	case argLen == 2:
		if arg[0] == '-' && (arg[1] == 'c' || arg[1] == 'm' || arg[1] == 'l' || arg[1] == 'w') {
			return true
		} else {
			fmt.Println(arg, " Is not A Good Option")
		}
	default:
		check_file_exist(arg)
	}
	return false
}
