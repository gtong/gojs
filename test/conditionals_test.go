package test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestConditionals(t *testing.T) {
	Convey("if", t, func() {
		assertEval("if (true) { 1; }", intVal(1))
		assertEval("a = 1; if (a == 1) { a = a + 1; } a;", intVal(2))
		assertEval("a = 1; if (a != 1) { a = a + 1; } a;", intVal(1))
	})

	Convey("if else", t, func() {
		assertEval("if (true) { 1; } else { 2; }", intVal(1))
		assertEval("if (false) { 1; } else { 2; }", intVal(2))
	})
}
