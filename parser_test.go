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
	v := make([]Xlsx, 10)
	err = Unmarshal(xlsxFile, &v)
	assert.Equal(t, err ,nil)
	for _, val := range v {
		t.Logf("v : %+v", val)
	}
}