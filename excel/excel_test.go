package excel

import "testing"

func TestExcel(t *testing.T) {
	file, err := OpenFile("./test.xlsx")
	if err != nil {
		t.Fatal(err)
	}

	rows, hasMore, err := file.GetRows(0)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("has more?", hasMore)

	for _, row := range rows {
		t.Log(row)
	}

	row, column, err := file.Size(0)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(row, column)

	t.Log(file.Close())
}
