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
			for k, val := range row.Cells {
				SetValue(k, val, cellValue.Index(index))
			}
		}
	}
	return nil
}

func SetValue(i int, value *xlsx.Cell, v reflect.Value) {
	// TODO: interface judge
	//tv := reflect.ValueOf(v)
	tv := reflect.Indirect(v)
	logrus.Infof("i: %v, num %v", i, tv.NumField())
	if reflect.Indirect(tv).NumField() < i-1 {
		return
	}
	stv := reflect.Indirect(tv).Field(i)
	//handle cunstomer unmarshal
	u, _, _ := indirect(stv, false)
	if u != nil {
		err := u.UnmarshalXlsx(value.Value)
		logrus.Errorf("err: %v", err)
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