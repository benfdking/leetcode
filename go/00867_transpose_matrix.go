package _go

func transpose(matrix [][]int) [][]int {
	if len(matrix) == 0 {
		return [][]int{}
	}
	out := make([][]int, len(matrix[0]))
	for i := range out {
		out[i] = make([]int, len(matrix))
	}

	for i := range matrix {
		for j := range matrix[i] {
			out[j][i] = matrix[i][j]
		}
	}

	return out
}
