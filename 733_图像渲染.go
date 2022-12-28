package main

// 岛屿问题的变形
func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	old := image[sr][sc]
	fill(image, sr, sc, color, old)
	return image

}

func fill(image [][]int, sr, sc int, newColor, old int) {
	if !inArea(image, sr, sc) {
		return
	}

	if image[sr][sc] == newColor {
		return
	}

	if image[sr][sc] != old {
		return
	}

	image[sr][sc] = newColor
	// 填充颜色
	fill(image, sr+1, sc, newColor, old)
	fill(image, sr-1, sc, newColor, old)
	fill(image, sr, sc+1, newColor, old)
	fill(image, sr, sc-1, newColor, old)

}

func inArea(image [][]int, r, c int) bool {
	return r >= 0 && r < len(image) && c >= 0 && c < len(image[0])
}
