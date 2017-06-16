package main

import (
	"fmt"
	"math"
)

func euclideanDistance(dataset [][][]int, inputs [][]int) [][][]float64 {
	var distances [][][]float64

	for _, inputLine := range inputs {
		for _, datasetLine := range dataset {
			datasetLineFloat := arrayIntToFloat64(datasetLine[0])
			inputLineFloat := arrayIntToFloat64(inputLine)
			r := calcDistances(inputLineFloat, datasetLineFloat)
			//fmt.Println(r)
			rArr := []float64{r}
			var newDatasetLine [][]float64
			newDatasetLine = append(newDatasetLine, datasetLineFloat)
			newDatasetLine = append(newDatasetLine, rArr)
			distances = append(distances, newDatasetLine)
			//passar tot els int de values a float64
		}
	}
	//fmt.Println(train)
	return distances
}

func calcDistances(in []float64, ds []float64) float64 {
	var r float64
	r = 0
	for i := 0; i < len(in); i++ {
		r = r + (in[i] * ds[i])
	}
	r = math.Sqrt(r)
	return r
}

func getShortestDistance(dist [][][]float64) {
	min := dist[0][0]
	max := dist[0][0]
	for i := 0; i < len(dist); i++ {
		if dist[i][1][0] > max[0] {
			max = dist[i][0]
		}
		if dist[i][0][0] < min[0] {
			min = dist[i][0]
		}
		//fmt.Println(dist[i][1])
	}
	fmt.Println(min)
	fmt.Println(max)
}
