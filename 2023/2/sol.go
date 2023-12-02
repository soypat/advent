package two

import (
	"bufio"
	"bytes"
	"strconv"
)

func CalculatePower(buf *bufio.Reader) int {
	totalPower := 0
	for {
		l, ispfx, err := buf.ReadLine()
		if ispfx {
			panic("line too long")
		}
		if err != nil {
			break
		}
		_, l, _ = bytes.Cut(l, []byte{':', ' '}) // Remove game ID part.

		parseOK := true
		var minR, minG, minB int
		for parseOK {
			var game []byte
			game, l, parseOK = bytes.Cut(l, []byte{';'})
			r, g, b := parseCubeshow(game)
			minR = max(r, minR)
			minG = max(g, minG)
			minB = max(b, minB)
		}
		power := minR * minG * minB
		totalPower += power
	}
	return totalPower
}

func SumPossible(buf *bufio.Reader, rMax, gMax, bMax int) int {
	idSum := 0
	id := 0
	for {
		id++
		l, ispfx, err := buf.ReadLine()
		if ispfx {
			panic("line too long")
		}
		if err != nil {
			break
		}
		_, l, _ = bytes.Cut(l, []byte{':', ' '}) // Remove game ID part.
		idValid := true
		parseOK := true
		for parseOK && idValid {
			var game []byte
			game, l, parseOK = bytes.Cut(l, []byte{';'})
			r, g, b := parseCubeshow(game)
			showValid := r <= rMax && g <= gMax && b <= bMax
			idValid = idValid && showValid
		}
		if idValid {
			idSum += id
		}
	}
	return idSum
}
func parseCubeshow(buf []byte) (r, g, b int) {
	ok := true
	for ok {
		var cube []byte
		buf = bytes.TrimSpace(buf)
		cube, buf, ok = bytes.Cut(buf, []byte{',', ' '})
		number, color, _ := bytes.Cut(cube, []byte{' '})
		gotnum, err := strconv.ParseUint(string(number), 10, 32)
		if err != nil {
			panic(err.Error())
		}
		if color[0] == 'r' {
			r = int(gotnum)
		} else if color[0] == 'g' {
			g = int(gotnum)
		} else if color[0] == 'b' {
			b = int(gotnum)
		}
	}
	return r, g, b
}
