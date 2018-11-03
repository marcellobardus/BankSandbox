package utils

import (
	"math"

	"gopkg.in/fogleman/gg.v1"
)

type bank struct {
	connections []int
	x           float64
	y           float64
}

func GetGraphVisulisation(width uint16, height uint16) {

	var banks = []bank{
		bank{connections: []int{1, 2}},
		bank{connections: []int{0}},
		bank{connections: []int{0, 3}},
		bank{connections: []int{2, 1}},
	}
	dc := gg.NewContext(1000, 1000)

	dc.DrawPoint(float64(width/2), float64(height/2), 5)
	dc.SetRGB(255, 255, 255)

	dc.DrawLine(10, 20, 400, 700)

	r := 400

	for i := 0; i < len(banks); i++ {

		circlePercentage := float64(i) / float64(len(banks))
		alpha := circlePercentage * 2 * math.Pi
		x := float64(r) * math.Cos(alpha)
		y := float64(r) * math.Sin(alpha)

		banks[i].x = x + float64(width/2)
		banks[i].y = y + float64(height/2)

		dc.DrawPoint(banks[i].x, banks[i].y, 5)

	}

	for i := 0; i < len(banks); i++ {
		for j := i; j < len(banks[i].connections); j++ {
			dc.DrawLine(banks[i].x, banks[i].y, banks[banks[i].connections[j]].x, banks[banks[i].connections[j]].y)
		}
	}

	dc.Fill()
	dc.SavePNG("out.png")
}
