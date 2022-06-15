package config_test

import (
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"

	"shima/internal/config"
)

// ==========================================================
// Public functions
// ----------------------------------------------------------

func TestXAxis(t *testing.T) {
	c := initConfig("/64x4x32")
	assertAxis(t, c, "X", 64, 4, 32, config.DefGridC, config.DefDivC, true, true, false)
	assertAxis(t, c, "Y", 64, 4, 32, config.DefGridC, config.DefDivC, true, true, false)
	assertType(t, c, "standard")

	c = initConfig("/64x4x0x000")
	assertAxis(t, c, "X", 64, 4, 0, "#000", config.DefDivC, true, false, false)
	assertAxis(t, c, "Y", 64, 4, 0, "#000", config.DefDivC, true, false, false)
	assertType(t, c, "standard")

	c = initConfig("/32x2x12x0-0-0-.3xFFF")
	assertAxis(t, c, "X", 32, 2, 12, "rgba(0,0,0,.3)", "#FFF", true, true, false)
	assertAxis(t, c, "Y", 32, 2, 12, "rgba(0,0,0,.3)", "#FFF", true, true, false)
	assertType(t, c, "standard")

	c = initConfig("/1x2x-3x000xFFF")
	assertAxis(t, c, "X", 1, 2, 3, "#000", "#FFF", true, true, true)
	assertAxis(t, c, "Y", 1, 2, 3, "#000", "#FFF", true, true, true)
	assertType(t, c, "standard")
}

func TestYAxis(t *testing.T) {
	c := initConfig("/64x4x32/32x2x12")
	assertAxis(t, c, "X", 64, 4, 32, config.DefGridC, config.DefDivC, true, true, false)
	assertAxis(t, c, "Y", 32, 2, 12, config.DefGridC, config.DefDivC, true, true, false)
	assertType(t, c, "standard")

	c = initConfig("/64x4x0x000/32x2x12xFFF")
	assertAxis(t, c, "X", 64, 4, 0, "#000", config.DefDivC, true, true, false)
	assertAxis(t, c, "Y", 32, 2, 12, "#FFF", config.DefDivC, true, false, false)
	assertType(t, c, "standard")

	c = initConfig("/32x2x12x0-0-0-.3xFFF/64x4x32x000x0-0-0-.3")
	assertAxis(t, c, "X", 32, 2, 12, "rgba(0,0,0,.3)", "#FFF", true, true, false)
	assertAxis(t, c, "Y", 64, 4, 32, "#000", "rgba(0,0,0,.3)", true, true, false)
	assertType(t, c, "standard")

	c = initConfig("/1x2x3x000xFFF/1x2x-3x000xFFF")
	assertAxis(t, c, "X", 1, 2, 3, "#000", "#FFF", true, true, false)
	assertAxis(t, c, "Y", 1, 2, 3, "#000", "#FFF", true, true, true)
	assertType(t, c, "standard")
}

func TestType(t *testing.T) {
	c := initConfig("/64x4x32/isometric")
	assertAxis(t, c, "X", 64, 4, 32, config.DefGridC, config.DefDivC, true, true, false)
	assertAxis(t, c, "Y", 64, 4, 32, config.DefGridC, config.DefDivC, true, true, false)
	assertType(t, c, "isometric")

	c = initConfig("/64x4x32x000/32x2x0xFFF/isometric")
	assertAxis(t, c, "X", 64, 4, 32, "#000", config.DefDivC, true, false, false)
	assertAxis(t, c, "Y", 32, 2, 0, "#FFF", config.DefDivC, true, true, false)
	assertType(t, c, "isometric")
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

// assertSvg fails the test if "got" isn't equal to "want"
func assertEquals(t *testing.T, got, want interface{}) {
	assert(t, reflect.DeepEqual(got, want), "= '%v', want '%v'", got, want)
}

// assertAxis fails the test if the axis is incorrect
func assertAxis(t *testing.T, c interface{}, axis string, grid, div, gutter int, gridC, divC string, show, gutterShow, GutterBefore bool) {
	cElem := reflect.ValueOf(c).Elem()
	assertEquals(t, cElem.FieldByName(axis+"Grid").Int(), int64(grid))
	assertEquals(t, cElem.FieldByName(axis+"Div").Int(), int64(div))
	assertEquals(t, cElem.FieldByName(axis+"Gutter").Int(), int64(gutter))
	assertEquals(t, cElem.FieldByName(axis+"GridC").String(), gridC)
	assertEquals(t, cElem.FieldByName(axis+"DivC").String(), divC)
	assertEquals(t, cElem.FieldByName(axis+"Show").Bool(), show)
	assertEquals(t, cElem.FieldByName(axis+"GutterShow").Bool(), gutterShow)
	assertEquals(t, cElem.FieldByName(axis+"GutterBefore").Bool(), GutterBefore)
}

// assertType fails the test if the type is incorrect
func assertType(t *testing.T, c interface{}, gridType string) {
	cElem := reflect.ValueOf(c).Elem()
	assertEquals(t, cElem.FieldByName("GridType").String(), gridType)
}

// initConfig Initialize a Config instance based on the provided URI
func initConfig(path string) interface{} {
	c := config.New()
	c.ParseFromUri(path)
	return c
}

// ==========================================================
