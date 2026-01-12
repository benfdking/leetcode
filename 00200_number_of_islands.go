package leetcode

const (
	numIslandsNotLand = 48
	numIslandsLand    = 49
)

func numIslands(grid [][]byte) int {
	count := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == numIslandsLand {
				count++
				grid = markIsland(grid, p{i, j})
			}
		}
	}
	return count
}

func markIsland(grid [][]byte, position p) [][]byte {
	grid[position[0]][position[1]] = numIslandsNotLand
	positions := []p{
		{position[0], position[1] - 1},
		{position[0], position[1] + 1},
		{position[0] - 1, position[1]},
		{position[0] + 1, position[1]},
	}
	for _, newPosition := range positions {
		if !numIslandsIsOutOfImage(grid, newPosition) &&
			grid[newPosition[0]][newPosition[1]] == numIslandsLand {
			grid = markIsland(grid, newPosition)
		}
	}
	return grid
}

func numIslandsIsOutOfImage(img [][]byte, p p) bool {
	return p[0] < 0 || p[0] >= len(img) || p[1] < 0 || p[1] >= len(img[0])
}
