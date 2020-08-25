// Package lengthconv performs Feet and Meter computations.
package lengthconv

import "fmt"

// Feet
type Feet float64

// Meter
type Meter float64

const (
	// MarathonM is length of full marathon in Meter
	MarathonM Meter = 42195
	// EverestM is height of the Everest in Meter
	EverestM Meter = 8848
)

// FToM converts a length in Feet to Meter.
func FToM(f Feet) Meter { return Meter(f * 0.3048) }

// CToK converts a length in Meter to Feet.
func MToF(m Meter) Feet { return Feet(m / 0.3048) }

func (f Feet) String() string  { return fmt.Sprintf("%gft", f) }
func (m Meter) String() string { return fmt.Sprintf("%gm", m) }
