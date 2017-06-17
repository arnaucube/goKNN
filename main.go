package main

import (
	"fmt"
	"strconv"
)

func main() {
	dataset := readDataset("datasets/poker/poker-hand-training-true.data", "\n", ",")
	//fmt.Println(dataset)
	inputs := readInputs("datasets/prova.data", "\n", ",")
	//fmt.Println(inputs)

	var datasetDistances [][][]float64
	for k, input := range inputs {
		datasetDistances = euclideanDistance(dataset, input)
		r := getShortestDistance(datasetDistances)
		fmt.Println("iteration " + strconv.Itoa(k))
		fmt.Println(input)
		fmt.Println(r)
		fmt.Print("result: ")
		fmt.Println(r[1])
		fmt.Println("---")
	}
}
