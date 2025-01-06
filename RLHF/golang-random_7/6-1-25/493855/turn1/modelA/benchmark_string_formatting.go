package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
	"time"
)

const numFormats = 1000000

func benchmarkFmtSprintf() time.Duration {
	start := time.Now()
	for i := 0; i < numFormats; i++ {
		fmt.Sprintf("Hello, %d!", i)
	}

	return time.Since(start)
}

func benchmarkStringBuilder() time.Duration {
	start := time.Now()
	var b strings.Builder
	for i := 0; i < numFormats; i++ {
		b.WriteString("Hello, ")
		b.WriteString(fmt.Sprintf("%d", i))
		b.WriteString("!")
		b.Reset()
	}
	return time.Since(start)
}

func benchmarkStrconvFormat() time.Duration {
	start := time.Now()
	for i := 0; i < numFormats; i++ {
		strconv.FormatInt(int64(i), 10)
	}
	return time.Since(start)
}

func benchmarkCharm() time.Duration {
	start := time.Now()
	for i := 0; i < numFormats; i++ {
		charm.Printf("Hello, %d!", i)
	}
	return time.Since(start)
}

func benchmarkBufio() time.Duration {
	start := time.Now()
	var buf strings.Builder
	for i := 0; i < numFormats; i++ {
		bufio.NewWriter(&buf).WriteString("Hello, ")
		bufio.NewWriter(&buf).WriteString(fmt.Sprintf("%d", i))
		bufio.NewWriter(&buf).WriteString("!")
		buf.Reset()
	}
	return time.Since(start)
}

func main() {
	fmt.Println("Benchmarking fmt.Sprintf:", benchmarkFmtSprintf())
	fmt.Println("Benchmarking strings.Builder:", benchmarkStringBuilder())
	fmt.Println("Benchmarking strconv.Format:", benchmarkStrconvFormat())
	fmt.Println("Benchmarking charm:", benchmarkCharm())
	fmt.Println("Benchmarking bufio:", benchmarkBufio())
}
