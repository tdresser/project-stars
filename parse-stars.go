package main

import (
	"encoding/csv"
	"fmt"
	"github.com/ajstarks/svgo"
	"os"
	"sort"
	"strconv"
)

type Star struct {
	Proper string
	Mag    float64
	X      float64
	Y      float64
	Z      float64
}

type Stars []Star

func (l Stars) Len() int {
	return len(l)
}

func (l Stars) Less(i, j int) bool {
	return l[i].Mag < l[j].Mag
}

func (l Stars) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func getStars() Stars {
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

	var stars Stars

	for i, row := range rawCSVdata {
		if i == 0 {
			continue
		}

		var proper = row[6]
		var mag, _ = strconv.ParseFloat(row[13], 64)
		var x, _ = strconv.ParseFloat(row[17], 64)
		var y, _ = strconv.ParseFloat(row[18], 64)
		var z, _ = strconv.ParseFloat(row[19], 64)

		if proper == "Sol" {
			continue
		}
		star := Star{
			Proper: proper,
			Mag:    mag,
			X:      x,
			Y:      y,
			Z:      z,
		}
		stars = append(stars, star)
	}

	sort.Sort(stars)
}

func main() {
	var stars Stars = getStars()

	for i, star := range stars {
		if i > 20 {
			break
		}
		fmt.Printf("%s: %f - (%f, %f, %f)\n",
			star.Proper, star.Mag, star.X, star.Y, star.Z)
	}
}
