package main

import (
	"encoding/csv"
	"fmt"
	"github.com/ajstarks/svgo"
	"math"
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
	var stars Stars

	csvfile, err := os.Open("HYG-Database/hygdata_v3.csv")

	if err != nil {
		fmt.Println(err)
		return stars
	}

	defer csvfile.Close()

	reader := csv.NewReader(csvfile)
	rawCSVdata, err := reader.ReadAll()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

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
	return stars;
}

const scale = 10

func draw(stars Stars) {
	f, _ := os.Create("stars.svg")
	defer f.Close()

	width := 500
	height := 500
	canvas := svg.New(f)
	canvas.Start(width, height)
	defer canvas.End()

	for i, star := range stars {
		if i > 10000 {
			break
		}

		var x = star.X / star.Z * 100
		var y = star.Y / star.Z * 100
		var mag = math.Pow((5 - star.Mag), 0.3) * 0.8

		canvas.Circle(
			int(x * scale), 
			int(y * scale), 
			int(mag * scale))

//		fmt.Printf("%s: %f - (%f, %f, %f)\n",
//			star.Proper, star.Mag, star.X, star.Y, star.Z)
	}
}

func main() {
	var stars Stars = getStars()
	draw(stars)
}
