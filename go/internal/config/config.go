// Package config manages a configuration object.
package config

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"shima/internal/util"
)

// config defines the configuration to generate the SVG.
type config struct {
	XGrid, XDiv, XGutter              int
	YGrid, YDiv, YGutter              int
	XGridC, XDivC                     string
	YGridC, YDivC                     string
	GridType, GridTypeMatrix          string
	XShow, XGutterShow, XGutterBefore bool
	YShow, YGutterShow, YGutterBefore bool

	XCell, XGridGutter int
	YCell, YGridGutter int
}

// types defines a list of transformation available for the grid.
var types = map[string]string{
	"standard":        "",
	"oblique":         util.Matrix(-45, 45),
	"isometric":       util.Matrix(-60, 30),
	"2-1-isometric":   util.Matrix(-63, 27),
	"trimetric-left":  util.Matrix(35, -9),
	"trimetric-right": util.Matrix(-25, 9),
}

const (
	DefGridC = "rgba(0,0,0,.2)"
	DefDivC  = "rgba(0,0,0,.1)"
)

// Regex to validate the URI.
// validPath: /<validXAxis>/<validYAxis>/<GridType>
// validAxis: <grid>x<div>x<gutter>x<gridColor>x<gridDiv>
//            <gridColor> and <gridDiv> are optional
// Ex (validPath):
// * /64x4x32xF2Fx2F2/
// * /64x4x32xF2Fx2F2/64x4x32x2F2xF2F/
// * /64x4x32xF2Fx2F2/isometric
// * /64x4x32xF2Fx2F2/64x4x32x2F2xF2F/isometric
// Ex (validAxis):
// * 64x4x32
// * 64x4x32x#FF0011
// * 64x4x32x#FF0011x#0000FF
// * 64x4x32x#255-255-200x#0000FF
// * 64x4x32x#255-255-200-0.5x#0000FF
var (
	validAxis = regexp.MustCompile("^([0-9]+)x([0-9]+)x(-)?([0-9]+)(?:x([[:xdigit:]\\-.]+))?(?:x([[:xdigit:]\\-.]+))?$")
	validType = regexp.MustCompile("^([a-z\\-]+)$")
)

// ==========================================================
// Public functions
// ----------------------------------------------------------

// New defines the config constructor.
func New() *config {
	c := config{
		XGrid: 160, XDiv: 5, XGutter: 32, XShow: true, XGutterShow: true,
		YGrid: 160, YDiv: 5, YGutter: 32, YShow: true, YGutterShow: true,
		XGridC: DefGridC, XDivC: DefDivC,
		YGridC: DefGridC, YDivC: DefDivC,
	}

	c.setGridType("standard")
	c.EndConfig()
	return &c
}

// EndConfig calculates computed variables.
func (c *config) EndConfig() {
	if c.XGrid == 0 && c.XGutter == 0 {
		c.YShow = false
	}
	if c.YGrid == 0 && c.YGutter == 0 {
		c.XShow = false
	}
	if c.XGutter == 0 || !c.YShow {
		c.YGutterShow = false
	}
	if c.YGutter == 0 || !c.XShow {
		c.XGutterShow = false
	}
	c.XGrid = util.Ife(c.XGrid == 0, c.YGrid, c.XGrid)
	c.YGrid = util.Ife(c.YGrid == 0, c.XGrid, c.YGrid)
	c.XCell = c.XGrid / util.Max(1, c.XDiv)
	c.YCell = c.YGrid / util.Max(1, c.YDiv)
	c.XGridGutter = c.XGrid + c.XGutter
	c.YGridGutter = c.YGrid + c.YGutter
}

// ParseFromUri creates a new configuration objet based on the provided URI.
func (c *config) ParseFromUri(uri string) error {
	var errXAxis, errYAxis, errGType error
	var skip bool

	opt := strings.Split(strings.Trim(uri, "/"), "/")
	oLen := len(opt)

	// A complete URI needs 3 items:
	// * The first item is always the X axis
	// * The second item can be the Y axis or the type of grid
	// * The third item is always the type of grid
	switch {
	case oLen == 1 && opt[0] == "":
		skip = true
	case oLen == 1:
		opt = append(opt, opt[0], c.GridType)
	case oLen == 2 && c.isGridType(opt[1]):
		opt = append([]string{opt[0]}, opt...)
	case oLen == 2:
		opt = append(opt, c.GridType)
	}

	if !skip {
		errXAxis = c.parseAxisFromUri(opt[0], "X")
		errYAxis = c.parseAxisFromUri(opt[1], "Y")
		errGType = c.setGridType(opt[2])
		c.EndConfig()
	}

	return util.ReturnFirstError(errXAxis, errYAxis, errGType)
}

// String displays a config struct.
func (c *config) String() string {
	return fmt.Sprintf("\n[Config]\n"+
		"*  x: %3d, %2d, %2d (xCell: %2d)\n"+
		"*  y: %3d, %2d, %2d (yCell: %2d)\n"+
		"* bX: XShow:%5t, XGutterShow:%5t, XGutterBefore:%5t\n"+
		"* bY: YShow:%5t, YGutterShow:%5t, YGutterBefore:%5t\n"+
		"* cX: %s, %s\n"+
		"* cY: %s, %s\n"+
		"*  t: %s, %s\n",
		c.XGrid, c.XDiv, c.XGutter, c.XCell,
		c.YGrid, c.YDiv, c.YGutter, c.YCell,
		c.XShow, c.XGutterShow, c.XGutterBefore,
		c.YShow, c.YGutterShow, c.YGutterBefore,
		c.XGridC, c.XDivC,
		c.YGridC, c.YDivC,
		c.GridType, c.GridTypeMatrix,
	)
}

// ==========================================================
// Private functions
// ----------------------------------------------------------

// ParseFromUri extracts configuration for an axis based on the provided substring URI.
func (c *config) parseAxisFromUri(uri, kAxis string) error {
	m := validAxis.FindStringSubmatch(uri)
	if m == nil {
		return errors.New("Invalid " + kAxis + " axis")
	}

	c.setValue(kAxis+"Grid", m[1])
	c.setValue(kAxis+"Div", m[2])
	c.setValue(kAxis+"Gutter", m[4])
	c.setValue(kAxis+"GutterBefore", m[3] == "-")
	c.setValue(kAxis+"GridC", util.GetColor(util.GetIndex(m, 5, "")))
	c.setValue(kAxis+"DivC", util.GetColor(util.GetIndex(m, 6, "")))

	return nil
}

// setValue set a value inside a specific field of a *config struct.
func (c *config) setValue(field string, value interface{}) {
	v := reflect.ValueOf(c).Elem().FieldByName(field)
	if !v.IsValid() {
		return
	}

	switch {
	case v.Kind() == reflect.Int:
		//iValue, _ := strconv.Atoi(value)
		//v.SetInt(int64(iValue))
		iValue, _ := strconv.ParseInt(value.(string), 10, 64)
		v.SetInt(iValue)
	case v.Kind() == reflect.Bool:
		v.SetBool(value.(bool))
	case value != "":
		v.SetString(value.(string))
	}
}

// isGridType checks if gridType is a valid type.
func (c *config) isGridType(gridType string) bool {
	_, ok := types[gridType]
	return ok
}

// setGridType assigns a type to the grid.
func (c *config) setGridType(gridType string) error {
	_, ok := types[gridType]

	if !ok {
		gridType = "standard"
	}

	c.GridType = gridType
	c.GridTypeMatrix = types[gridType]

	if c.XGutterBefore || c.YGutterBefore {
		xGutterBefore := util.Ife(c.XGutterBefore, c.XGutter, 0)
		yGutterBefore := util.Ife(c.YGutterBefore, c.YGutter, 0)
		c.GridTypeMatrix += fmt.Sprintf("translate(%d,%d)", xGutterBefore, yGutterBefore)
	}

	if !ok {
		return errors.New("invalid grid type")
	}

	return nil
}

// ==========================================================
