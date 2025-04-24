package easy

func flipAndInvertImage(image [][]int) [][]int {
	for i := range image {
		for j, k := 0, len(image[i])-1; j <= k; j, k = j+1, k-1 {
			image[i][j], image[i][k] = image[i][k]^1, image[i][j]^1
		}
	}

	return image
}
