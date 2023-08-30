package position

type Position struct {
	Filename string
	Offset   int
	Line     int
	Column   int
}

func (pos *Position) isValid() bool {
	return pos.Line > 0
}

func New(fileName string, offset int, line int, column int) *Position {
	p := &Position{Filename: fileName, Offset: offset, Line: line, Column: column}
	return p
}
