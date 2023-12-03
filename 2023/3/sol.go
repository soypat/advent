package three

import (
	"bufio"
	"errors"
	"fmt"
	"strconv"
)

const invalidSymbol = 0xff

func SumPartNumbers(buf *bufio.Reader) int {
	MAT, err := parsePartMatrix(buf)
	if err != nil {
		panic(err)
	}
	sum := 0
	r, c := MAT.Dims()
	for i := 0; i < r; i++ {

		var numStart = -1
		var isValidNum bool
		for j := 0; j <= c; j++ {
			char := MAT.At(i, j)
			isdig := isdigit(char)
			if numStart == -1 && !isdig {
				continue // Not a number nor are we keeping track of one.
			} else if numStart == -1 && isdig {
				numStart = j // Begin keeping track of number.
			}
			if isdig {
				// We are keeping track of number, check if valid if not yet confirmed.
				if !isValidNum {
					isValidNum = MAT.IsSymbolNeighbor(i, j)
				}
				continue
			}
			// Gotten to this point we were tracking a number but it's over. Parse and add if valid
			// We are no longer in a number, if is valid we parse and add.
			if isValidNum {
				// Parse and add.
				offset := i * MAT.cols
				numstr := string(MAT.data[offset+numStart : offset+j])
				v, err := strconv.ParseUint(numstr, 10, 32)
				if err != nil {
					panic(err)
				}
				fmt.Println("adding", v)
				sum += int(v)
				isValidNum = false
			}
			numStart = -1 // Reset number tracking.
		}
	}
	return sum
}

func parsePartMatrix(buf *bufio.Reader) (MAT partmatrix, err error) {
	rows := 0
	for {
		l, pfx, err := buf.ReadLine()
		if pfx {
			return MAT, errors.New("line too long")
		}
		if err != nil || len(l) <= 1 {
			break
		}
		MAT.data = append(MAT.data, l...)
		if MAT.cols == 0 {
			MAT.cols = len(l)
		} else if MAT.cols != len(l) {
			return MAT, errors.New("inconsistent row length")
		}
		rows++
	}
	MAT.rows = rows
	return MAT, err
}

type partmatrix struct {
	data []byte
	cols int
	rows int
}

func (m *partmatrix) Dims() (r, c int) {
	return m.rows, m.cols
}

func (m *partmatrix) At(r, c int) byte {
	if uint(r) >= uint(m.rows) || uint(c) >= uint(m.cols) || m.data[r*m.cols+c] == '.' {
		return invalidSymbol
	}
	return m.data[r*m.cols+c]
}

func isdigit(char byte) bool {
	return char >= '0' && char <= '9'
}

func (m *partmatrix) IsSymbolNeighbor(r, c int) bool {
	return isSymbol(m.At(r-1, c-1)) || isSymbol(m.At(r-1, c)) || isSymbol(m.At(r-1, c+1)) ||
		isSymbol(m.At(r, c-1)) || isSymbol(m.At(r, c+1)) || isSymbol(m.At(r+1, c-1)) ||
		isSymbol(m.At(r+1, c)) || isSymbol(m.At(r+1, c+1))
}

func isSymbol(c byte) bool {
	return c != invalidSymbol && !isdigit(c)
}
