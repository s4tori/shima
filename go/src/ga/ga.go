// Package ga manages Google Analytics.
package ga

import (
	"log"
	"math/rand"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

const (
	cidMin  = 100000000
	cidMax  = 999999999
	gaUrl   = "https://www.google-analytics.com/collect"
	trackId = "UA-110169890-1"
)


// ==========================================================
// Public functions
// ----------------------------------------------------------

// GetCid retrieves "cid" (client ID) from "_ga" cookie.
// If the cookie does not exist, a random CID is generated.
// Example: GA1.1.524171068.1511793622 => 524171068.1511793622.
func GetCid(r *http.Request) string {
	var cookie, err = r.Cookie("_ga")
	if err != nil {
		return generateCid()
	}

	var cookievalue = cookie.Value
	var opt = strings.Split(cookievalue, ".")

	if len(opt) < 4 {
		return generateCid()
	}

	return opt[2] + "." + opt[3]
}

// Get the IP of the current user. Useful only to report correct geolocation.
// The IP address of the sender will be anonymized.
// See https://developers.google.com/analytics/devguides/collection/protocol/v1/parameters#uip
func GetIp(r *http.Request) string {
	hdrRealIP := r.Header.Get("X-Real-Ip")
	hdrFwdIP  := r.Header.Get("X-Forwarded-For")

	if len(hdrFwdIP) == 0 && len(hdrRealIP) == 0 {
		ip, _, _ := net.SplitHostPort(r.RemoteAddr)
		return ip
	}

	// X-Forwarded-For is potentially a list of addresses separated with ","
	for _, addr := range strings.Split(hdrFwdIP, ",") {
		fIp := net.ParseIP(strings.TrimSpace(addr))
		if fIp != nil && !fIp.IsLoopback() {
			return fIp.String()
		}
	}

	return hdrRealIP
}

// TrackPageEvent sends data to GA using the Measurement Protocol.
func TrackPageEvent(r *http.Request) {
	values := url.Values{}
	values.Set("v"  , "1")
	values.Add("t"  , "pageview")
	values.Add("aip", "1")
	values.Add("tid", trackId)
	values.Add("uip", GetIp(r))
	values.Add("cid", GetCid(r))
	values.Add("dp" , r.URL.Path)

	go gaRequest(values, r.UserAgent())
}


// ==========================================================
// Private functions
// ----------------------------------------------------------

// generateCid generates a random CID.
func generateCid() string {
	return strconv.Itoa(rand.Intn(cidMax - cidMin) + cidMin)
}

// gaRequest creates an HTTP request to GA.
func gaRequest(values url.Values, userAgent string) {
	req, err := http.NewRequest("POST", gaUrl, strings.NewReader(values.Encode()))
	if err != nil {
		log.Println(err)
		return
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("User-Agent"  , userAgent)
	_, err = http.DefaultClient.Do(req)

	if err != nil {
		log.Println(err)
	}
}


// ==========================================================
