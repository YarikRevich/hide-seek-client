package types

import (
	"github.com/kvartborg/vector"
)

type Vec2 struct{ X, Y float64 }

type Vec3 struct{ X, Y, Z float64 }

type Matrix4 [4][4]float64

func (m *Matrix4) GetTransposed() Matrix4 {
	var newMat Matrix4
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			newMat[i][j] = m[j][i]
		}
	}
	return newMat
}

func (m *Matrix4) GetForward() vector.Vector {
	vert := vector.Vector{
		m[2][0],
		m[2][1],
		m[2][2]}
	return vert.Unit()
}

func (m *Matrix4) GetMultipied(other Matrix4) Matrix4 {
	newMat := Matrix4{
		{1, 0, 0, 0},
		{0, 1, 0, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 1},
	}

	newMat[0][0] = m[0][0]*other[0][0] + m[0][1]*other[1][0] + m[0][2]*other[2][0] + m[0][3]*other[3][0]
	newMat[1][0] = m[1][0]*other[0][0] + m[1][1]*other[1][0] + m[1][2]*other[2][0] + m[1][3]*other[3][0]
	newMat[2][0] = m[2][0]*other[0][0] + m[2][1]*other[1][0] + m[2][2]*other[2][0] + m[2][3]*other[3][0]
	newMat[3][0] = m[3][0]*other[0][0] + m[3][1]*other[1][0] + m[3][2]*other[2][0] + m[3][3]*other[3][0]

	newMat[0][1] = m[0][0]*other[0][1] + m[0][1]*other[1][1] + m[0][2]*other[2][1] + m[0][3]*other[3][1]
	newMat[1][1] = m[1][0]*other[0][1] + m[1][1]*other[1][1] + m[1][2]*other[2][1] + m[1][3]*other[3][1]
	newMat[2][1] = m[2][0]*other[0][1] + m[2][1]*other[1][1] + m[2][2]*other[2][1] + m[2][3]*other[3][1]
	newMat[3][1] = m[3][0]*other[0][1] + m[3][1]*other[1][1] + m[3][2]*other[2][1] + m[3][3]*other[3][1]

	newMat[0][2] = m[0][0]*other[0][2] + m[0][1]*other[1][2] + m[0][2]*other[2][2] + m[0][3]*other[3][2]
	newMat[1][2] = m[1][0]*other[0][2] + m[1][1]*other[1][2] + m[1][2]*other[2][2] + m[1][3]*other[3][2]
	newMat[2][2] = m[2][0]*other[0][2] + m[2][1]*other[1][2] + m[2][2]*other[2][2] + m[2][3]*other[3][2]
	newMat[3][2] = m[3][0]*other[0][2] + m[3][1]*other[1][2] + m[3][2]*other[2][2] + m[3][3]*other[3][2]

	newMat[0][3] = m[0][0]*other[0][3] + m[0][1]*other[1][3] + m[0][2]*other[2][3] + m[0][3]*other[3][3]
	newMat[1][3] = m[1][0]*other[0][3] + m[1][1]*other[1][3] + m[1][2]*other[2][3] + m[1][3]*other[3][3]
	newMat[2][3] = m[2][0]*other[0][3] + m[2][1]*other[1][3] + m[2][2]*other[2][3] + m[2][3]*other[3][3]
	newMat[3][3] = m[3][0]*other[0][3] + m[3][1]*other[1][3] + m[3][2]*other[2][3] + m[3][3]*other[3][3]

	return newMat
}
