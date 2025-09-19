package medium

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	s := SpreadsheetConstructor(3)
	assert.Equal(t, 12, s.GetValue("=5+7"))
	s.SetCell("A1", 10)
	assert.Equal(t, 16, s.GetValue("=A1+6"))
	s.SetCell("B2", 15)
	assert.Equal(t, 25, s.GetValue("=A1+B2"))
	s.ResetCell("A1")
	assert.Equal(t, 15, s.GetValue("=A1+B2"))
}

func TestUnsetCellsAreZero(t *testing.T) {
	s := SpreadsheetConstructor(5)
	assert.Equal(t, 0, s.GetValue("=A1+B3"))
}

func TestSetAndReset(t *testing.T) {
	s := SpreadsheetConstructor(10)
	s.SetCell("C4", 7)
	s.SetCell("D4", 8)
	assert.Equal(t, 15, s.GetValue("=C4+D4"))
	s.ResetCell("C4")
	assert.Equal(t, 8, s.GetValue("=C4+D4"))
}

func TestSetZeroBehavesAsZero(t *testing.T) {
	s := SpreadsheetConstructor(2)
	s.SetCell("A1", 0)
	s.SetCell("B2", 9)
	assert.Equal(t, 9, s.GetValue("=A1+B2"))
}

func TestMultiplePlusTerms(t *testing.T) {
	s := SpreadsheetConstructor(3)
	s.SetCell("A1", 1)
	s.SetCell("B1", 2)
	s.SetCell("C1", 3)
	assert.Equal(t, 15, s.GetValue("=A1+B1+C1+4+5"))
}

func TestNumbersOnly(t *testing.T) {
	s := SpreadsheetConstructor(1)
	assert.Equal(t, 0, s.GetValue("=0+0"))
	assert.Equal(t, 105, s.GetValue("=105+0"))
}

func TestWhitespaceInFormulaTerms(t *testing.T) {
	s := SpreadsheetConstructor(3)
	s.SetCell("A2", 11)
	assert.Equal(t, 14, s.GetValue("=  3+ A2"))
}

func TestHighColumnLetters(t *testing.T) {
	s := SpreadsheetConstructor(3)
	s.SetCell("Z3", 21)
	assert.Equal(t, 25, s.GetValue("=Z3+4"))
}

func TestOverwriteCell(t *testing.T) {
	s := SpreadsheetConstructor(4)
	s.SetCell("D1", 5)
	assert.Equal(t, 5, s.GetValue("=D1+0"))
	s.SetCell("D1", 17)
	assert.Equal(t, 20, s.GetValue("=D1+3"))
}

func TestManyOperationsSequence(t *testing.T) {
	s := SpreadsheetConstructor(6)
	assert.Equal(t, 0, s.GetValue("=A1+A2"))
	assert.Equal(t, 15, s.GetValue("=7+8"))
	s.SetCell("A1", 2)
	s.SetCell("A2", 3)
	assert.Equal(t, 5, s.GetValue("=A1+A2"))
	s.SetCell("B6", 100)
	assert.Equal(t, 102, s.GetValue("=B6+A1"))
	s.SetCell("A1", 0)
	assert.Equal(t, 100, s.GetValue("=A1+B6"))
	s.ResetCell("B6")
	assert.Equal(t, 3, s.GetValue("=A2+B6"))
}
