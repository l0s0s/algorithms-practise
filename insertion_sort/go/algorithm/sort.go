package algorithm

func Sort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		j := i - 1
		for j >= 0 && arr[j] > arr[j+1] {
			arr[j], arr[j+1] = arr[j+1], arr[j]
			j--
		}
	}

	return arr
}
