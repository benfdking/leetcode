package leetcode

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	return floodFillWithP(image, p{sr, sc}, image[sr][sc], color)
}

type p [2]int

func floodFillWithP(image [][]int, position p, originalNumber int, color int) [][]int {
	image[position[0]][position[1]] = color
	positions := []p{
		{position[0], position[1] - 1},
		{position[0], position[1] + 1},
		{position[0] - 1, position[1]},
		{position[0] + 1, position[1]},
	}
	for _, newPosition := range positions {
		if !isOutOfImage(image, newPosition) &&
			image[newPosition[0]][newPosition[1]] == originalNumber &&
			image[newPosition[0]][newPosition[1]] != color {
			image = floodFillWithP(image, newPosition, originalNumber, color)
		}
	}
	return image
}

func isOutOfImage(img [][]int, p p) bool {
	return p[0] < 0 || p[0] >= len(img) || p[1] < 0 || p[1] >= len(img[0])
}
