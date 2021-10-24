package main

import (
	"fmt"
	"math/rand"
	"sort/src"
	"time"
)

var (
	oriArr = []int{12, 23, 43, 54, 1, 42, 34, 54, 65, 76, 87, 11, 32}
)

func main() {
	//trace.Start(os.Stderr)
	//defer trace.Stop()

	//++++++++++++++++++++++++++++++++++
	// 稳定排序
	//++++++++++++++++++++++++++++++++++

	// BubbleSort
	fmt.Println("origin arr", oriArr)
	fmt.Println("sorted arr", src.BubbleSort(oriArr, len(oriArr)), "BubbleSort")

	shuffle(oriArr)
	fmt.Println("origin arr", oriArr)
	fmt.Println("sorted arr", src.InsertSort(oriArr, len(oriArr)), "InsertSort")

	shuffle(oriArr)
	fmt.Println("origin arr", oriArr)
	fmt.Println("sorted arr", src.MergeSort(oriArr, len(oriArr)), "MergeSort")

	//++++++++++++++++++++++++++++++++++
	// 不稳定排序
	//++++++++++++++++++++++++++++++++++

	shuffle(oriArr)
	fmt.Println("origin arr", oriArr)
	fmt.Println("sorted arr", src.QuickSort(oriArr, len(oriArr)), "QuickSort")

	shuffle(oriArr)
	fmt.Println("origin arr", oriArr)
	fmt.Println("sorted arr", src.SelectSort(oriArr, len(oriArr)), "SelectSort")

	shuffle(oriArr)
	fmt.Println("origin arr", oriArr)
	fmt.Println("sorted arr", src.HeapSort(oriArr, len(oriArr)), "HeapSort")

	fmt.Println("over")
}

// 随机打乱数组
func shuffle(slice []int) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for len(slice) > 0 {
		n := len(slice)
		randIndex := r.Intn(n)
		slice[n-1], slice[randIndex] = slice[randIndex], slice[n-1]
		slice = slice[:n-1]
	}
}
