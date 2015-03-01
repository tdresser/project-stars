package main

import (
	"fmt"
	"encoding/csv"	
	"os"
	"strconv"
)

func main() {
	csvfile, err := os.Open("HYG-Database/hygdata_v3.csv")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer csvfile.Close()

	reader := csv.NewReader(csvfile)
	rawCSVdata, err := reader.ReadAll()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, row := range rawCSVdata {
		var proper = row[6]
		var mag, _ = strconv.ParseFloat(row[13], 64)
		var x, _ = strconv.ParseFloat(row[17], 64)
		var y, _ = strconv.ParseFloat(row[18], 64)
		var z, _ = strconv.ParseFloat(row[19], 64)
		fmt.Printf("%s: %f - (%f, %f, %f)\n", proper, mag, x, y, z)
	}
}
