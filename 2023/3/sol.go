package three

import (
	"bufio"
	"errors"
	"fmt"
	"strconv"
)

const invalidSymbol = 0xff

func SumGearRatios(buf *bufio.Reader) int {
	MAT, err := parsePartMatrix(buf)
	if err != nil {
		panic(err)
	}
	possibleGears := make(map[[2]int][2]int)
	addToGear := func(r, c, value int) {
		idx := [2]int{r, c}
		got, ok := possibleGears[idx]
		if !ok {
			got[0] = value
		} else if got[1] != 0 {
			got = [2]int{-1, -1} // Invalidate gear.
		} else if got[0] != 0 {
			fmt.Println("got second val", value, "for", got[0])
			got[1] = value
		}
		possibleGears[idx] = got
	}
	startIdxOffsets := [5][2]int{
		{-1, -1}, {-1, 0},
		{0, -1},
		{1, -1}, {1, 0},
	}
	middleIdxOffsets := [2][2]int{
		{-1, 0},
		{1, 0},
	}
	endIdxOffsets := [5][2]int{
		{-1, 0}, {-1, 1},
		{0, 1},
		{1, 0}, {1, 1},
	}
	forEachPart(&MAT, func(r, cs, ce, partNo int) {
		var offsets [][2]int
		for i := cs; i <= ce; i++ {
			if i == cs {
				offsets = startIdxOffsets[:]
			} else if i == ce {
				offsets = endIdxOffsets[:]
			} else {
				offsets = middleIdxOffsets[:]
			}
			for _, offset := range offsets {
				if isGear(MAT.At(r+offset[0], i+offset[1])) {
					addToGear(r+offset[0], i+offset[1], partNo)
				}
			}
		}
	})

	gearSum := 0
	for _, gear := range possibleGears {
		if gear[0] <= 0 || gear[1] <= 0 {
			continue
		}
		fmt.Println("add gear", gear)
		gearSum += gear[0] * gear[1]

	}
	return gearSum
}

func SumPartNumbers(buf *bufio.Reader) int {
	MAT, err := parsePartMatrix(buf)
	if err != nil {
		panic(err)
	}
	sum := 0
	forEachPart(&MAT, func(r, cs, ce, partNo int) {
		sum += partNo
	})
	return sum
}

func forEachPart(MAT *partmatrix, fn func(r, cs, ce, partNo int)) {
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
				fn(i, numStart, j, int(v))
				isValidNum = false
			}
			numStart = -1 // Reset number tracking.
		}
	}
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

func isGear(c byte) bool {
	return c == '*'
}
