package main

import (
	"fmt"
	"math/rand"
)

const N = 1000

func multiplyMatricesSequential(matrixA [N][N]float64, matrixB [N][N]float64) [N][N]float64 {
	var result [N][N]float64;
	for y = 0; y < N; y++ {
		for x = 0; x < N; x++ {
			sum = 0
			for i = 0; i < N; i++ {
				sum += matrixA[y][i] * matrixB[i][x]
			}
			result[y][x] = sum
		}
	}
	return result
}

func multiplyMatricesConcurrent(matrixA [N][N]float64, matrixB [N][N]float64) [N][N]float64 {
	var result [N][N]float64;
	/*for y = 0; y < N; y++ {
		for x = 0; x < N; x++ {
			sum = 0
			for i = 0; i < N; i++ {
				sum += matrixA[y][i] * matrixB[i][x]
			}
			result[y][x] = sum
		}
	}*/
	return result
}

func main() {
	var matrixA [N][N]float64;
	var matrixB [N][N]float64;

	for _, matrix := range [][N][N]float64 {matrixA, matrixB} {
		for _, row := range matrix {
			for x := range row {
				row[x] = rand.Float64()
			}
		}
	}

	fmt.Println(multiplyMatricesSequential(matrixA, matrixB))
	fmt.Println(multiplyMatricesConcurrent(matrixA, matrixB))
}
