package sorting

import (
	"cmp"
	"slices"
)

func QuickSort[T cmp.Ordered](slice []T) []T {
	if len(slice) < 2 {
		return slice
	}
	left, right := 0, len(slice)-1
	pivot := len(slice) / 2
	slice[pivot], slice[right] = slice[right], slice[pivot]
	for i := range slice {
		if slice[i] < slice[right] {
			slice[left], slice[i] = slice[i], slice[left]
			left++
		}
	}
	slice[left], slice[right] = slice[right], slice[left]
	QuickSort(slice[:left])
	QuickSort(slice[left+1:])
	return slice
}

func ParallelQuickSort[T cmp.Ordered](slice []T) []T {
	if len(slice) < 2 {
		return slice
	}
	left, right := 0, len(slice)-1
	pivot := len(slice) / 2
	slice[pivot], slice[right] = slice[right], slice[pivot]
	for i := range slice {
		if slice[i] < slice[right] {
			slice[left], slice[i] = slice[i], slice[left]
			left++
		}
	}
	slice[left], slice[right] = slice[right], slice[left]
	if len(slice) <= 1024 {
		ParallelQuickSort(slice[:left])
		ParallelQuickSort(slice[left+1:])
		return slice
	} else {
		l, r := make(chan bool), make(chan bool)
		go func() {
			ParallelQuickSort(slice[:left])
			l <- true
		}()
		go func() {
			ParallelQuickSort(slice[left+1:])
			r <- true
		}()
		<-l
		<-r
		return slice
	}
}

func MergeSort[T cmp.Ordered](slice []T) []T {
	if len(slice) < 2 {
		return slice
	}
	mid := len(slice) / 2
	return MergeSlices(MergeSort(slice[:mid]), MergeSort(slice[mid:]))
}

func ParallelMergeSort[T cmp.Ordered](slice []T) []T {
	if len(slice) < 2 {
		return slice
	}
	mid := len(slice) / 2
	if len(slice) <= 1024 {
		return MergeSlices(ParallelMergeSort(slice[:mid]), ParallelMergeSort(slice[mid:]))
	}
	left := make(chan []T)
	right := make(chan []T)
	go func() {
		left <- ParallelMergeSort(slice[:mid])
	}()
	go func() {
		right <- ParallelMergeSort(slice[mid:])
	}()
	return MergeSlices(<-left, <-right)
}

func HeapSort[T cmp.Ordered](slice []T) []T {
	heap := slice
	MaxHeapify(heap)
	for range slice {
		MaxHeapPop(heap)
		heap = heap[:len(heap)-1]
	}
	return slice
}

func BuiltInSort[T cmp.Ordered](slice []T) []T {
	slices.Sort(slice)
	return slice
}
