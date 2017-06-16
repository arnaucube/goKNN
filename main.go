package main

func main() {
	dataset := readDataset("datasets/poker/poker-hand-training-true.data", "\n", ",")
	//fmt.Println(dataset)
	inputs := readInputs("datasets/prova.data", "\n", ",")
	//fmt.Println(inputs)

	distances := euclideanDistance(dataset, inputs)
	getShortestDistance(distances)
}
