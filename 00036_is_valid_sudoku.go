package leetcode

func isValidSudoku(board [][]byte) bool {
	// horizontals
	for i := 0; i < 9; i++ {
		var vs [9]byte
		for j := range vs {
			vs[j] = board[i][j]
		}
		if !isValidLine(vs) {
			return false
		}
	}

	// verticals
	for i := 0; i < 9; i++ {
		var outs [9]byte
		for j := 0; j < 9; j++ {
			outs[j] = board[j][i]
		}
		if !isValidLine(outs) {
			return false
		}
	}

	// squares
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {

			var vs [9]byte
			for k := 0; k < 3; k++ {
				for l := 0; l < 3; l++ {
					vs[k*3+l] = board[j*3+k][i*3+l]
				}
			}
			if !isValidLine(vs) {
				return false
			}

		}
	}

	return true
}

func isValidLine(values [9]byte) bool {
	vs := makeCount(values)
	for _, v := range vs {
		if v > 1 {
			return false
		}
	}
	return true
}

func makeCount(values [9]byte) [9]int {
	var out [9]int
	for _, v := range values {
		position, ok := valuesMap[v]
		if ok {
			out[position]++
		}
	}
	return out
}

func returnMap() map[byte]int {
	m := make(map[byte]int, len(values))
	for i, v := range values {
		m[v] = i
	}
	return m
}

var valuesMap = returnMap()

var values = []byte{'1', '2', '3', '4', '5', '6', '7', '8', '9'}
