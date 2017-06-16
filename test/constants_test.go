package test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestConstants(t *testing.T) {
	Convey("numbers", t, func() {
		assertEval("1;", intVal(1))
	})

	Convey("NaN", t, func() {
		assertEval("NaN;", nanVal())
	})

	Convey("strings", t, func() {
		assertEval("'1';", strVal("1"))
	})

	Convey("booleans", t, func() {
		assertEval("true;", boolVal(true))
		assertEval("false;", boolVal(false))
	})

	Convey("variables", t, func() {
		assertEval("x = 1; x;", intVal(1))
		assertError("x;", "x is not defined")
	})
}
