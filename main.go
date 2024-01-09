package main

import (
	"benchviz/benchmark"
	"benchviz/sorting"
	"benchviz/viz"
	"fmt"
	"os/exec"
)

func ExecBenchmark(benchmarkName, testName string) benchmark.Benchmark {
	out, err := exec.Command("go", "test", "benchviz/sorting", "-bench", testName, "-benchmem").Output()
	if err != nil {
		panic(err)
	}
	return sorting.ParseBenchmark(benchmarkName, string(out))
}

func RenderSortAlgoLineChart() {
	xAxis := make([]string, len(sorting.TestCases))
	for i, testCase := range sorting.TestCases {
		xAxis[i] = fmt.Sprintf("%d", testCase.Size)
	}
	mergeSortBenchmark := ExecBenchmark("MergeSort", "BenchmarkIntMergeSort")
	heapSortBenchmark := ExecBenchmark("HeapSort", "BenchmarkIntHeapSort")
	quickSortBenchmark := ExecBenchmark("QuickSort", "BenchmarkIntQuickSort")
	mergeSortParallelBenchmark := ExecBenchmark("MergeSortParallel", "BenchmarkIntParallelMergeSort")
	quickSortParallelBenchmark := ExecBenchmark("QuickSortParallel", "BenchmarkIntParallelQuickSort")

	linechart := viz.NewLineChart(
		"Sorting Algos in Go",
		fmt.Sprintf("%s, %s", mergeSortBenchmark.GoOS, mergeSortBenchmark.GoArch),
		xAxis,
		[]viz.LineSeries{
			viz.BenchmarkToLineSeries(mergeSortBenchmark),
			viz.BenchmarkToLineSeries(heapSortBenchmark),
			viz.BenchmarkToLineSeries(quickSortBenchmark),
			viz.BenchmarkToLineSeries(mergeSortParallelBenchmark),
			viz.BenchmarkToLineSeries(quickSortParallelBenchmark),
		},
	)
	linechart.Render("linechart.html")

}

func RenderSortAlgoBarChart() {
	xAxis := make([]string, len(sorting.TestCases))
	for i, testCase := range sorting.TestCases {
		xAxis[i] = fmt.Sprintf("%d", testCase.Size)
	}
	mergeSortBenchmark := ExecBenchmark("MergeSort", "BenchmarkIntMergeSort")
	heapSortBenchmark := ExecBenchmark("HeapSort", "BenchmarkIntHeapSort")
	quickSortBenchmark := ExecBenchmark("QuickSort", "BenchmarkIntQuickSort")
	mergeSortParallelBenchmark := ExecBenchmark("MergeSortParallel", "BenchmarkIntParallelMergeSort")
	quickSortParallelBenchmark := ExecBenchmark("QuickSortParallel", "BenchmarkIntParallelQuickSort")

	barchart := viz.NewBarChart(
		"Sorting Algos in Go",
		fmt.Sprintf("%s, %s", mergeSortBenchmark.GoOS, mergeSortBenchmark.GoArch),
		xAxis,
		[]viz.BarSeries{
			viz.BenchmarkToBarSeries(mergeSortBenchmark),
			viz.BenchmarkToBarSeries(heapSortBenchmark),
			viz.BenchmarkToBarSeries(quickSortBenchmark),
			viz.BenchmarkToBarSeries(mergeSortParallelBenchmark),
			viz.BenchmarkToBarSeries(quickSortParallelBenchmark),
		},
	)
	barchart.Render("barchart.html")
}

func main() {
	RenderSortAlgoLineChart()
	RenderSortAlgoBarChart()
}