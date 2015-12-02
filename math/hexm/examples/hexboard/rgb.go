package main

import (
	"math"

	"github.com/gitchander/go-lang/cairo"
)

type RGB struct {
	R, G, B byte
}

func setSourceColor(c *cairo.Canvas, q RGB) {

	const t = 1.0 / math.MaxUint8

	c.SetSourceRGB(
		float64(q.R)*t,
		float64(q.G)*t,
		float64(q.B)*t,
	)
}

var (
	palettes = [][]RGB{
		[]RGB{
			RGB{255, 0, 0},
			RGB{0, 255, 0},
			RGB{0, 0, 255},
		},
		[]RGB{
			RGB{16, 16, 16},
			RGB{255, 255, 255},
			RGB{127, 127, 127},
		},
		[]RGB{
			RGB{3, 89, 146},
			RGB{211, 211, 211},
			RGB{100, 153, 188},
		},
		[]RGB{
			RGB{167, 79, 7},
			RGB{241, 208, 187},
			RGB{221, 165, 90},
		},
		[]RGB{
			RGB{70, 136, 202},
			RGB{216, 205, 84},
			RGB{247, 62, 162},
		},
	}
)
