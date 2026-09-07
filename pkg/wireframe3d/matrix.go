package wireframe3d

import "math"

type Point3D struct {
	X, Y, Z float64
}

type Point2D struct {
	X, Y int
}

func ProjectPoint(p Point3D, angleRad float64, scale float64) Point2D {
	cosA := math.Cos(angleRad)
	sinA := math.Sin(angleRad)
	rotX := p.X*cosA - p.Z*sinA
	rotZ := p.X*sinA + p.Z*cosA + 5.0
	projX := int((rotX / rotZ) * scale)
	projY := int((p.Y / rotZ) * scale)
	return Point2D{X: projX, Y: projY}
}
