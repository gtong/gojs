package test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestUnaryOperators(t *testing.T) {
	Convey("increment", t, func() {
		assertEval("x = 0; x++;", intVal(0))
		assertEval("x = 0; x++; x;", intVal(1))

		assertEval("x = true; x++;", intVal(1))
		assertEval("x = true; x++; x;", intVal(2))

		assertEval("x = 'a'; x++;", nanVal())

		assertError("1++;", "Invalid left-hand side expression")
		assertError("'a'++;", "Invalid left-hand side expression")
		assertError("true++;", "Invalid left-hand side expression")
	})

	Convey("decrement", t, func() {
		assertEval("x = 3; x--;", intVal(3))
		assertEval("x = 3; x--; x;", intVal(2))
	})
}

func TestBinaryOperators(t *testing.T) {
	Convey("addition", t, func() {
		// Int tests
		assertEval("1 + 1;", intVal(2))
		assertEval("1 + 2 + 3;", intVal(6))
		assertEval("1 + '2';", strVal("12"))
		assertEval("1 + NaN;", nanVal())
		assertEval("1 + true;", intVal(2))
		assertEval("1 + false;", intVal(1))

		// NaN tests
		assertEval("NaN + 1;", nanVal())
		assertEval("NaN + '2';", strVal("NaN2"))
		assertEval("NaN + true;", nanVal())

		// String tests
		assertEval("'1' + 1;", strVal("11"))
		assertEval("'1' + 2 + 3;", strVal("123"))
		assertEval("'1' + NaN;", strVal("1NaN"))
		assertEval("'1' + true;", strVal("1true"))

		// Boolean tests
		assertEval("true + true;", intVal(2))
		assertEval("true + false;", intVal(1))
		assertEval("true + NaN;", nanVal())
		assertEval("true + '1';", strVal("true1"))
	})

	Convey("subtraction", t, func() {
		// Int tests
		assertEval("5 - 2;", intVal(3))
		assertEval("5 - 2 - 1;", intVal(2))
		assertEval("1 - 2;", intVal(-1))
		assertEval("3 - '1';", intVal(2))
		assertEval("1 - NaN;", nanVal())
		assertEval("3 - true;", intVal(2))
		assertEval("3 - 'a';", nanVal())
		assertEval("2 - NaN;", nanVal())

		// NaN tests
		assertEval("NaN - 1;", nanVal())

		// String tests
		assertEval("'a' - 1;", nanVal())

		// Boolean tests
		assertEval("true - true;", intVal(0))
		assertEval("false - 1;", intVal(-1))
	})

	Convey("assignment", t, func() {
		assertEval("x = 1;", intVal(1))
		assertEval("x = y = 1;", intVal(1))
		assertEval("x = y = z = 1;", intVal(1))

		assertEval("x = 1; y = x + 1;", intVal(2))

		assertEval("x = 1; y = x; x = 2; y;", intVal(1))
		assertEval("x = 1; y = x; x = 2; y + 0;", intVal(1))

		assertEval("x = true;", boolVal(true))

		assertError("1 = 1;", "Invalid left-hand side in assignment")
		assertError("'a' = 'a';", "Invalid left-hand side in assignment")
		assertError("true = true;", "Invalid left-hand side in assignment")
	})

	Convey("equality", t, func() {
		// Int tests
		assertEval("1 == 1;", boolVal(true))
		assertEval("1 == 2;", boolVal(false))
		assertEval("1 == '1';", boolVal(true))
		assertEval("1 == true;", boolVal(true))
		assertEval("1 + 1 == 3 - 1;", boolVal(true))

		// String tests
		assertEval("'dog' == 'dog';", boolVal(true))
		assertEval("'dog' == 'cat';", boolVal(false))

		// Bool tests
		assertEval("true == true;", boolVal(true))
		assertEval("true == false;", boolVal(false))
		assertEval("true == 'true';", boolVal(false))
		assertEval("true == 1;", boolVal(true))
		assertEval("true == '1';", boolVal(true))

		// NaN tests
		assertEval("NaN == NaN;", boolVal(false))
		assertEval("NaN == 'hi';", boolVal(false))

		// Variables
		assertEval("x = 1; x == '1';", boolVal(true))

		// Strict
		assertEval("true === true;", boolVal(true))
		assertEval("true === 1;", boolVal(false))
		assertEval("x = 'a'; x === 'a';", boolVal(true))
	})

	Convey("inequality", t, func() {
		// Int tests
		assertEval("1 != 2;", boolVal(true))
		assertEval("1 != 1;", boolVal(false))
		assertEval("1 != '1';", boolVal(false))

		// String tests
		assertEval("'dog' != 'dog';", boolVal(false))
		assertEval("'dog' != 'cat';", boolVal(true))

		// Bool tests
		assertEval("true != true;", boolVal(false))
		assertEval("true != false;", boolVal(true))

		// NaN tests
		assertEval("NaN != NaN;", boolVal(true))

		// Strict
		assertEval("1 !== 2;", boolVal(true))
		assertEval("NaN !== NaN;", boolVal(true))
		assertEval("1 !== '1';", boolVal(true))
	})

	Convey("compare", t, func() {
		// Greater Than tests
		assertEval("true > false;", boolVal(true))
		assertEval("true > true;", boolVal(false))
		assertEval("2 > 1;", boolVal(true))
		assertEval("2 + 2 > 2 + 1;", boolVal(true))
		assertEval("y = 2; x = 1; y > x;", boolVal(true))
		assertEval("y = 2; x = 1; x > y;", boolVal(false))
		assertEval("y = 2; x = 1; x + 2 > y;", boolVal(true))
		assertEval("'z' > 'a';", boolVal(true))
		assertEval("'a' > 'Z';", boolVal(true))
		assertEval("'A' > '9';", boolVal(true))
		assertEval("'9' > '-';", boolVal(true))
		assertEval("'a' + 1 > 'a';", boolVal(true))

		// Less Than tests
		assertEval("true < false;", boolVal(false))
		assertEval("1 < 2;", boolVal(true))
		assertEval("'abc' < 'abd';", boolVal(true))

		// Greater Than or Equal tests
		assertEval("3 >= 2;", boolVal(true))
		assertEval("2 >= 2;", boolVal(true))
		assertEval("'a' >= 'a';", boolVal(true))

		// Less Than or Equal tests
		assertEval("1 <= 2;", boolVal(true))
		assertEval("2 <= 2;", boolVal(true))
		assertEval("'a' <= 'a';", boolVal(true))
	})
}
