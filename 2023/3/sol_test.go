package three

import (
	"bufio"
	"fmt"
	"os"
	"testing"
	"time"
)

func TestSumPartNumbersReal(t *testing.T) {
	testSumPartNumbers(t, "input.txt", 530849)
}

func TestSumPartNumbers(t *testing.T) {
	testSumPartNumbers(t, "test.txt", 4361)
}

func testSumPartNumbers(t *testing.T, filename string, expected int) {
	start := time.Now()
	defer func() {
		fmt.Printf("elapsed: %v\n", time.Since(start))
	}()
	fd, _ := os.Open(filename)
	buf := bufio.NewReader(fd)
	sum := SumPartNumbers(buf)
	if sum != expected {
		t.Errorf("got %v, expected %v", sum, expected)
	}
}
