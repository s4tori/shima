// Package util implements some utility functions.
package util

import (
	"fmt"
	"math"
	"math/rand"
	"regexp"
	"strings"
)

var (
	regHex  = regexp.MustCompile("^(?:[[:xdigit:]]{3}|[[:xdigit:]]{6})$")
	regRgb  = regexp.MustCompile("^[0-9]{1,3}-[0-9]{1,3}-[0-9]{1,3}$")
	regRbga = regexp.MustCompile("^[0-9]{1,3}-[0-9]{1,3}-[0-9]{1,3}-(?:0|1|0?.[0-9])$")
	emojis  = []string{"(-_-')", "(;-;)", "(&gt;_&lt;)", "(˚Δ˚)"}
)

// Ife evaluates a condition, if true returns the first parameter otherwise the second.
func Ife(cond bool, condTrue, condFalse int) int {
	if cond {
		return condTrue
	}
	return condFalse
}

// Min returns the smallest integer.
// See: https://mrekucci.blogspot.fr/2015/07/dont-abuse-mathmax-mathmin.html
func Min(x, y int) int {
	return Ife(x < y, x, y)
}

// Max returns the largest integer.
// See: https://mrekucci.blogspot.fr/2015/07/dont-abuse-mathmax-mathmin.html
func Max(x, y int) int {
	return Ife(x > y, x, y)
}

// Abs returns the absolute value of x.
func Abs(x int) int {
	return Ife(x < 0, -x, x)
}

// ReturnFirstError returns the first valid error (from left to right)
func ReturnFirstError(who ...error) error {
	for _, err := range who {
		if err != nil {
			return err
		}
	}

	return nil
}

// Matrix returns the transformation matrix relative to xAngle and yAngle.
func Matrix(xAngle, yAngle int) string {
	fXAngle := float64(xAngle) * math.Pi / 180
	fYAngle := float64(yAngle) * math.Pi / 180
	return fmt.Sprintf("matrix(1,%.3f,%.3f,1,0,0) scale(%.3f,%.3f)", math.Tan(fYAngle), math.Tan(fXAngle), math.Cos(fYAngle), math.Cos(fXAngle))
}

// GetIndex gets the value at a certain index of object.
// If the resolved value is undefined, def is returned in its place.
func GetIndex(values []string, i int, def string) string {
	if i >= 0 && i < len(values) {
		return values[i]
	}

	return def
}

// GetColor converts a raw "value" to a valid color string (hex, rgb or rgba).
func GetColor(value string) string {
	switch {
	case regHex.MatchString(value):
		return "#" + value
	case regRgb.MatchString(value):
		return "rgb(" + strings.Replace(value, "-", ",", 2) + ")"
	case regRbga.MatchString(value):
		return "rgba(" + strings.Replace(value, "-", ",", 3) + ")"
	}

	return value
}

// GetEmoji returns an emoji randomly
func GetEmoji() string {
	i := rand.Intn(len(emojis))
	return emojis[i]
}
