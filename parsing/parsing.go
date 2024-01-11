package parsing

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Pranish-Pantha/benchviz/benchmark"
)

// Parses a benchmark output from the go test -bench command
func ParseBenchmark(name, rawBenchmark string) benchmark.Benchmark {
	split := strings.Split(rawBenchmark, "\n")
	goOS := strings.Split(split[0], " ")[1]
	goArch := strings.Split(split[1], " ")[1]
	benchmark := benchmark.Benchmark{
		GoOS:   goOS,
		GoArch: goArch,
		Name:   name,
		Tests:  make([]benchmark.BenchMarkTestCase, 0),
	}
	for _, line := range split[3:] {
		arr := strings.Fields(line)
		if len(arr) <= 1 {
			break
		}
		testCase, err := ParseTestCase(arr)
		if err != nil {
			fmt.Printf("Error parsing line: %s. %v\n", line, err)
			continue
		}
		benchmark.Tests = append(benchmark.Tests, testCase)
	}
	return benchmark
}

// Parses a single line of a benchmark output representing a single test case
func ParseTestCase(line []string) (benchmark.BenchMarkTestCase, error) {
	if len(line) < 7 {
		return benchmark.BenchMarkTestCase{}, fmt.Errorf("line is too short")
	}
	iterations, err := strconv.ParseInt(line[1], 10, 0)
	if err != nil {
		return benchmark.BenchMarkTestCase{}, err
	}
	nanosecondsPerOp, err := strconv.ParseFloat(line[2], 32)
	if err != nil {
		return benchmark.BenchMarkTestCase{}, err
	}
	return benchmark.BenchMarkTestCase{
		Iterations:       int(iterations),
		NanosecondsPerOp: int(nanosecondsPerOp),
	}, nil
}
