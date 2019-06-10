package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type inputData struct {
	Numbers string
	Coef    float64
}

func parseInputData(s string) []float64 {
	result := []float64{}
	s = strings.TrimSpace(s)
	for _, v := range strings.Split(s, " ") {
		if number, err := strconv.ParseFloat(v, 64); err == nil {
			result = append(result, number)
		}
	}
	return result
}

func prepareResponseResult(points []Point, result float64) []byte {
	buf := strconv.FormatFloat(result, 'f', 5, 64) + "\n"
	for _, v := range points {
		buf += strconv.FormatFloat(v.X, 'f', 5, 64) + " " + strconv.FormatFloat(v.Y, 'f', 5, 64) + "\n"
	}
	return []byte(buf)
}

// ExpHandler for exp. forecast
func ExpHandler(w http.ResponseWriter, r *http.Request) {
	var inData inputData
	json.NewDecoder(r.Body).Decode(&inData)

	if len(inData.Numbers) == 0 {
		w.WriteHeader(400)
		w.Write([]byte("Нет входных данных!"))
		return
	}

	w.WriteHeader(200)
	w.Write(prepareResponseResult(ExpSmooth(parseInputData(inData.Numbers), inData.Coef)))
}

// MiddleSliceHandler for middle slice forecast
func MiddleSliceHandler(w http.ResponseWriter, r *http.Request) {
	var inData inputData
	json.NewDecoder(r.Body).Decode(&inData)

	if len(inData.Numbers) == 0 {
		w.WriteHeader(400)
		w.Write([]byte("Нет входных данных!"))
		return
	}

	w.WriteHeader(200)
	w.Write(prepareResponseResult(MiddleSlice(parseInputData(inData.Numbers), inData.Coef)))
}

// RegressionHandler for regression forecast
func RegressionHandler(w http.ResponseWriter, r *http.Request) {
	var inData inputData
	json.NewDecoder(r.Body).Decode(&inData)

	if len(inData.Numbers) == 0 {
		w.WriteHeader(400)
		w.Write([]byte("Нет входных данных!"))
		return
	}

	w.WriteHeader(200)
	w.Write(prepareResponseResult(linearRegressionLSE(parseInputData(inData.Numbers))))
}






