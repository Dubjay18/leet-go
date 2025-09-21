package main

func isValidSudoku(board [][]byte) bool {
	var rows, cols, boxes [9]int
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			ch := board[r][c]
			if ch == '.' {
				continue
			}
			d := int(ch - '1')
			mask := 1 << d
			bi := (r/3)*3 + (c / 3)

			// If the bit is already set in any of the three, it's a duplicate.
			if rows[r]&mask != 0 || cols[c]&mask != 0 || boxes[bi]&mask != 0 {
				return false
			}
			// Mark the digit as seen in the row, column, and box.
			rows[r] |= mask
			cols[c] |= mask
			boxes[bi] |= mask
		}
	}
	return true
}

func main() {

}
