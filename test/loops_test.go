package test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLoops(t *testing.T) {
	Convey("while", t, func() {
		assertEval("x = 0; while (x < 3) { x++; } x;", intVal(3))
		assertEval("x = 0; while (x < 2) { x = x + 1; }", intVal(2))
		assertEval("x = 1; while (false) { x++; } x;", intVal(1))

		assertEval("x = 0; while (x < 2) { if (x == 1) { break; } x++; } x;", intVal(1))

		assertEval("a = 0; b = 0; while (a < 2) { a++; continue; } b;", intVal(0))
	})
}
