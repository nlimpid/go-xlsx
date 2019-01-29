package go_xlsx

import (
	"testing"

	"github.com/magiconair/properties/assert"
	"github.com/tealeg/xlsx"
)

type Xlsx struct {
	Status         int8     `xlsx:"status"`
	OrderID        uint64   `xlsx:"order_id"`
	CreateTime     xlsxTime `xlsx:"create_time"`
	ExpressCompany string   `xlsx:"express_company"`
	ExpressNumber  string   `xlsx:"物流号"`
}

func TestUnmarshal(t *testing.T) {
	xlsxFile, err := xlsx.OpenFile("file.xlsx")
	assert.Equal(t, err, nil)
	v := make([][]Xlsx, 10)
	for k := range v {
		v[k] = make([]Xlsx, 10)
	}
	err = Unmarshal(xlsxFile, &v)
	assert.Equal(t, err, nil)
	for _, val := range v {
		for _, val2 := range val {
			if val2.Status != 0 {
				t.Logf("v : %+v, express_number: %v", val2.Status, val2.ExpressNumber)
			}
		}
	}
}
