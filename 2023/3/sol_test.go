package three

import (
	"bufio"
	"fmt"
	"os"
	"testing"
	"time"
)

func TestSumGearRatios(t *testing.T) {
	testSumGearRatios(t, "test.txt", 467835)
}

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

func testSumGearRatios(t *testing.T, filename string, expected int) {
	start := time.Now()
	defer func() {
		fmt.Printf("elapsed: %v\n", time.Since(start))
	}()
	fd, _ := os.Open(filename)
	buf := bufio.NewReader(fd)
	sum := SumGearRatios(buf)
	if sum != expected {
		t.Errorf("got %v, expected %v", sum, expected)
	}
}
