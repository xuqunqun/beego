package controllers

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMakeExcelByStructs(t *testing.T) {
	Convey("test MakeExcelByStructs", t, func() {
		tests := []struct {
			SheetName   string
			Expect int
		}{
			{"cyanxu",1},
			{"",0},
		}

		for _, tt := range tests {
			res := MakeExcelByStructs(tt.SheetName)
			So(res, ShouldEqual, tt.Expect)
		}
	})
}