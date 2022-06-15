// Package grid outputs a valid SVG.
package grid

import (
	"encoding/base64"
	"fmt"
	"io"
	"log"

	"shima/internal/config"
	"shima/internal/svg"
	"shima/internal/util"
)

// RenderBase64 writes a valid SVG element into "w" using Base64 encodings.
func RenderBase64(w io.Writer, uri string) {
	fmt.Fprint(w, "data:image/svg+xml;base64,")
	b64Writer := base64.NewEncoder(base64.StdEncoding, w)
	Render(b64Writer, uri)
	b64Writer.Close()
}

// Render writes a valid SVG element into "w".
func Render(w io.Writer, uri string) {
	c := config.New()
	err := c.ParseFromUri(uri)
	log.Println(c)

	s := svg.New(w)
	s.Start()
	s.Open("defs")

	// Inner grid
	s.Pattern("innerGrid", c.XCell, c.YCell, "")
	s.GBlurry()
	if c.XShow {
		s.Path(fmt.Sprintf("M 0 0 L %d 0", c.XCell), c.XDivC)
	}
	if c.YShow {
		s.Path(fmt.Sprintf("M 0 0 L 0 %d", c.YCell), c.YDivC)
	}
	s.Close("g")
	s.Close("pattern")

	// Outer grid
	s.Pattern("outerGrid", c.XGridGutter, c.YGridGutter, c.GridTypeMatrix)
	s.Rect(util.Ife(c.YShow, 1, 0), util.Ife(c.XShow, 1, 0), c.XGrid+util.Ife(c.YShow, -1, 1), c.YGrid+util.Ife(c.XShow, -1, 1), "url(#innerGrid)")
	s.GBlurry()
	if c.XShow {
		s.Path(fmt.Sprintf("M 0 0 L %d 0", c.XGrid), c.XGridC)
	}
	if c.YShow {
		s.Path(fmt.Sprintf("M 0 0 L 0 %d", c.YGrid), c.YGridC)
	}
	if c.XGutterShow {
		s.Path(fmt.Sprintf("M 0 %d L %d %d", c.YGrid, c.XGrid, c.YGrid), c.XGridC)
	} // gutters
	if c.YGutterShow {
		s.Path(fmt.Sprintf("M %d 0 L %d %d", c.XGrid, c.XGrid, c.YGrid), c.YGridC)
	} // gutters
	s.Close("g")
	s.Close("pattern")

	s.Close("defs")
	s.RectFill("url(#outerGrid)")

	// Display errors inside the SVG
	if err != nil {
		s.TextError(err)
	}

	s.End()
}
