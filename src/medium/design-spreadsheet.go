package medium

import (
	"strconv"
	"strings"
)

type Spreadsheet struct {
	cols [26]map[int]int
}

func SpreadsheetConstructor(rows int) Spreadsheet {
	var cols [26]map[int]int
	for i := 0; i < 26; i++ {
		cols[i] = make(map[int]int)
	}
	return Spreadsheet{cols: cols}
}

func (s *Spreadsheet) parseCell(cell string) (col, row int) {
	col = int(cell[0] - 'A')
	row, _ = strconv.Atoi(cell[1:])
	return
}

func (s *Spreadsheet) SetCell(cell string, value int) {
	col, row := s.parseCell(cell)
	if value == 0 {
		delete(s.cols[col], row)
		return
	}
	s.cols[col][row] = value
}

func (s *Spreadsheet) ResetCell(cell string) {
	col, row := s.parseCell(cell)
	delete(s.cols[col], row)
}

func (s *Spreadsheet) GetValue(formula string) int {
	// "=X+Y+..."
	if len(formula) == 0 || formula[0] != '=' {
		return 0
	}
	sum := 0
	for _, term := range strings.Split(formula[1:], "+") {
		t := strings.TrimSpace(term)
		if t == "" {
			continue
		}

		if t[0] >= '0' && t[0] <= '9' {
			n, _ := strconv.Atoi(t)
			sum += n
			continue
		}

		c, r := s.parseCell(t)
		sum += s.cols[c][r]
	}
	return sum
}
