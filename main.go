package main

import (
	"fmt"
	"math"
)

func main() {
	height, weight, err := readUserInput()
	if err != nil {
		panic(err)
	}

	result := calculateBMI(height, weight)
	fmt.Println(roundFloat(result, 1))
}

func readUserInput() (height, weight float64, err error) {
	fmt.Println("Введите массу тела: ")
	_, err = fmt.Scan(&weight)
	if err != nil {
		return 0, 0, fmt.Errorf("ошибка: %v", err)
	}

	fmt.Println("Введите ваш рост: ")
	_, err = fmt.Scan(&height)
	if err != nil {
		return 0, 0, fmt.Errorf("height: %v", err)
	}
	return height, weight, nil
}

func calculateBMI(height, weight float64) float64 {
	height = height / 100
	return weight / (height * height)
}

func roundFloat(f float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(f*ratio) / ratio
}
