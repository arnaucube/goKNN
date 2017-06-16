package main

import (
	"math"
)

func euclideanDistance(dataset [][][]int, inputLine []int) [][][]float64 {
	var distances [][][]float64

	for _, datasetLine := range dataset {
		datasetLineFloat := arrayIntToFloat64(datasetLine[0])
		resultDatasetLineFloat := arrayIntToFloat64(datasetLine[1])
		inputLineFloat := arrayIntToFloat64(inputLine)
		r := calcDistances(inputLineFloat, datasetLineFloat)
		distance := []float64{r}
		var newDatasetLine [][]float64
		newDatasetLine = append(newDatasetLine, datasetLineFloat)
		newDatasetLine = append(newDatasetLine, resultDatasetLineFloat)
		newDatasetLine = append(newDatasetLine, distance)
		distances = append(distances, newDatasetLine)
		//distances contains: [[datasetLine], [datasetLineResult], [distanceWithInput]]
	}
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

func getShortestDistance(dist [][][]float64) [][]float64 {
	min := dist[0][2]
	max := dist[0][2]
	var datasetNearResult [][]float64
	for i := 0; i < len(dist); i++ {
		//fmt.Println(dist[i][1])
		if dist[i][2][0] > max[0] {
			max = dist[i][0]
		}
		/*fmt.Println(min[0])
		fmt.Println(dist[i][1][0])*/
		if dist[i][2][0] < min[0] {
			min = dist[i][0]
			datasetNearResult = dist[i]
		}
		//fmt.Println(dist[i][1])
	}
	/*fmt.Println(min)
	fmt.Println(max)*/
	return datasetNearResult
}
