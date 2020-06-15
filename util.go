package main

import (
	"encoding/csv"
	"io"
	"os"
	"strconv"
	"strings"
)

// Note should load it into rows with maps from col index to 'header'

// Process the csv file and place the data in a map with the keys from the header line.
// If a single entry in the csv file cannot be parsed the read is aborted and an error returned
func ReadRecords(file *os.File) (map[string][]float64, error) {
	data := make(map[string][]float64)
	reader := csv.NewReader(file)
	reader.Comma = ','

	// First line should contain all the keys
	record, err := reader.Read()
	if err != nil {
		return nil, err
	}
	orderedKeys := []string{}
	for _, k := range record {
		key := strings.Trim(strings.TrimSpace(k), "#")
		orderedKeys = append(orderedKeys, key)
		data[key] = []float64{}
	}

	// Now place the data into the map, one slice of floats per key
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		for i, stringVal := range record {
			val, err := strconv.ParseFloat(strings.TrimSpace(stringVal), 64)
			if err != nil {
				return nil, err
			}
			data[orderedKeys[i]] = append(data[orderedKeys[i]], val)
		}
	}
	return data, nil
}


func Transpose(mat [][]float64) [][]float64 {
	xl := len(mat[0])
	yl := len(mat)
	result := make([][]float64, xl)
	for i := range result {
		result[i] = make([]float64, yl)
	}
	for i := 0; i < xl; i++ {
		for j := 0; j < yl; j++ {
			result[i][j] = mat[j][i]
		}
	}
	return result
}

func Add(s1 []float64, c float64) []float64 {
	result := make([]float64, len(s1))
	for i:=0;i<len(s1);i++ {
		result[i] = s1[i] + c
	}
	return result
}

func Sub(s1 []float64, c float64) []float64 {
	result := make([]float64, len(s1))
	for i:=0;i<len(s1);i++ {
		result[i] = s1[i] - c
	}
	return result
}

func Scale(s1 []float64, c float64) []float64 {
	result := make([]float64, len(s1))
	for i:=0;i<len(s1);i++ {
		result[i] = s1[i] * c
	}
	return result
}

func AddElem(s1,s2 []float64) []float64 {
	result := make([]float64, len(s1))
	for i:=0;i<len(s1);i++ {
		result[i] = s1[i] + s2[i]
	}
	return result
}

func SubElem(s1,s2 []float64) []float64 {
	result := make([]float64, len(s1))
	for i:=0;i<len(s1);i++ {
		result[i] = s1[i] - s2[i]
	}
	return result
}

func MulElem(s1,s2 []float64) []float64 {
	result := make([]float64, len(s1))
	for i:=0;i<len(s1);i++ {
		result[i] = s1[i] * s2[i]
	}
	return result
}

func DivElem(s1,s2 []float64) []float64 {
	result := make([]float64, len(s1))
	for i:=0;i<len(s1);i++ {
		result[i] = s1[i] / s2[i]
	}
	return result
}