package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args[1:]) != 1 {
		return
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal("error: ", err)
	}
	defer file.Close()

	data := []float64{}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		for _, item := range strings.Fields(line) {
			number, err := strconv.ParseFloat(item, 64)
			if err != nil {
				log.Fatal("error: ", err)
			}
			data = append(data, number)
		}
	}

	n := float64(len(data))

	var sumX, sumY, sumX2, sumY2, sumXY float64
	for i, y := range data {
		x := float64(i)
		sumX += x
		sumY += y
		sumX2 += x * x
		sumY2 += y * y
		sumXY += x * y
	}

	m := (n*sumXY - sumX*sumY) / (n*sumX2 - sumX*sumX)
	b := (sumY - m*sumX) / n

	pearson := (n*sumXY - sumX*sumY) / math.Sqrt((n*sumX2-sumX*sumX)*(n*sumY2-sumY*sumY))

	fmt.Printf("Linear Regression Line: y = %.6fx + %.6f\n", m, b)
	fmt.Printf("Pearson Correlation Coefficient: %.10f\n", pearson)
}
