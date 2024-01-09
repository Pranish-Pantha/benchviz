package benchmark

type Benchmark struct {
	GoOS   string
	GoArch string
	Name   string
	Tests  []BenchMarkTestCase
}

type BenchMarkTestCase struct {
	InputSize        int
	Iterations       int
	NanosecondsPerOp int
	BytesPerOp       int
	AllocsPerOp      int
}
