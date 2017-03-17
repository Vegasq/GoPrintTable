package GoPrintTable

import "fmt"

func getMaxColWidth(table [][]string) []int {
	if len(table) == 0 {
		return nil
	}
	columnCount := getMaxColCount(table)
	var result []int
	var colW int

	for i:=0; i<columnCount; i+=1 {
		result = append(result, 0)
	}

	for _, row := range table {
		for colI, col := range row {
			colW = len(col)
			if colW > result[colI] {
				result[colI] = colW
			}
		}
	}

	return result
}

func getMaxColCount(table [][]string) int {
	result := 0
	for _, row := range table {
		if len(row) > result {
			result = len(row)
		}
	}
	return result
}

func fillTableValues(table [][]string, sizes []int) [][]string {
	var diffW int
	var colW int
	var tailEnd string

	for rowI, row := range table {
		for colI, col := range row {
			colW = len(col)
			diffW = sizes[colI] - colW

			tailEnd = ""
			for i:=0; i!=diffW; i++ {
				tailEnd += " "
			}
			table[rowI][colI] = table[rowI][colI] + tailEnd
		}
	}

	return table
}

func fillTableWithColumns(table [][]string, count int) [][]string {
	var wasChanged bool
	for ;; {
		wasChanged = false
		for i, row := range table {
			if len(row) != count {
				table[i] = append(table[i], "-")
				wasChanged = true
			}
		}
		if wasChanged == false {
			return table
		}

	}
}

func printLine(width int){
	for i:=1; i!=width; i+=1{
		fmt.Print("-")
	}
	fmt.Println("")
}

func printIt(table [][]string, withHeader bool){
	if len(table) == 0 {
		return
	}

	var rowStr string
	var rows []string
	for _, row := range table {
		rowStr = "| "
		for _, col := range row {
			rowStr = rowStr + col + " | "
		}
		rows = append(rows, rowStr)
	}
	tableWidth := len(rows[0])

	printLine(tableWidth)
	for i, row := range rows{
		fmt.Println(row)
		if i == 0 && withHeader {
			printLine(tableWidth)
		}
	}
	printLine(tableWidth)
}

func printTable(table [][]string, withHeader bool) {
	maxColWidth := getMaxColWidth(table)
	maxColCount := getMaxColCount(table)


	tableSized := fillTableWithColumns(table, maxColCount)
	tableFilled := fillTableValues(tableSized, maxColWidth)

	printIt(tableFilled, withHeader)
}

func PrintTableWithHeader(table [][]string) {
	printTable(table, true)
}

func PrintTable(table [][]string) {
	printTable(table, false)
}