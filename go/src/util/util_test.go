package util_test

import (
	"errors"
	"fmt"
	"reflect"
	"path/filepath"
	"runtime"
	"testing"
	"util"
)


// ==========================================================
// Public functions
// ----------------------------------------------------------

func TestIfe(t *testing.T) {
	assertEquals(t, util.Ife(true , 1, 2), 1)
	assertEquals(t, util.Ife(true , 2, 1), 2)
	assertEquals(t, util.Ife(false, 1, 2), 2)
	assertEquals(t, util.Ife(false, 2, 1), 1)
}

func TestMin(t *testing.T) {
	assertEquals(t, util.Min( 1, 2),  1)
	assertEquals(t, util.Min( 2, 1),  1)
	assertEquals(t, util.Min(-2, 1), -2)
	assertEquals(t, util.Min(1, -2), -2)
}

func TestMax(t *testing.T) {
	assertEquals(t, util.Max( 1, 2), 2)
	assertEquals(t, util.Max( 2, 1), 2)
	assertEquals(t, util.Max(-2, 1), 1)
	assertEquals(t, util.Max(1, -2), 1)
}

func TestAbs(t *testing.T) {
	assertEquals(t, util.Abs( 1), 1)
	assertEquals(t, util.Abs( 2), 2)
	assertEquals(t, util.Abs(-1), 1)
	assertEquals(t, util.Abs(-2), 2)
	assertEquals(t, util.Abs(0) , 0)
}

func TestReturnFirstError(t *testing.T) {
	error1 := errors.New("error 1")
	error2 := errors.New("error 2")
	assertEquals(t, util.ReturnFirstError(error1, error2), error1)
	assertEquals(t, util.ReturnFirstError(error2, error1), error2)
	assertEquals(t, util.ReturnFirstError(nil   , error2), error2)
	assertEquals(t, util.ReturnFirstError(error1, nil   ), error1)
	assertEquals(t, util.ReturnFirstError(nil   , nil   ), nil)
}

func TestGetIndex(t *testing.T) {
	assertEquals(t, util.GetIndex([]string{"a","b"}, -1, "c"), "c")
	assertEquals(t, util.GetIndex([]string{"a","b"},  0, "c"), "a")
	assertEquals(t, util.GetIndex([]string{"a","b"},  1, "c"), "b")
	assertEquals(t, util.GetIndex([]string{"a","b"},  2, "c"), "c")
}

func TestGetColor(t *testing.T) {
	assertEquals(t, util.GetColor("0F3")         , "#0F3")
	assertEquals(t, util.GetColor("0FFF33")      , "#0FFF33")
	assertEquals(t, util.GetColor("0-255-51")    , "rgb(0,255,51)")
	assertEquals(t, util.GetColor("0-255-51-0")  , "rgba(0,255,51,0)")
	assertEquals(t, util.GetColor("0-255-51-1")  , "rgba(0,255,51,1)")
	assertEquals(t, util.GetColor("0-255-51-0.5"), "rgba(0,255,51,0.5)")
	assertEquals(t, util.GetColor("0-255-51-.5") , "rgba(0,255,51,.5)")

	// Invalid colors
	assertEquals(t, util.GetColor("0-AAA-FFF")   , "0-AAA-FFF")
	assertEquals(t, util.GetColor("0-A-F")       , "0-A-F")
	assertEquals(t, util.GetColor("foobar")      , "foobar")
}


// ==========================================================
// Private functions
// ----------------------------------------------------------

// assert fails the test if the condition is false.
// See: https://github.com/benbjohnson/testing
func assert(t *testing.T, condition bool, msg string, v ...interface{}) {
	if !condition {
		pc, file, line, _ := runtime.Caller(2)
		method := runtime.FuncForPC(pc).Name()
		fmt.Printf("\033[31m%s:%d: %s() "+msg+"\033[39m\n\n", append([]interface{}{filepath.Base(file), line, method}, v...)...)
		t.FailNow()
	}
}

// assertSvg fails the test if "got" isn't equal to "want".
func assertEquals(t *testing.T, got, want interface{}) {
	assert(t, reflect.DeepEqual(got, want), " = '%v', want '%v'", got, want)
}


// ==========================================================
// Example functions
// ----------------------------------------------------------

func ExampleIfe_output() {
	fmt.Println(util.Ife(true , 1, 2))
	fmt.Println(util.Ife(false, 1, 2))
	// Output:
	// 1
	// 2
}

func ExampleMin_output() {
	fmt.Println(util.Min(1, 2))
	fmt.Println(util.Min(-2, 1))
	// Output:
	// 1
	// -2
}

func ExampleMax_output() {
	fmt.Println(util.Max(2, 1))
	fmt.Println(util.Max(-2, 1))
	// Output:
	// 2
	// 1
}

func ExampleAbs_output() {
	fmt.Println(util.Abs(1))
	fmt.Println(util.Abs(-2))
	// Output:
	// 1
	// 2
}

func ExampleReturnFirstError_output() {
	error1 := errors.New("error 1")
	error2 := errors.New("error 2")
	fmt.Println(util.ReturnFirstError(error1, error2))
	fmt.Println(util.ReturnFirstError(nil   , error2))
	// Output:
	// error 1
	// error 2
}

func ExampleGetColor_output() {
	fmt.Println(util.GetColor("0F3"))
	fmt.Println(util.GetColor("0FFF33"))
	fmt.Println(util.GetColor("0-255-51"))
	fmt.Println(util.GetColor("0-255-51-.5"))
	// Output:
	// #0F3
	// #0FFF33
	// rgb(0,255,51)
	// rgba(0,255,51,.5)
}


// ==========================================================
