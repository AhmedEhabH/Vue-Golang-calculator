package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type numbers struct {
	Num1 float64 `json:"num1"`
	Num2 float64 `json:"num2"`
}

type numsResponseData struct {
	Add float64 `json:"add"`
	Mul float64 `json:"mul"`
	Sub float64 `json:"sub"`
	Div float64 `json:"div"`
}

func process(numsdata numbers) numsResponseData {
	var numsRes numsResponseData
	numsRes.Add = numsdata.Num1 + numsdata.Num2
	numsRes.Mul = numsdata.Num1 * numsdata.Num2
	numsRes.Sub = numsdata.Num1 - numsdata.Num2
	numsRes.Div = numsdata.Num1 / numsdata.Num2
	fmt.Println("NUM RES")
	fmt.Println(numsRes)
	return numsRes
}

func calc(w http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
	// fmt.Println(decoder)

	var (
		numsData    numbers
		numsResData numsResponseData
	)

	decoder.Decode(&numsData)
	fmt.Println("NUM-1, NUM-2")
	fmt.Println(numsData)

	numsResData = process(numsData)
	fmt.Println("RES")
	fmt.Println(numsResData)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(numsResData); err != nil {
		panic(err)
	}
}

func main() {
	http.HandleFunc("/", calc)
	http.ListenAndServe(":8090", nil)
}
