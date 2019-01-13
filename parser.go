package go_xlsx

import (
	"github.com/sirupsen/logrus"
	"github.com/tealeg/xlsx"
	"reflect"
	"strconv"
)

func Unmarshal(f *xlsx.File, v interface{}) error {
	valueV := reflect.ValueOf(v)
	valueVV := reflect.Indirect(valueV)

	for sKey, sheet := range f.Sheets {
		cellValue := valueVV.Index(sKey)
		for index, row := range sheet.Rows {
			if nil == row {
				break
			}
			logrus.Infof("index: %v cap: %v", index, cellValue.Cap())
			// Get element of array, growing if necessary.
			// Grow slice if necessary
			if index >= cellValue.Cap() {
				newcap := cellValue.Cap() + cellValue.Cap()/2
				if newcap < 4 {
					newcap = 4
				}
				newcellValue := reflect.MakeSlice(cellValue.Type(), cellValue.Len(), newcap)
				reflect.Copy(newcellValue, cellValue)
				cellValue.Set(newcellValue)
			}
			if index >= cellValue.Len() {
				cellValue.SetLen(index + 1)
			}
			for k, val := range row.Cells {
				if k < cellValue.Len() {
					SetValue(k, val, cellValue.Index(index))
				}
			}
		}
	}
	return nil
}

func SetValue(i int, value *xlsx.Cell, v reflect.Value) {
	// TODO: interface judge
	//tv := reflect.ValueOf(v)
	tv := reflect.Indirect(v)
	//logrus.Infof("i: %v, num %v", i, tv.NumField())
	if reflect.Indirect(tv).NumField() < i-1 {
		return
	}
	stv := reflect.Indirect(tv).Field(i)
	//handle cunstomer unmarshal
	u, _, _ := indirect(stv, false)
	if u != nil {
		err := u.UnmarshalXlsx(value.Value)
		if err != nil {
			logrus.Errorf("err: %v", err)
		}
	}
	if stv.Kind() == reflect.Uint64 {

		strin, _ := value. GeneralNumericWithoutScientific()
		xx, _ := strconv.ParseUint(strin, 10, 64)
		stv.SetUint(xx)
	}
	if stv.Kind() == reflect.Int8 {
		xx, _ := strconv.ParseInt(value.Value, 10, 8)
		stv.SetInt(xx)
	}
	if stv.Kind() == reflect.String {
		stv.SetString(value.Value)
	}

}