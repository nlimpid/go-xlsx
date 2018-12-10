package xlsx_parser

import (
	"github.com/sirupsen/logrus"
	"github.com/tealeg/xlsx"
	"strconv"
	"time"
)

type xlsxTime struct {
	Time time.Time
}

func (x *xlsxTime) UnmarshalXlsx(data string) error {
	// Ignore null, like in the main JSON package.
	if string(data) == "null" {
		return nil
	}
	// Fractional seconds are handled implicitly by Parse.
	var err error
	t, err := strconv.ParseFloat(string(data), 10)
	x.Time = xlsx.TimeFromExcelTime(t, true)
	logrus.Infof("x.Time: %v", x.Time)
	return err
}