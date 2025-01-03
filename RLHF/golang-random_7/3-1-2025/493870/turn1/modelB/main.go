package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"runtime"
	"strconv"
	"strings"
	"testing"
	"time"
)

func BenchmarkSprintf(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < b.N; i++ {
		x := rand.Intn(1000)
		y := rand.Intn(1000)
		fmt.Sprintf("%d + %d = %d", x, y, x+y)
	}
}

func BenchmarkBufferSprintf(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	buf := bytes.NewBuffer(make([]byte, 0, 1024))
	for i := 0; i < b.N; i++ {
		x := rand.Intn(1000)
		y := rand.Intn(1000)
		buf.Reset()
		fmt.Fprintf(buf, "%d + %d = %d", x, y, x+y)
	}
}

func BenchmarkIntToString(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < b.N; i++ {
		x := rand.Intn(1000)
		y := rand.Intn(1000)
		strconv.Itoa(x) + " + " + strconv.Itoa(y) + " = " + strconv.Itoa(x+y)
	}
}

func BenchmarkStringBuilder(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	var sb strings.Builder
	for i := 0; i < b.N; i++ {
		x := rand.Intn(1000)
		y := rand.Intn(1000)
		sb.Reset()
		sb.WriteString(strconv.Itoa(x))
		sb.WriteString(" + ")
		sb.WriteString(strconv.Itoa(y))
		sb.WriteString(" = ")
		sb.WriteString(strconv.Itoa(x + y))
		sb.String()
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	testing.Benchmark(BenchmarkSprintf)
	testing.Benchmark(BenchmarkBufferSprintf)
	testing.Benchmark(BenchmarkIntToString)
	testing.Benchmark(BenchmarkStringBuilder)
}
