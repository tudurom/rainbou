package util

import "fmt"

// Color stores the basic components of any color for the purpose of expressing in a multitude of formats.
type Color struct {
	R, G, B uint8
}

// Returns an HTML color without the leading #
func (c *Color) Hex() string {
	return fmt.Sprintf("%02x%02x%02x", c.R, c.G, c.B)
}

// Returns an HTML color with it components reversed
func (c *Color) HexBGR() string {
	return fmt.Sprintf("%02x%02x%02x", c.B, c.G, c.R)
}

// Returns an HTML-like color with its components doubled
func (c *Color) DHex() string {
	return fmt.Sprintf("%02x%02x%02x%02x%02x%02x", c.R, c.R, c.G, c.G, c.B, c.B)
}

// Returns an array of color color components
func (c *Color) RGB() []uint8 {
	return []uint8{
		c.R,
		c.G,
		c.B,
	}
}

// Returns an array of color components where their values are between 0 and 1
func (c *Color) SRGB() []float64 {
	return []float64{
		float64(c.R) / 255,
		float64(c.G) / 255,
		float64(c.B) / 255,
	}
}
