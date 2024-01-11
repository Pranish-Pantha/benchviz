package sorting

import (
	"fmt"
	"slices"
	"testing"
)

// Template for a sorting benchmark
func RunIntSortBenchmark(b *testing.B, sortFunc func([]int) []int) {
	for _, v := range TestCases {
		b.Run(fmt.Sprintf("input_size_%d", v.Size), func(b *testing.B) {
			slice := GenerateSlice(v.Size, RandomIntGenerator(10000))
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				sortFunc(slice)
			}
		})
	}
}

func BenchmarkIntQuickSort(b *testing.B) {
	RunIntSortBenchmark(b, QuickSort)
}

func BenchmarkIntMergeSort(b *testing.B) {
	RunIntSortBenchmark(b, MergeSort)
}

func BenchmarkIntBuiltInSort(b *testing.B) {
	RunIntSortBenchmark(b, BuiltInSort)
}

func BenchmarkIntParallelMergeSort(b *testing.B) {
	RunIntSortBenchmark(b, ParallelMergeSort)
}

func BenchmarkIntParallelQuickSort(b *testing.B) {
	RunIntSortBenchmark(b, ParallelQuickSort)
}

func BenchmarkIntHeapSort(b *testing.B) {
	RunIntSortBenchmark(b, HeapSort)
}

// Template for a sorting test
func ValidateIntSortFunc(t *testing.T, sortFunc func([]int) []int, name string) {
	for _, v := range TestCases {
		t.Run(fmt.Sprintf("input_size_%d", v.Size), func(t *testing.T) {
			slice := GenerateSlice(v.Size, RandomIntGenerator(10000))
			sorted := make([]int, len(slice))
			copy(sorted, slice)
			slices.Sort((sorted))
			slice = sortFunc(slice)
			if !slices.Equal(slice, sorted) {
				t.Errorf("%s did not sort correctly", name)
			}
		})
	}
}

func TestMergeSort(t *testing.T) {
	ValidateIntSortFunc(t, MergeSort, "MergeSort")
}

func TestQuickSort(t *testing.T) {
	ValidateIntSortFunc(t, QuickSort, "QuickSort")
}

func TestParallelMergeSort(t *testing.T) {
	ValidateIntSortFunc(t, ParallelMergeSort, "ParallelMergeSort")
}

func TestParallelQuickSort(t *testing.T) {
	ValidateIntSortFunc(t, ParallelQuickSort, "ParallelQuickSort")
}

func TestHeapSort(t *testing.T) {
	ValidateIntSortFunc(t, HeapSort, "HeapSort")
}
