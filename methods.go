package main

import (
	"fmt"
	"math"
)

import "github.com/VividCortex/ewma"

// Point used for linear regression
type Point struct {
	X float64
	Y float64
}

//ExpSmooth - метод экспоненциального сглаживания
func ExpSmooth(inputRow []float64, alpha float64) ([]Point, float64) {
	outputRow := []float64{}
	resultRow := []float64{}
	resultPoints := make([]Point, len(inputRow))

	var result float64
	var s0 float64 //Сумма первых нескольких элементов ряда, деленное на их количество

	if len(inputRow) >= 3 {
		s0 = (inputRow[0] + inputRow[1] + inputRow[2]) / 3
	} else {
		s0 = inputRow[0]
	}

	fmt.Println("\nInput row: ", inputRow, "\n")

	var t float64
	for i := 0; i < len(inputRow); i++ {
		if i == 0 {
			t = (1-alpha)*inputRow[i] + alpha*s0
			fmt.Println("Iteration", i, ": (1 -", alpha, ") *", inputRow[i], "+", alpha, "*", s0, "=", t)
		} else {
			t = (1-alpha)*inputRow[i] + alpha*outputRow[i-1]
			fmt.Println("Iteration", i, ": (1 -", alpha, ") *", inputRow[i], "+", alpha, "*", outputRow[i-1], "=", t)
		}
		outputRow = append(outputRow, math.Round(t/0.05)*0.05)
	}

	fmt.Println("\nOutput row: ", outputRow)
	fmt.Println("\n\n")

	for i := 0; i < len(inputRow); i++ {
		resultRow = append(resultRow, math.Pow(inputRow[i]-outputRow[i], 2))
		fmt.Println(inputRow[i], " => ", outputRow[i])
		result += resultRow[i]
		resultPoints[i] = Point{float64(i), resultRow[i]}
	}

	fmt.Println("Result: ", result)
	return resultPoints, result
}

//MiddleSlice - метод скользящей средней
func MiddleSlice(inputRow []float64, smoothK float64) ([]Point, float64) {
	a := ewma.NewMovingAverage(smoothK)
	for _, f := range inputRow {
		a.Add(f)
	}

	return nil, a.Value()
}
//Линен регр
func linearRegressionLSE(inputRow []float64) ([]Point, float64) {
	q := len(inputRow)
	series := make([]Point, q, q)
	resultPoints := make([]Point, len(inputRow))

	for i, v := range inputRow {
		series[i] = Point{float64(i), v}
	}

	p := float64(q)
	sumX, sumY, sumXx, sumXy := 0.0, 0.0, 0.0, 0.0

	for _, p := range series {
		sumX += p.X
		sumY += p.Y
		sumXx += p.X * p.X
		sumXy += p.X * p.Y
	}

	m := (p*sumXy - sumX*sumY) / (p*sumXx - sumX*sumX)
	b := (sumY / p) - (m * sumX / p)

	for i, p := range series {
		resultPoints[i] = Point{p.X, (p.X*m + b)}
	}

	return resultPoints, resultPoints[len(resultPoints)-1].Y
}
