package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// storing available option
var options = map[string]bool{
	"c": false,
	"m": false,
	"l": false,
	"w": false,
}

var files []string

func main() {
	// parse arguments
	// fmt.Println(len(os.Args))
	parseArguments()
	if len(files) > 0 {
		for _, filename := range files {
			content, err := ioutil.ReadFile(filename)
			if err != nil {
				log.Fatalf("Error reading file: %v", err)
			}
			// fmt.Println("Content: ", string(content))
			result := word_count(string(content))
			printWordCount(result)
		}
	} else {

	}

	// fmt.Println(options)
	// fmt.Println("Options: ", options)
	// fmt.Println("Files: ", files)
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

func check_file_exist(filePath string) bool {
	if _, err := os.Stat(filePath); err == nil {
		files = append(files, filePath)
		return true
	} else if os.IsNotExist(err) {
		fmt.Println("File does not exist.")
	} else {
		fmt.Println("Error:", err)
	}
	return false
}

func checkArg(arg string) bool {
	// if the length of the arg is less or equal to 2
	// then you have to check if it is a file or not
	// for now you have to throw an exception
	// otherwise you have to check if it is a valid option

	argLen := len(arg)

	switch {
	case argLen == 3:
		if arg[0] == '-' && arg[1] == '-' && (arg[2] == 'c' || arg[2] == 'm' || arg[2] == 'l' || arg[2] == 'w') {
			options[string(arg[2])] = true
			return true
		} else {
			fmt.Println(arg, " Is not A Good Option")
		}
	case argLen == 2:
		if arg[0] == '-' && (arg[1] == 'c' || arg[1] == 'm' || arg[1] == 'l' || arg[1] == 'w') {
			options[string(arg[1])] = true
			return true
		} else {
			fmt.Println(arg, " Is not A Good Option")
		}
	default:
		return (check_file_exist(arg))
	}
	return false
}

func word_count(data string) map[string]int {
	var result = make(map[string]int)

	if options["c"] {
		result["c"] = countBytes(data)
	}

	if options["m"] {
		result["m"] = countChars(data)
	}

	if options["l"] {
		result["l"] = countLines(data)
	}

	if options["w"] {
		result["w"] = countWords(data)
	}
	return result
}

func countBytes(data string) int {
	return len(data)
}

func countChars(data string) int {
	return len([]rune(data))
}

func countLines(data string) int {
	lines := strings.Split(data, "\n")
	return len(lines)
}

func countWords(data string) int {
	words := strings.Fields(data)
	return len(words)
}


func printWordCount(result map[string]int) {
	if val, ok := result["c"]; ok {
		fmt.Print(val, " ")
	}
	if val, ok := result["m"]; ok {
		fmt.Print(val, " ")
	}
	if val, ok := result["l"]; ok {
		fmt.Print(val, " ")
	}
	if val, ok := result["w"]; ok {
		fmt.Print(val)
	}
	fmt.Print("\n")
}