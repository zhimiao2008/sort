package src

// 堆排序
func HeapSort(arr []int, size int) []int {
	buildHeap(arr, size)
	//fmt.Println("heap arr", arr)
	heapLength := size
	for i := size - 1; i >= 0; i-- {
		// 交换堆顶和最后一个元素，即每次将剩余元素中的最大者放到最后面
		arr[i], arr[0] = arr[0], arr[i]

		heapLength--
		//重新调整堆顶节点成为大顶堆
		adjustHeap(arr, 0, heapLength)
		//fmt.Println("heap new arr", arr, arr[i])
	}
	return arr
}

func buildHeap(arr []int, size int) {
	// 从最后一个非叶 结点 到 第1个结点 依次进行向上调整
	for i := size/2; i >= 1; i-- {
		adjustHeap(arr, i, size)
	}
}

// 堆实际上是一棵完全二叉树，其任何一非叶节点满足性质：
//  Key[i]<=key[2i+1]&&Key[i]<=key[2i+2] 或者 Key[i]>=Key[2i+1]&&key>=key[2i+2]
// 即任何一非叶节点的关键字不大于或者不小于其左右孩子节点的关键字。
//
// 堆分为大顶堆和小顶堆，满足
// - 大顶堆： Key[i]>=Key[2i+1]&&key>=key[2i+2]
// - 小顶堆： Key[i]<=key[2i+1]&&Key[i]<=key[2i+2]
// 由上述性质可知大顶堆的堆顶的关键字肯定是所有关键字中最大的，小顶堆的堆顶的关键字是所有关键字中最小的。
func adjustHeap(heap []int, i int, length int) {
	// 左子节点2i+1, 右子节点2i+2, 父节点(i-1)/2

	// 左子节点2i+1
	lchild := 2*i + 1
	// 右子节点2i+2
	rchild := 2*i + 2
	// 临时变量
	max := 0

	// 如果i不是叶节点就不用进行调整
	if i <= length/2 {
		//首先判断他和他左儿子的关系，并用 max 记录值较大的结点编号
		if lchild < length && heap[lchild] > heap[i] {
			max = lchild
		} else {
			max = i
		}

		// 如果他有右儿子的情况下，再对右儿子进行讨论
		if rchild < length && heap[rchild] > heap[max] {
			max = rchild
		}

		// 如果发现最大的结点编号不是自己，说明子结点中有比父结点更大的
		if max != i {
			// 交换它们，注意swap函数需要自己来写
			heap[i], heap[max] = heap[max], heap[i]
			//需要继续比较起父节点,  避免调整之后以max为父节点的子树不是堆
			adjustHeap(heap, max, length)
		}
	}
}
