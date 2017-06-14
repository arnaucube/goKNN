package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func readDataFile(path string, lineSeparation string, valueSeparation string) [][][]int {
	var dataset [][][]int
	b, err := ioutil.ReadFile(path)
	check(err)
	str := string(b)
	str = strings.Replace(str, "\r", "", -1)
	lines := strings.Split(str, lineSeparation)
	for _, v1 := range lines {
		params := strings.Split(v1, valueSeparation)
		var datasetLine [][]int
		var datasetLineEntry []int
		var lastLineValue []int
		for k2, v2 := range params {
			value, err := strconv.Atoi(v2)
			check(err)
			if k2 < len(params)-1 {
				datasetLineEntry = append(datasetLineEntry, value)
			} else {
				lastLineValue = append(lastLineValue, value)
			}
		}
		datasetLine = append(datasetLine, datasetLineEntry)
		datasetLine = append(datasetLine, lastLineValue)
		dataset = append(dataset, datasetLine)
	}
	return dataset
}
