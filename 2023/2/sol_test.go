package two

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"testing"
	"time"
)

func TestCalculatePowerTest(t *testing.T) {
	testCalculatePower(t, "test.txt", 2286)
}

func TestCalculatePowerReal(t *testing.T) {
	testCalculatePower(t, "input.txt", 65122)
}

func testCalculatePower(t *testing.T, filename string, expected int) {
	start := time.Now()
	defer func() {
		fmt.Printf("elapsed: %v\n", time.Since(start))
	}()
	fd, _ := os.Open(filename)
	buf := bufio.NewReader(fd)
	sum := CalculatePower(buf)
	if sum != expected {
		t.Errorf("got %v, expected %v", sum, expected)
	}
}

func TestSumPossible(t *testing.T) {
	testSumPossible(t, "test.txt", 8)
}

func TestSumPossibleRealInput(t *testing.T) {
	testSumPossible(t, "input.txt", 2879)
}

func testSumPossible(t *testing.T, filename string, expect int) {
	start := time.Now()
	defer func() {
		fmt.Printf("elapsed: %v\n", time.Since(start))
	}()
	fd, _ := os.Open(filename)
	buf := bufio.NewReader(fd)
	sum := SumPossible(buf, 12, 13, 14)
	if sum != expect {
		t.Errorf("got %v, expected %v", sum, expect)
	}
}

func BenchmarkPartOne(b *testing.B) {
	buf, _ := os.ReadFile("input.txt")
	srcBuf := bytes.NewBuffer(buf)
	bbuf := bufio.NewReader(srcBuf)
	for i := 0; i < b.N; i++ {
		buf := *srcBuf // Hard copy of the buffer to prevent allocations.
		bbuf.Reset(&buf)
		sum := SumPossible(bbuf, 12, 13, 14)
		if sum != 2879 {
			panic("bad sum")
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
		sum := CalculatePower(bbuf)
		if sum != 65122 {
			panic("bad sum")
		}
	}
}
