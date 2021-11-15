package helper

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func WriteFile(filename string, data [][]string) {
	fmt.Printf("writing data to file - %s \n", filename)

	f, err := os.Create(filename)
	defer f.Close()

	if err != nil {
		log.Fatalln("failed to open file:", err)
	}

	w := csv.NewWriter(f)
	w.WriteAll(data) // calls Flush internally

	if err := w.Error(); err != nil {
		log.Fatalln("error writing csv:", err)
	}

}

func ReadFile(filename string) [][]string {
	fmt.Printf("reading data from file - %s \n", filename)

	f, err := os.Open(filename)
	defer f.Close()

	if err != nil {
		log.Fatal(err)
	}

	r := csv.NewReader(f)
	data, err := r.ReadAll()

	if err != nil {
		log.Fatal(err)
	}

	return data
}
