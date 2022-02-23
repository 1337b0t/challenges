package chessboard

// Declare a type named Rank which stores if a square is occupied by a piece - this will be a slice of bools
type Rank []bool

// Declare a type named Chessboard which contains a map of eight Ranks, accessed with keys from "A" to "H"
type Chessboard map[string]Rank

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank
func CountInRank(cb Chessboard, rank string) int {
	count := 0
	if _, exists := cb[rank]; exists {
		for _, v := range cb[rank] {
			if v {
				count++
			}
		}
	}
	return count
}

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file
func CountInFile(cb Chessboard, file int) int {
	count := 0
	if validFile(file) {
		for rank := range cb {
			for fileIdx, v := range cb[rank] {
				if fileIdx+1 == file {
					if v {
						count++
					}
				}
			}
		}
	}

	return count
}

// CountAll should count how many squares are present in the chessboard
func CountAll(cb Chessboard) int {
	count := 0
	for rank := range cb {
		for range cb[rank] {
			count++
		}
	}
	return count
}

// CountOccupied returns how many squares are occupied in the chessboard
func CountOccupied(cb Chessboard) int {
	count := 0
	for rank := range cb {
		for _, occupied := range cb[rank] {
			if occupied {
				count++
			}
		}
	}
	return count
}

func validFile(file int) bool {
	return file > 0 && file <= 8
}
