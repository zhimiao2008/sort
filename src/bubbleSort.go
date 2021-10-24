package src

// 冒泡排序
func BubbleSort(arr []int, size int) []int {
	for i := 0; i < size; i++ {
		for j := 0; j < size-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}
