package main

import "fmt"

func main() {
	dataset := readDataset("datasets/poker/poker-hand-training-true.data", "\n", ",")
	//fmt.Println(dataset)
	inputs := readInputs("datasets/prova.data", "\n", ",")
	//fmt.Println(inputs)

	var datasetDistances [][][]float64
	for _, input := range inputs {
		datasetDistances = euclideanDistance(dataset, input)
		r := getShortestDistance(datasetDistances)
		fmt.Println(input)
		fmt.Println(r)
		fmt.Println(r[1])
	}
}
