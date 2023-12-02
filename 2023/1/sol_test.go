package one

import (
	"bufio"
	"bytes"
	"os"
	"testing"
)

func BenchmarkPartOne(b *testing.B) {
	buf, _ := os.ReadFile("input.txt")
	srcBuf := bytes.NewBuffer(buf)
	bbuf := bufio.NewReader(srcBuf)
	for i := 0; i < b.N; i++ {
		buf := *srcBuf // Hard copy of the buffer to prevent allocations.
		bbuf.Reset(&buf)
		sum := ONE(bbuf)
		if sum == 0 {
			panic("sum is zero")
		}
	}
}

func BenchmarkPartTwo(b *testing.B) {
	buf, _ := os.ReadFile("input.txt")
	srcBuf := bytes.NewBuffer(buf)
	bbuf := bufio.NewReader(srcBuf)
	for i := 0; i < b.N; i++ {
		buf := *srcBuf // Hard copy of the buffer to prevent allocations.
		bbuf.Reset(&buf)
		sum := TWO(bbuf)
		if sum == 0 {
			panic("sum is zero")
		}
	}
}
