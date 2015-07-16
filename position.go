package object

import (
	"fmt"
	"github.com/go-gl/mathgl/mgl32"
)

type Position struct {

	//quat is the objects current rotation
	quat  mgl32.Quat
	trans mgl32.Vec3
}

func NewPosition() *Position {
	return &Position{
		quat:  mgl32.QuatIdent(),
		trans: mgl32.Vec3{0, 0, 0},
	}
}

func PositionOfVertex(v *Vertex) *Position {
	t := mgl32.Vec3{-v.cvec.X(), -v.cvec.Y(), -v.cvec.Z()}
	return &Position{
		trans: t,
		quat:  v.obj.cp.quat.Inverse(),
	}
}

func PositionOfVertexNoRotation(v *Vertex) *Position {
	t := mgl32.Vec3{-v.cvec.X(), -v.cvec.Y(), -v.cvec.Z()}
	return &Position{
		trans: t,
		quat:  mgl32.QuatIdent(),
	}
}

func (p *Position) String() string {
	return fmt.Sprintf("Position Quat: w:%f, i:%f, j%f, k%f Translation x:%f, y%f, z%f",
		p.quat.W, p.quat.X(), p.quat.Y(), p.quat.Z(),
		p.trans.X(), p.trans.Y(), p.trans.Z())
}

func (p *Position) Copy() *Position {
	return &Position{
		quat:  p.quat,
		trans: p.trans,
	}
}

func (p *Position) Overwrite(p2 *Position) {
	p.quat = p2.quat
	p.trans = p2.trans
}
