package src

// 选择排序
func SelectSort(arr []int, size int) []int {
	for i := 0; i < size; i++ {
		min := i
		for j := i; j < size; j++ {
			if arr[min] > arr[j] {
				min = j
			}
		}
		arr[min], arr[i] = arr[i], arr[min]
	}
	return arr
}