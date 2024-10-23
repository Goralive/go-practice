package main

type cost struct {
	day   int
	value float64
}

func sum(nums ...int) int {
	var returnValue int = 0
	for i := 0; i < len(nums); i++ {
		returnValue += nums[i]
	}
	return returnValue
}

func getCostsByDay(costs []cost) []float64 {
	if len(costs) == 0 {
		return []float64{}
	}
	maxDays := 0

	for _, c := range costs {
		if c.day > maxDays {
			maxDays = c.day
		}
	}

	result := make([]float64, maxDays+1)

	for _, cost := range costs {
		result[cost.day] += cost.value
	}

	return result
}

func createMatrix(rows, cols int) [][]int {
	matrix := make([][]int, rows)
	for i := 0; i < rows; i++ {
		matrix[i] = make([]int, cols)
		for j := 0; j < cols; j++ {
			matrix[i][j] = i * j
		}
	}
	return matrix
}

func indexOfFirstBadWord(msg []string, badWords []string) int {
	for i, word := range msg {
		for _, badWord := range badWords {
			if word == badWord {
				return i
			}
		}
	}
	return -1
}
