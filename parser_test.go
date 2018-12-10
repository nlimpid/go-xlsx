package xlsx_parser

import (
	"github.com/tealeg/xlsx"
	"testing"
)

type Xlsx struct {
	OrderID        uint64
	Status         int8
	CreateTime     xlsxTime
	ExpressCompany string
	ExpressNumber  string
}

func TestUnmarshal(t *testing.T) {
	xlsxFile, err := xlsx.OpenFile("file.xlsx")
	if err != nil {
		t.Errorf("open file err = %v\n", err)
	}
	v := Xlsx{}
	Unmarshal(xlsxFile, &v)
	t.Logf("v : %+v", v)
}