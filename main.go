package main

import "fmt"

func main() {
	dataset := readDataFile("datasets/poker/poker-hand-training-true.data", "\n", ",")
	fmt.Println(dataset)
}
