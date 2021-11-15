package main

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/hash-sha512/octo-couscous/helper"
)

func main() {

	data := helper.ReadFile("input_3.csv")

	var header = [][]string{append(data[0][:1], data[0][30:32]...)}
	records := data[1:]
	output := make([][]string, len(records))

	for i := range records {
		output[i] = append(records[i][:1], records[i][30:32]...)
	}

	fmt.Println("running sort algorithm on the records")

	sort.Slice(output, func(i, j int) bool {
		//index - 1->yellow, 2->red
		a, _ := strconv.Atoi(output[i][2])
		b, _ := strconv.Atoi(output[j][2])
		c, _ := strconv.Atoi(output[i][1])
		d, _ := strconv.Atoi(output[j][1])

		if a != b {
			return a > b //based on red cards
		}
		return c > d //on equal no. of red cards, considers yellow cards
	})

	output = append(header, output...)

	helper.WriteFile("output_3.csv", output)

	fmt.Println("fin")
}
