package ga_test

import (
	"bufio"
	"fmt"
	"net/http"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
	"testing"
	"ga"
)


// ==========================================================
// Public functions
// ----------------------------------------------------------

func TestGetCid(t *testing.T) {
	req := initReq("cookie1=value1;_ga=GA1.1.524171068.1511793623")
	assertEquals(t, ga.GetCid(req), "524171068.1511793623")

	req = initReq("_ga=GA1.2.3.4;cookie2=value2")
	assertEquals(t, ga.GetCid(req), "3.4")

	req = initReq("_ga=value1;cookie2=value2")
	assertNotEquals(t, ga.GetCid(req), "3.4")

	req = initReq("cookie1=GA1.2.3.4;cookie2=value2")
	assertNotEquals(t, ga.GetCid(req), "3.4")
}

func TestGetIp(t *testing.T) {
	req := initReqIp("", "1.2.3.4", "")
	assertEquals(t, ga.GetIp(req), "1.2.3.4")

	req = initReqIp("", "", "5.6.7.8:4321")
	assertEquals(t, ga.GetIp(req), "5.6.7.8")

	req = initReqIp("", "1.2.3.4", "5.6.7.8:4321")
	assertEquals(t, ga.GetIp(req), "1.2.3.4")

	req = initReqIp("", "", "")
	assertEquals(t, ga.GetIp(req), "")

	req = initReqIp("127.0.0.1", "1.2.3.4", "")
	assertEquals(t, ga.GetIp(req), "1.2.3.4")

	req = initReqIp("127.0.0.1,5.6.7.8", "1.2.3.4", "")
	assertEquals(t, ga.GetIp(req), "5.6.7.8")

	req = initReqIp("a.b.c.d, 5.6.7.8 , 127.0.0.1", "1.2.3.4", "9.10.11.12")
	assertEquals(t, ga.GetIp(req), "5.6.7.8")
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

// assertSvg fails the test if "got" is equal to "want".
func assertNotEquals(t *testing.T, got, want interface{}) {
	assert(t, reflect.DeepEqual(got, want) == false, " = '%v', want '%v'", got, want)
}

// initReq simulates an HTTP request withs cookies
func initReq(rawCookies string) *http.Request {
	rawRequest := fmt.Sprintf("GET / HTTP/1.0\r\nCookie: %s\r\n\r\n", rawCookies)
	req, _     := http.ReadRequest(bufio.NewReader(strings.NewReader(rawRequest)))

	return req
}

// initReq simulates an HTTP request with a specific IP
func initReqIp(xFrdFor, xReadlIp, remoteIp string) *http.Request {
	req := initReq("")

	if xFrdFor != "" {
		req.Header.Set("X-Forwarded-For", xFrdFor)
	}

	if xReadlIp != "" {
		req.Header.Set("X-Real-Ip", xReadlIp)
	}

	if remoteIp != "" {
		req.RemoteAddr = remoteIp
	}

	return req
}


// ==========================================================
