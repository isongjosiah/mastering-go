package main

import "math/rand"

func random(min, max int) float64 {
	return float64(rand.Intn(max-min) + min)
}

func getCofactor(A [][]float64, temp [][]float64, p int, q int, n int) {
	i := 0
	j := 0

	for row := 0; row < n; row++ {
		for col := 0; col < n; col++ {
			if row != p && col != q {
				temp[i][j] = A[row][col]
				j++
				if j == n-1 {
					j = 0
					i++
				}
			}
		}
	}
}

func determinant(A [][]float64, n int) float64 {
	D := float64(0)
	if n == 1 {
		return A[0][0]
	}

	temp := createMatrix(n, n)
	sign := 1

	for f := 0; f < n; f++ {
		getCofactor(A, temp, 0, f, n)
		D += float64(sign) * A[0][f] * determinant(temp, n-1)
		sign = -sign
	}

	return D
}

func adjoint(A [][]float64) ([][]float64, error) {
	N := len(A)
	adj := createMatrix(N, N)
	if N == 1 {
		adj[0][0] = 1
		return adj, nil
	}
	sign := 1
	var temp = createMatrix(N, N)

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			g

		}
	}
}
