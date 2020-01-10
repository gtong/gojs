package test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestFunctions(t *testing.T) {

	Convey("functions", t, func() {
		assertEval("function() { return 1; }();", intVal(1))
		assertEval("a = function() { return 1; }; a();", intVal(1))

		assertEval("a = function() { a = 1; return a; }; a();", intVal(1))
		assertEval("a = function() { b = 1; b++; return b; }; a();", intVal(2))

		assertEval("function() { return 2; return 1; }();", intVal(2))
	})

	Convey("functions with arguments", t, func() {
		assertEval("function(a) { return a; }(1);", intVal(1))
		assertEval("function(a, b) { return a + b; }(1, 2);", intVal(3))
		assertEval("function(a) { return a == null; }();", boolVal(true))
		assertEval("function(a) { return a; }(1 + 1);", intVal(2))
		assertEval("function(a) { return a; }(1 + 1, 3);", intVal(2))
		assertEval("f = function(a) { return a + 1; }; function(a) { return a + 2; }( f(1) );", intVal(4))
	})

	Convey("closures", t, func() {
		assertEval("a = 1; f = function() { return a; }; f();", intVal(1))
		assertEval("a = 1; f = function() { return a++; }; f(); a;", intVal(2))
		assertEval("f = function() { a = 1; return function() { return a++; }; }(); f(); f();", intVal(2))
		assertEval("f = function() { a = 1; b = function() { return a; }; a = 3; return b; }(); f();", intVal(3))
	})

}
