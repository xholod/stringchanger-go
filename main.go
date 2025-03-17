package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 4 {
		log.Printf("\nThe application requires 3 parameters:\n\t - directory\n\t - source string\n\t - target string\n")
		log.Printf("\nExample: stringchanger-go /home/username/customDir sourceText targetText\n")
		panic("The application has terminated.")
	}

	directory := os.Args[1]
	source_string := os.Args[2]
	target_string := os.Args[3]

	log.Printf("\nTarget directory:\t\t%s\nSource string:\t\t\t%s\nTarget string:\t\t\t%s\n\n", directory, source_string, target_string)

	dirEntries, err := os.ReadDir(directory)
	if err != nil {
		panic(err)
	}

Loop:
	for _, dirEntry := range dirEntries {
		if !dirEntry.IsDir() {
			filename := directory + "/" + dirEntry.Name()

			fmt.Printf("\nDo you want to change the file (%s)?\n", filename)
			fmt.Print("Your answer (y/n): ")
			reader := bufio.NewReader(os.Stdin)
			answer, _, readRuneErr := reader.ReadRune()
			if readRuneErr != nil {
				fmt.Println(readRuneErr)
			}

			if answer == 'n' {
				log.Printf("File %s was skipped\n", filename)
				continue Loop
			} else if answer == 'y' {

				source_data := readFile(filename)
				target_data := strings.ReplaceAll(source_data, source_string, target_string)
				writeErr := os.WriteFile(filename, []byte(target_data), 0644)
				if writeErr != nil {
					panic(writeErr)
				}
				log.Printf("File %s is OK\n", filename)

			} else {
				log.Printf("File %s was skipped\n", filename)
			}
		}
	}

	fmt.Println()
	log.Println("There are no files to modify.")
	log.Println("The application has terminated.")

}

func readFile(filepath string) string {
	data, err := os.ReadFile(filepath)
	if err != nil {
		panic(err)
	}
	dataResult := string(data)

	return dataResult
}
