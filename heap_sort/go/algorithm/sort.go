package algorithm

func heapify(arr []int, N int, i int) {
	largest := i
	l := 2*i + 1
	r := 2*i + 2

	if l < N && arr[l] > arr[largest] {
		largest = l
	}

	if r < N && arr[r] > arr[largest] {
		largest = r
	}

	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		heapify(arr, N, largest)
	}
}

func Sort(arr []int) []int {
	N := len(arr)

	for i := (N / 2) - 1; i >= 0; i-- {
		heapify(arr, N, i)
	}

	for i := N - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		heapify(arr, i, 0)
	}

	return arr
}
