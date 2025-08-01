package _go

func rotate(matrix [][]int) {
	n := len(matrix)

	verticalLength := n / 2
	horizontalLength := n/2 + n%2

	for i := 0; i < verticalLength; i++ {
		for j := 0; j < horizontalLength; j++ {

			side := n - 1
			lu := p{i, j}
			ru := p{lu[1], side - lu[0]}
			rd := p{ru[1], side - ru[0]}
			ld := p{rd[1], side - rd[0]}

			temp := matrix[lu[0]][lu[1]]
			matrix[lu[0]][lu[1]] = matrix[ld[0]][ld[1]]
			matrix[ld[0]][ld[1]] = matrix[rd[0]][rd[1]]
			matrix[rd[0]][rd[1]] = matrix[ru[0]][ru[1]]
			matrix[ru[0]][ru[1]] = temp
		}
	}
}
