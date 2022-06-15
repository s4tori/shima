// Package svg helps to generate a SVG.
// Inspiration: https://github.com/ajstarks/svgo
package svg

import (
	"fmt"
	"io"

	"shima/internal/util"
)

// SVG defines the generated SVG.
type SVG struct {
	Writer io.Writer
}

const (
	ff = "system-ui,-apple-system,BlinkMacSystemFont,'Segoe UI',Roboto,Oxygen,Ubuntu,Cantarell,'Open Sans','Helvetica Neue',sans-serif"
)

// ==========================================================
// Public functions
// ----------------------------------------------------------

// New defines the SVG constructor.
func New(w io.Writer) *SVG {
	return &SVG{w}
}

// Start begins the SVG document.
func (svg *SVG) Start() {
	svg.println(`<?xml version="1.0"?>`)
	svg.println(`<svg width="100%" height="100%" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink">`)
}

// End ends the SVG document.
func (svg *SVG) End() {
	svg.println("</svg>")
}

// Open opens a tag.
func (svg *SVG) Open(tag string) {
	svg.printf("<%s>\n", tag)
}

// Close closes a tag.
func (svg *SVG) Close(tag string) {
	svg.printf("</%s>\n", tag)
}

// Pattern defines a pattern.
func (svg *SVG) Pattern(id string, w, h int, transform string) {
	svg.printf(`<pattern id="%s" x="0" y="0" width="%d" height="%d" patternUnits="userSpaceOnUse"`, id, w, h)
	if transform != "" {
		svg.printf(` patternTransform="%s"`, transform)
	}
	svg.println(">")
}

// GBlurry begins a group with a translation of 0.5.
func (svg *SVG) GBlurry() {
	svg.println(`<g transform="translate(0.5, 0.5)">`)
}

// Path draws an arbitrary path.
func (svg *SVG) Path(path, stroke string) {
	svg.printf(`<path d="%s" fill="none" stroke="%s" stroke-width="1" />%s`, path, stroke, "\n")
}

// Rect draws a rectangle.
func (svg *SVG) Rect(x, y, w, h int, fill string) {
	svg.printf(`<rect x="%d" y="%d" width="%d" height="%d" fill="%s" />%s`, x, y, w, h, fill, "\n")
}

// RectFill draws a rectangle width 100% width and height.
func (svg *SVG) RectFill(fill string) {
	svg.printf(`<rect x="0" y="0" width="100%%" height="100%%" fill="%s" />%s`, fill, "\n")
}

// TextError defines a text error (at the center of the SVG).
func (svg *SVG) TextError(err error) {
	emoji := util.GetEmoji()
	svg.RectFill("rgba(0,0,0,.8)")
	svg.printf(`<text x="50%%" y="50%%" fill="#f03" font-family="%s" font-size="30px" text-anchor="middle">Error %s</text>%s`, ff, emoji, "\n")
	svg.printf(`<text x="50%%" y="50%%" fill="#f03" font-family="%s" font-size="20px" transform="translate(0, 30)" text-anchor="middle">%s</text>%s`, ff, err.Error(), "\n")
}

// ==========================================================
// Private functions
// ----------------------------------------------------------

// print calls "print" using the SVG writer.
func (svg *SVG) print(a ...interface{}) (n int, err error) {
	return fmt.Fprint(svg.Writer, a...)
}

// println calls "println" using the SVG writer.
func (svg *SVG) println(a ...interface{}) (n int, err error) {
	return fmt.Fprintln(svg.Writer, a...)
}

// printf calls "printf" using the SVG writer.
func (svg *SVG) printf(format string, a ...interface{}) (n int, err error) {
	return fmt.Fprintf(svg.Writer, format, a...)
}

// ==========================================================
