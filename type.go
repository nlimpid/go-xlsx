package go_xlsx

import (
	"strconv"
	"time"

	"github.com/tealeg/xlsx"
)

type xlsxTime struct {
	Time time.Time
}

func (x *xlsxTime) UnmarshalXlsx(data string) error {
	// Ignore null, like in the main JSON package.
	if string(data) == "null" || string(data) == "" {
		return nil
	}
	// Fractional seconds are handled implicitly by Parse.
	var err error
	t, err := strconv.ParseFloat(string(data), 10)
	x.Time = xlsx.TimeFromExcelTime(t, false)
	return err
}
