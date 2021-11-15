package main

import (
	"fmt"
	"math"
	"strconv"

	"github.com/hash-sha512/octo-couscous/helper"
)

func main() {

	records := helper.ReadFile("input_2.csv")

	hashmap := make(map[string]map[string]int)

	for _, record := range records {

		age, _ := strconv.Atoi(record[1])
		_, exists := hashmap[record[3]]

		if exists {
			hashmap[record[3]]["minimum"] = int(math.Min(float64(hashmap[record[3]]["minimum"]), float64(age)))
			hashmap[record[3]]["maximum"] = int(math.Max(float64(hashmap[record[3]]["maximum"]), float64(age)))
		} else {
			hashmap[record[3]] = map[string]int{"minimum": age, "maximum": age}
		}
	}

	var output [][]string

	for key := range hashmap {
		row := []string{key, strconv.Itoa(hashmap[key]["minimum"]), strconv.Itoa(hashmap[key]["maximum"])}
		output = append(output, row)
	}

	header := [][]string{{"occupation", "minimum", "maximum"}}
	output = append(header, output...)

	helper.WriteFile("output_2.csv", output)

	fmt.Println("fin")

}
