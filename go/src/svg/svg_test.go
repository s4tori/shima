package svg_test

import (
	"bytes"
	"fmt"
	"path/filepath"
	"runtime"
	"testing"
	"svg"
)


// ==========================================================
// Public functions
// ----------------------------------------------------------

func TestEnd(t *testing.T) {
	s, b := initSvg()
	s.End()
	assertSvg(t, b, "</svg>\n")
}

func TestOpen(t *testing.T) {
	s, b := initSvg()
	s.Open("tag")
	assertSvg(t, b, "<tag>\n")

	b.Reset()
	s.Open("g")
	assertSvg(t, b, "<g>\n")
}

func TestClose(t *testing.T) {
	s, b := initSvg()
	s.Close("tag")
	assertSvg(t, b, "</tag>\n")

	b.Reset()
	s.Close("g")
	assertSvg(t, b, "</g>\n")
}

func TestPath(t *testing.T) {
	s, b := initSvg()
	s.Path("M 0 0 L 0 32", "red")
	assertSvg(t, b, fmt.Sprintf(`<path d="M 0 0 L 0 32" fill="none" stroke="red" stroke-width="1" />%s`, "\n"))
}

func TestRect(t *testing.T) {
	s, b := initSvg()
	s.Rect(1, 2, 3, 4, "blue")
	assertSvg(t, b, fmt.Sprintf(`<rect x="1" y="2" width="3" height="4" fill="blue" />%s`, "\n"))
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

// assertSvg fails the test if the buffer "b" isn't equal to "want"
func assertSvg(t *testing.T, b *bytes.Buffer, want string) {
	got := b.String()
	assert(t, got == want, "= %q, want %q", got, want)
}

// initSvg Initialize a SVG instance attached to a buffer
func initSvg() (*svg.SVG, *bytes.Buffer) {
	b := new(bytes.Buffer)
	s := svg.New(b)
	return s, b
}


// ==========================================================
