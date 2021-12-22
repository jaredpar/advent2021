package util

type Grid struct {
	Values       []int
	ColumnLength int
}

func NewGrid(values []int, columnLength int) *Grid {
	return &Grid{Values: values, ColumnLength: columnLength}
}

func (g *Grid) Count() int {
	return len(g.Values)
}

func (g *Grid) Index(row, column int) int {
	return (row * g.ColumnLength) + column
}

func (g *Grid) Rows() int {
	return len(g.Values) / g.ColumnLength
}

func (g *Grid) Columns() int {
	return g.ColumnLength
}

func (g *Grid) Value(row, column int) int {
	index := g.Index(row, column)
	return g.Values[index]
}

func (g *Grid) SetValue(row, column, value int) {
	index := g.Index(row, column)
	g.Values[index] = value
}
