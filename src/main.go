package main

import (
	"fmt"
	"os"
	"log"
	"encoding/csv"
	"bufio"
	"io"
	"strconv"
	"path/filepath"
)


type IrisData struct{
	sepalLen float64
	sepalWid float64
	petalLen float64
	petalWid float64
	class string
}

func main() {
	// Testing comments
	absPath, _ := filepath.Abs("./datasets/iris.data")
	irisFile, err := os.Open(absPath)
	if err != nil {
		log.Fatal(err)
	}
	iris_reader := csv.NewReader(bufio.NewReader(irisFile))

	var iris []IrisData

	for {
		line, err := iris_reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		sL, _ := strconv.ParseFloat(line[0], 64)
		sW, _ := strconv.ParseFloat(line[1], 64)
		pL, _ := strconv.ParseFloat(line[2], 64)
		pW, _ := strconv.ParseFloat(line[3], 64)

		iris = append(iris, IrisData{
			sepalLen: sL,
			sepalWid: sW,
			petalLen: pL,
			petalWid: pW,
			class:    line[4],
		}, )
	}

	fmt.Printf("I just read in %d data points\n", len(iris))
	fmt.Println("The first record is:")
	fmt.Printf("%+v\n", iris[0])
}

