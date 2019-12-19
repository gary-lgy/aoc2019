package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"strings"
)

func init() {
	solvers["8a"] = solve8a
	solvers["8b"] = solve8b
}

func getImageLayers(image []byte, width, height int) ([][][]byte, error) {
	pixels := width * height
	if len(image)%pixels != 0 {
		return nil, fmt.Errorf("invalid image input")
	}
	layersCount := len(image) / width / height
	layers := make([][][]byte, layersCount)
	for i := 0; i < layersCount; i++ {
		layer := make([][]byte, height)
		layers[i] = layer
		for j := 0; j < height; j++ {
			line := make([]byte, width)
			layer[j] = line
			for k := 0; k < width; k++ {
				line[k] = image[i*pixels+j*width+k]
			}
		}
	}
	return layers, nil
}

func solve8a(input io.Reader) (string, error) {
	width, height := 25, 6
	numbers, err := ioutil.ReadAll(input)
	if err != nil {
		return "", err
	}
	layers, err := getImageLayers(numbers, width, height)
	if err != nil {
		return "", err
	}
	minZeros, ones, twos := math.MaxInt32, 0, 0
	for _, layer := range layers {
		zeroCount, oneCount, twoCount := 0, 0, 0
		for i := 0; i < height; i++ {
			for j := 0; j < width; j++ {
				switch layer[i][j] {
				case '0':
					zeroCount++
				case '1':
					oneCount++
				case '2':
					twoCount++
				}
			}
		}
		if zeroCount < minZeros {
			minZeros, ones, twos = zeroCount, oneCount, twoCount
		}
	}
	return fmt.Sprint(ones * twos), nil
}

func displayImage(image [][]byte) {
	for _, row := range image {
		r := strings.NewReplacer("0", " ", "1", "■", "2", " ")
		fmt.Println(r.Replace(string(row)))
	}
}

func solve8b(input io.Reader) (string, error) {
	width, height := 25, 6
	numbers, err := ioutil.ReadAll(input)
	if err != nil {
		return "", err
	}
	layers, err := getImageLayers(numbers, width, height)
	if err != nil {
		return "", err
	}

	image := make([][]byte, height)
	for i := 0; i < height; i++ {
		row := make([]byte, width)
		image[i] = row
		for j := 0; j < width; j++ {
			image[i][j] = '2'
			for k := 0; k < len(layers) && image[i][j] == '2'; k++ {
				image[i][j] = layers[k][i][j]
			}
		}
	}
	displayImage(image)
	return "", nil
}
