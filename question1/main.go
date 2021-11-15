package main

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/hash-sha512/octo-couscous/helper"
)

func aggregatePerDecade(data [][]string, c chan []string) {
	var aggregate = make([]string, 12)
	aggregate[0] = data[0][0]           //decade starting year
	aggregate[1] = data[len(data)-1][1] //population on the closing decade year
	for row := range data {
		for field := 2; field < 12; field++ {
			//if not for strict type conversions and strconv return of (value, err) ... this woulda just been a += c
			sum, _ := strconv.Atoi(data[row][field])
			currentValue, _ := strconv.Atoi(aggregate[field])
			aggregate[field] = strconv.Itoa(sum + currentValue)
		}
	}
	c <- aggregate
	defer close(c)
}

func main() {

	data := helper.ReadFile("input_1.csv")

	header := data[0:1]
	records := data[1:]

	//TODO array of channels(dynamic length based on len(data)/10) vs just a single multiplexed buffered channel
	c1 := make(chan []string)
	c2 := make(chan []string)
	c3 := make(chan []string)
	c4 := make(chan []string)
	c5 := make(chan []string)
	c6 := make(chan []string)

	for i := 0; i < len(records); i = i + 10 {
		//? playing devil's advocate - do eval("'c' + i") and skip the multiple if ...
		if i == 0 {
			go aggregatePerDecade(records[i:i+10], c1)
		} else if i == 10 {
			go aggregatePerDecade(records[i:i+10], c2)
		} else if i == 20 {
			go aggregatePerDecade(records[i:i+10], c3)
		} else if i == 30 {
			go aggregatePerDecade(records[i:i+10], c4)
		} else if i == 40 {
			go aggregatePerDecade(records[i:i+10], c5)
		} else if i == 50 {
			go aggregatePerDecade(records[i:], c6)
		}
	}

	fmt.Println("all goroutines spawned")

	var output [][]string

	for iter := 0; iter < 6; iter++ {
		select {
		case msg1 := <-c1:
			output = append(output, msg1[:])
		case msg2 := <-c2:
			output = append(output, msg2[:])
		case msg3 := <-c3:
			output = append(output, msg3[:])
		case msg4 := <-c4:
			output = append(output, msg4[:])
		case msg5 := <-c5:
			output = append(output, msg5[:])
		case msg6 := <-c6:
			output = append(output, msg6[:])
		}
	}

	fmt.Println("sorting because the execution of goroutines(threads) is not sequential")

	sort.Slice(output, func(i, j int) bool {
		a, _ := strconv.Atoi(output[i][0])
		b, _ := strconv.Atoi(output[j][0])
		return a < b
	})

	output = append(header, output...) //prepend the header row
	helper.WriteFile("output_1.csv", output)

	fmt.Println("fin")

}
