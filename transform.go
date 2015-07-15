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
//translation
func NewTransform(x, y, z float32, rot RotationType, xt, yt, zt float32, tran TranslationType) *Transform {
	return &Transform{
		quat:  mgl32.AnglesToQuat(z, y, x, mgl32.ZYX),
		rType: rot,
		trans: mgl32.Vec3{xt, yt, zt},
		tType: tran,
	}
}

func (t *Transform) TransformPosition(p *Position) {
	switch t.rType {
	case RotAbs:
		p.Quat = t.quat
	case RotAddAbs:
		p.Quat = t.quat.Mul(p.Quat)
	case RotAddLocalRef:
		p.Quat = p.Quat.Mul(t.quat)
	}
	switch t.tType {
	case TranAbs:
		p.CurOffset = t.trans
	case TranRel:
		p.CurOffset = p.CurOffset.Add(t.trans)
	}
}
