package one

import (
	"bufio"
)

func TWO(buf *bufio.Reader) int {
	var sum int
	for {
		line, _, err := buf.ReadLine()
		if err != nil {
			break
		}
		first := -1
		i := 0
		for i < len(line) {
			digit, advance := getDigit2(line[i:])
			i += advance
			if digit >= 0 {
				first = digit
				break
			}
		}

		last := first
		j := len(line) - 1
		for j >= i {
			digit, _ := getDigit2(line[j:])
			j--
			if digit >= 0 {
				last = digit
				break
			}
		}
		sum += 10*first + last
	}
	return sum
}

var letterIdx = [256][]int{
	'o': {1},
	't': {2, 3},
	'f': {4, 5},
	's': {6, 7},
	'e': {8},
	'n': {9},
}

var nums = [10]string{
	1: "one",
	2: "two",
	3: "three",
	4: "four",
	5: "five",
	6: "six",
	7: "seven",
	8: "eight",
	9: "nine",
}

func getDigit2(buf []byte) (int, int) {
	c := buf[0]
	if c >= '0' && c <= '9' {
		return int(c) - '0', 1
	}
	for _, i := range letterIdx[c] {
		if string(buf[:min(len(buf), len(nums[i]))]) == nums[i] {
			return i, len(nums[i])
		}
	}
	return -1, 1
}

func ONE(buf *bufio.Reader) int {
	var sum int
	for {
		line, _, err := buf.ReadLine()
		if err != nil {
			break
		}
		first := -1
		last := -1
		for i := 0; i < len(line); i++ {
			digit := getDigit(line[i])
			if digit < 0 {
				continue
			}
			if first == -1 {
				first = digit
			}
			last = digit
		}
		sum += 10*first + last
	}
	return sum
}

func getDigit(c byte) int {
	if c >= '0' && c <= '9' {
		return int(c) - '0'
	}
	return -1
}
