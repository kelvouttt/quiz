package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	var input int

	f, err := os.Open("problems.csv")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	csvReader := csv.NewReader(f)
	for {
		rec, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%v\n", rec[0])
		fmt.Print("Please enter the answer: ")
		fmt.Scan(&input)
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
