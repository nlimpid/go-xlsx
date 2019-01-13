package go_xlsx

import (
	"github.com/magiconair/properties/assert"
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
	assert.Equal(t, err, nil)
	v := make([][]Xlsx, 10)
	for k := range v {
		v[k] = make([]Xlsx, 10)
	}
	err = Unmarshal(xlsxFile, &v)
	assert.Equal(t, err ,nil)
	for _, val := range v {
		for _, val2 := range val {
			t.Logf("v : %+v", val2)
		}
	}
}