package src

// 归并排序
func MergeSort(arr []int, size int) []int {
	if size == 1 {
		return arr
	}

	mid := size / 2

	left := arr[:mid]
	right := arr[mid:]

	ll := MergeSort(left, len(left))
	rl := MergeSort(right, len(right))
	return mergeArr(ll, rl)
}

func mergeArr(ll, rl []int) []int {
	var result []int
	mi, mj := len(ll), len(rl)
	count := mi + mj
	result = make([]int, count)
	var i, j int
	for i < mi && j < mj {
		if ll[i] < rl[j] {
			result[i+j] = ll[i]
			i++
		} else {
			result[i+j] = rl[j]
			j++
		}
	}
	for i < mi {
		result[i+j] = ll[i]
		i++
	}
	for j < mj {
		result[i+j] = rl[j]
		j++
	}

	return result
}
