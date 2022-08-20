package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

var score int
var incorrect int
var input string

func main() {
	file, _ := openFile()
	defer file.Close()

	csvReader := csv.NewReader(file)

	for {
		rec, err := csvReader.Read()
		if err == io.EOF {
			fmt.Println("Thanks for taking the quiz :)")
			fmt.Printf("Your total score is %v with %v incorrect out of %v questions\n", score, incorrect, score+incorrect)
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%v\n", rec[0])
		fmt.Print("Answer: ")
		fmt.Scan(&input)

		if input == rec[1] {
			score++
			fmt.Printf("Your score is %v\n", score)
		} else {
			incorrect++
		}
	}
}

/*
func userInput() string {
	var input string

	fmt.Println("Answer:")
	fmt.Scan(&input)

	return input
}
*/

func openFile() (*os.File, error) {
	f, err := os.Open("problems.csv")
	if err != nil {
		log.Fatal(err)
	}
	return f, err
}
