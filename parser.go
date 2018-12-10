package xlsx_parser

import (
	"github.com/sirupsen/logrus"
	"github.com/tealeg/xlsx"
	"reflect"
	"strconv"
)

func Unmarshal(f *xlsx.File, v interface{}) error {
	for _, sheet := range f.Sheets {
		for _, row := range sheet.Rows {
			if nil == row {
				break
			}
			for k, val := range row.Cells {
				SetValue(k, val.Value, v)
			}
		}
	}
	return nil
}

func SetValue(i int, value string, v interface{}) {
	// TODO: interface judge
	tv := reflect.ValueOf(v)
	logrus.Infof("i: %v", i)
	if reflect.Indirect(tv).NumField() < i-1 {
		return
	}
	stv := reflect.Indirect(tv).Field(i)
	u, _, _ := indirect(stv, false)
	if u != nil {
		err := u.UnmarshalXlsx(value)
		logrus.Errorf("err: %v", err)
	}
	if stv.Kind() == reflect.Uint64 {
		xx, _ := strconv.ParseUint(value, 10, 64)
		stv.SetUint(xx)
	}
	if stv.Kind() == reflect.Int8 {
		xx, _ := strconv.ParseInt(value, 10, 8)
		stv.SetInt(xx)
	}
	if stv.Kind() == reflect.String {
		// handle time
		stv.SetString(value)
	}

}