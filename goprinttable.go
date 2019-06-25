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
		return nil, fmt.Errorf("table is empty")
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

func (g *goPrintTable) formatLine(width int) string {
	var l string
	for i := 1; i != width; i += 1 {
		l += fmt.Sprint("-")
	}
	l += fmt.Sprint("\n")
	return " " + l
}

func (g goPrintTable) expectedLength() int {
	var l int
	var max int
	var colCount int

	for _, row := range g.Table {
		l = 0
		colCount = 0
		for _, col := range row {
			l += len(col)
			colCount += 1
		}

		// 3 is spacer to the left from col " | ", and 2 last after last col " |"
		l = l + (colCount * 3) + 2
		if l > max {
			max = l
		}
	}
	return max
}

func (g *goPrintTable) formatRow(row []string, isHeader bool) string {
	var rowStr string

	tableWidth := g.expectedLength()

	for _, col := range row {
		rowStr = rowStr + " | " + col
	}
	rowStr = rowStr + " |\n"

	if isHeader {
		rowStr = g.formatLine(tableWidth) + rowStr + g.formatLine(tableWidth)
	}
	return rowStr
}


func (g *goPrintTable) formatTable(withHeader bool) string {
	if len(g.Table) == 0 {
		return ""
	}

	maxColWidth, err := g.getMaxColWidth()
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	maxColCount := g.getMaxColCount()

	g.fillTableWithColumns(maxColCount)
	g.fillTableValues(maxColWidth)


	var stringTable string

	tableWidth := g.expectedLength()

	for i, row := range g.Table {
		firstRow := i == 0
		stringTable += g.formatRow(row, firstRow && withHeader)
	}
	stringTable += g.formatLine(tableWidth)

	return stringTable
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


func GetStringTableWithHeader(table [][]string, header bool) string {
	var g = goPrintTable{table, header}
	return g.formatTable(true)
}