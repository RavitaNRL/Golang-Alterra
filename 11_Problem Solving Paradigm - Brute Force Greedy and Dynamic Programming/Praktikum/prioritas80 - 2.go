//SOAL
// Diberi bilangan bulat numRows, kembalikan numRows pertama dari segitiga Pascal. Dalam segitiga Pascal, setiap angka adalah jumlah dari dua angka tepat di atasnya seperti yang ditunjukkan:
// Input: numRows = 5
// Output:

package main

import "fmt"

func generatePascal(numRows int) [][]int {
	if numRows <= 0 {
		return nil
	}

	result := make([][]int, numRows)

	for n := range result {
		result[n] = make([]int, n+1)
		result[n][0], result[n][n] = 1, 1

		for j := 1; j < n; j++ {
			result[n][j] = result[n-1][j-1] + result[n-1][j]
		}
	}

	return result
}

func main() {
	fmt.Println(generatePascal(2))
	fmt.Println(generatePascal(4))
	fmt.Println(generatePascal(5))

}
