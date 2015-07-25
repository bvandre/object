package object

import (
	"github.com/go-gl/mathgl/mgl32"
)

type Transform struct {
	//quat is the rotation defined
	//in this orientation
	quat  mgl32.Quat
	rType RotationType
	//trans is the translation defined
	//in this orientation
	trans mgl32.Vec3
	tType TranslationType
}

var (
	xAxis mgl32.Vec3 = mgl32.Vec3{1, 0, 0}
	yAxis mgl32.Vec3 = mgl32.Vec3{0, 1, 0}
	zAxis mgl32.Vec3 = mgl32.Vec3{0, 0, 1}
)

type RotationType int

const (
	RotNone RotationType = iota
	RotAbs
	RotAddAbs
	RotAddLocalRef
)

type TranslationType int

const (
	TranNone TranslationType = iota
	TranAbs
	TranRel
)

//NewTransform defines a new Rotation and Translation
//x, y, z are the angle in degrees rotation around the repsective axis
//addition is true when this will be applied in addition to an existing orientation
//localref is true if the orientation should be applied in the object local reference
//or in the absolute reference.  If addition is false this has no effect
//xt, yt, zt are the translation offset to apply in the x, y, and z axis
//transabs is true when this translation will be applied in relation to the current
//translation ZXY YZX XYZ
func NewTransform(x, y, z float32, rot RotationType, xt, yt, zt float32, tran TranslationType) *Transform {
	return &Transform{
		quat:  mgl32.AnglesToQuat(mgl32.DegToRad(z), mgl32.DegToRad(y), mgl32.DegToRad(x), mgl32.ZYX),
		rType: rot,
		trans: mgl32.Vec3{xt, yt, zt},
		tType: tran,
	}
}

//Updates a Position to the transform represented by t
func (t *Transform) TransformPosition(p *Position) {
	switch t.rType {
	case RotAbs:
		p.quat = t.quat
	case RotAddAbs:
		p.quat = t.quat.Mul(p.quat)
	case RotAddLocalRef:
		p.quat = p.quat.Mul(t.quat)
	}
	switch t.tType {
	case TranAbs:
		p.trans = t.trans
	case TranRel:
		p.trans = p.trans.Add(t.trans)
	}
}
