package matrix

import (
	"errors"
	"strconv"
	"strings"
)

// Matrix type here.
type Matrix [][]int

func New(s string) (*Matrix, error) {
	rowsString := strings.Split(s, "\n")
	var m Matrix = make([][]int, len(rowsString))

	for i, rowString := range rowsString {
		var row []int
		vals := strings.Split(strings.TrimSpace(rowString), " ")

		if i > 0 && len(vals) != len(m[0]) {
			return nil, errors.New("matrix is not rectangular")
		}

		for _, v := range vals {
			v, err := strconv.Atoi(v)
			if err != nil {
				return nil, err
			}
			row = append(row, v)
		}
		m[i] = row
	}

	return &m, nil
}

// Cols and Rows must return the results without affecting the matrix.
func (m *Matrix) Cols() [][]int {
	newM := make([][]int, len((*m)[0]))
	for i := 0; i < len((*m)[0]); i++ {
		newM[i] = make([]int, len(*m))
		for j := 0; j < len(*m); j++ {
			newM[i][j] = (*m)[j][i]
		}
	}
	return newM
}

func (m *Matrix) Rows() [][]int {
	newM := make([][]int, len(*m))
	for i := 0; i < len(*m); i++ {
		newM[i] = make([]int, len((*m)[0]))
		copy(newM[i], (*m)[i])
	}
	return newM
}

func (m *Matrix) Set(row, col, val int) bool {
	if row < 0 || row >= len(*m) || col < 0 || col >= len((*m)[0]) {
		return false
	}
	(*m)[row][col] = val
	return true
}
