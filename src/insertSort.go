package src

// 插入排序
func InsertSort(arr []int, size int) []int {
	for i := 1; i < size; i++ {
		val := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > val {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = val
	}
	return arr
}