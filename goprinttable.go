package GoPrintTable

import "fmt"

// GoPrintTable print beautiful table.
type goPrintTable struct {
	Table      [][]string
	WithHeader bool
}

// getMaxColWidth Find biggest value (length) per column, and return slice of
// int, where each value represent column with same index.
func (g *goPrintTable) getMaxColWidth() ([]int, error) {
	if len(g.Table) == 0 {
		return nil, fmt.Errorf("Table is empty.")
	}
	columnCount := g.getMaxColCount()
	var result []int
	var colW int

	for i := 0; i < columnCount; i += 1 {
		result = append(result, 0)
	}

	for _, row := range g.Table {
		for colI, col := range row {
			colW = len(col)
			if colW > result[colI] {
				result[colI] = colW
			}
		}
	}

	return result, nil
}

// getMaxColCount Find row with Max count of elements and return len(row)
func (g *goPrintTable) getMaxColCount() int {
	result := 0
	for _, row := range g.Table {
		if len(row) > result {
			result = len(row)
		}
	}
	return result
}

// fillTableValues append spaces to column values to have same len(col) per col
func (g *goPrintTable) fillTableValues(sizes []int) {
	var diffW int
	var colW int
	var tailEnd string

	for rowI, row := range g.Table {
		for colI, col := range row {
			colW = len(col)
			diffW = sizes[colI] - colW

			tailEnd = ""
			for i := 0; i != diffW; i++ {
				tailEnd += " "
			}
			g.Table[rowI][colI] = g.Table[rowI][colI] + tailEnd
		}
	}
}

// fillTableWithColumns for each row, append empty values to slice til
// len(row) == sizes
func (g *goPrintTable) fillTableWithColumns(count int) {
	var wasChanged bool
	for {
		wasChanged = false
		for i, row := range g.Table {
			if len(row) != count {
				g.Table[i] = append(g.Table[i], "-")
				wasChanged = true
			}
		}
		if wasChanged == false {
			return
		}

	}
}

// printLine print separator line
func (g *goPrintTable) printLine(width int) {
	for i := 1; i != width; i += 1 {
		fmt.Print("-")
	}
	fmt.Println("")
}

// printIt join values and print them as table
func (g *goPrintTable) printIt(withHeader bool) {
	if len(g.Table) == 0 {
		return
	}

	var rowStr string
	var rows []string
	for _, row := range g.Table {
		rowStr = "| "
		for _, col := range row {
			rowStr = rowStr + col + " | "
		}
		rows = append(rows, rowStr)
	}
	tableWidth := len(rows[0])

	g.printLine(tableWidth)
	for i, row := range rows {
		fmt.Println(row)
		if i == 0 && withHeader {
			g.printLine(tableWidth)
		}
	}
	g.printLine(tableWidth)
}

// printTable do math and pass result to printIt()
func (g *goPrintTable) printTable() {
	maxColWidth, err := g.getMaxColWidth()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	maxColCount := g.getMaxColCount()

	g.fillTableWithColumns(maxColCount)
	g.fillTableValues(maxColWidth)

	g.printIt(g.WithHeader)
}

// PrintTableWithHeader does formatted print of [][]string in beautiful ascii
// table with separate header that should be first row in argument.
func PrintTableWithHeader(table [][]string) {
	var g = goPrintTable{table, true}
	g.printTable()
}

// PrintTable does formatted print of [][]string in beautiful ascii table.
func PrintTable(table [][]string) {
	var g = goPrintTable{table, false}
	g.printTable()
}
