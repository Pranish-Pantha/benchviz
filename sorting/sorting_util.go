package sorting

import (
	"benchviz/benchmark"
	"benchviz/parsing"
	"cmp"
	"math/rand"
)

var TestCases = []struct {
	Size int
}{
	{Size: 50000},
	{Size: 100000},
	{Size: 500000},
	{Size: 1000000},
}

func RandomIntGenerator(max int) func(int) int {
	return func(i int) int {
		return rand.Intn(max)
	}
}

func GenerateSlice[T cmp.Ordered](size int, generator func(int) T) []T {
	slice := make([]T, size)
	for i := 0; i < size; i++ {
		slice[i] = generator(i)
	}
	return slice
}

func MergeSlices[T cmp.Ordered](left, right []T) []T {
	result := make([]T, len(left)+len(right))
	i := 0
	for len(left) > 0 && len(right) > 0 {
		if left[0] <= right[0] {
			result[i] = left[0]
			left = left[1:]
		} else {
			result[i] = right[0]
			right = right[1:]
		}
		i++
	}
	for j := 0; j < len(left); j++ {
		result[i] = left[j]
		i++
	}
	for j := 0; j < len(right); j++ {
		result[i] = right[j]
		i++
	}
	return result
}

func MaxHeapShiftDown[T cmp.Ordered](slice []T, i, size int) {
	for i*2+1 < size {
		child := i*2 + 1
		if child+1 < size && slice[child+1] > slice[child] {
			child++
		}
		if slice[i] >= slice[child] {
			return
		}
		slice[i], slice[child] = slice[child], slice[i]
		i = child
	}
}

func MaxHeapPop[T cmp.Ordered](slice []T) {
	slice[0], slice[len(slice)-1] = slice[len(slice)-1], slice[0]
	MaxHeapShiftDown(slice, 0, len(slice)-1)
}

func MaxHeapify[T cmp.Ordered](slice []T) {
	for i := (len(slice) / 2) - 1; i >= 0; i-- {
		MaxHeapShiftDown(slice, i, len(slice))
	}
}

func ParseBenchmark(name, rawBenchmark string) benchmark.Benchmark {
	benchmark := parsing.ParseBenchmark(name, rawBenchmark)
	for i, testCase := range TestCases {
		benchmark.Tests[i].InputSize = testCase.Size
	}
	return benchmark
}
