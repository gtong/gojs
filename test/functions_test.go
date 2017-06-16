package test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestFunctions(t *testing.T) {
	Convey("functions", t, func() {
		assertEval("function() { return 1; }();", intVal(1))
		assertEval("a = function() { return 1; }; a();", intVal(1))
	})
}
