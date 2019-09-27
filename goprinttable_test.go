package GoPrintTable

import (
	"testing"
)

func getTableExample() [][]string {
	header := []string{"Node Name", "IP", "Status"}
	node1 := []string{"controller", "192.168.0.100", "Online"}
	node2 := []string{"minion1", "192.168.0.101", "Offline"}
	node3 := []string{"minion2", "192.168.0.102"}
	t := [][]string{header, node1, node2, node3}
	return t
}

func TestGetMaxColWidth(t *testing.T) {
	table := getTableExample()
	var g = goPrintTable{table, false}
	maxColWidth, _ := g.getMaxColWidth()
	if maxColWidth[0] != 10 {
		t.Errorf("maxColWidth[2] should be 10 but it's %d.",
			maxColWidth[0])
	}
	if maxColWidth[1] != 13 {
		t.Errorf("maxColWidth[2] should be 13 but it's %d.",
			maxColWidth[1])
	}
	if maxColWidth[2] != 7 {
		t.Errorf("maxColWidth[2] should be 7 but it's %d.",
			maxColWidth[2])
	}

}

func TestGetMaxColCount(t *testing.T) {
	table := getTableExample()
	var g = goPrintTable{table, false}
	maxColCount := g.getMaxColCount()
	if maxColCount != 3 {
		t.Errorf("MaxColCount should be 3 but it's %d.", maxColCount)
	}
}

func TestFillTableWithColumns(t *testing.T) {
	table := getTableExample()
	var g = goPrintTable{table, false}
	maxColCount := g.getMaxColCount()
	g.fillTableWithColumns(maxColCount)

	firstRowLen := len(g.Table[0])

	for k, v := range g.Table {
		if firstRowLen != len(v) {
			t.Errorf("Rows was not alligned properly. Row len(%s)"+
				" != len(%s) ", g.Table[0], g.Table[k])
		}
	}

}

func TestFillTableValues(t *testing.T) {
	table := getTableExample()

	var g = goPrintTable{table, false}

	maxColCount := g.getMaxColCount()
	maxColWidth, _ := g.getMaxColWidth()

	g.fillTableWithColumns(maxColCount)
	g.fillTableValues(maxColWidth)

	rowLen := len(g.Table[0])
	for i := 0; i < rowLen; i += 1 {
		basicLen := len(g.Table[0][i])

		for k := range g.Table {
			if len(g.Table[k][i]) != basicLen {
				t.Errorf("Column %s wasn't fill correctly.", g.Table[k][i])
			}
		}

	}
}

func TestGetStringTableWithHeader(t *testing.T) {
	type args struct {
		table  [][]string
		header bool
	}

	simpleTableResult :=
` ----------------
 | Name   | Age |
 ----------------
 | Donald | 3   |
 ----------------
`

	unevenTableResult :=
` ----------------
 | Name   | Age |
 ----------------
 | Donald | -   |
 ----------------
`

	tests := []struct {
		name string
		args args
		want string
	}{
		{"Simple table", args{[][]string{[]string{"Name", "Age"}, []string{"Donald", "3"}}, true}, simpleTableResult},
		{"Uneven table", args{[][]string{[]string{"Name", "Age"}, []string{"Donald"}}, true}, unevenTableResult},
		{"Empty table", args{[][]string{}, true}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetStringTableWithHeader(tt.args.table, tt.args.header); got != tt.want {
				t.Errorf("GetStringTableWithHeader() = %v, want %v", got, tt.want)
			}
		})
	}
}
