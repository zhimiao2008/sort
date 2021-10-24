package src

// 快速排序
func QuickSort(arr []int, size int) []int {
	return _quickSort(arr, 0, size-1)
}

func _quickSort(arr []int, left int, right int) []int {
	if left >= right {
		return arr
	}
	key := arr[left]
	low := left
	high := right
	for left < right {
		for left < right && arr[right] >= key {
			right--
		}
		arr[left] = arr[right]

		for left < right && arr[left] <= key {
			left++
		}
		arr[right] = arr[left]
	}
	arr[right] = key
	_quickSort(arr, low, left-1)
	_quickSort(arr, left+1, high)
	return arr
}
