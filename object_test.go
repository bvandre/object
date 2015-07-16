package object

import (
	"testing"
	"fmt"
)

var verts []*Vertex = []*Vertex{
	NewVertex(1, 0, 0, "x"),
	NewVertex(0, 1, 0, "y"),
	NewVertex(0, 0, 1, "z"),
}

var PiRotX []*Vertex = []*Vertex{
	XPiRotX,
	YPiRotX,
	ZPiRotX,
}

var Pi_4Rot []*Vertex = []*Vertex{
	X45Rot,
	Y45Rot,
	Z45Rot,
}

var (
	XPiRotX *Vertex = NewVertex(1, 0, 0, "x")
	X45Rot *Vertex = NewVertex(-1, 0, 0, "x")

	YPiRotX *Vertex = NewVertex(0, -1, 0, "y")
	Y45Rot *Vertex = NewVertex(0, 1, 0, "y")
	ZPiRotX *Vertex = NewVertex(0, 0, -1, "z")
	Z45Rot *Vertex = NewVertex(0, 0, -1, "z")
)

func TestObjectCreation(t *testing.T) {
	_, err := NewObject("fail", nil)
	if err == nil {
		t.Error("object should fail with nil vertex slice")
	}
	o, err := NewObject("test", verts)
	if err != nil {
		t.Errorf("error creating object: %s", err)
	}
	if len(o.Verts) != len(verts)+1 {
		t.Error("NewObject should create origin")
	}
}

func TestAbsoluteObjectRotation(t *testing.T) {
	o, _ := NewObject("test", verts)
	errPrefix := "object absolute rotation not correct"
	o.AbsRotate(180, 0, 0)
	vertexSliceTest(o.Verts, PiRotX, errPrefix, t)

	o.AbsRotate(45, 45, 45)
	fmt.Println(o.cp.quat)
	vertexSliceTest(o.Verts, Pi_4Rot, errPrefix, t)
}

func TestRelativeAbsoluteObjectRotation(t *testing.T) {
	o, _ := NewObject("test", verts)
	errPrefix := "object relative absolute rotation not correct"
	o.RelRotateAbsRef(90, 0, 0)
	o.RelRotateAbsRef(90, 0, 0)
	vertexSliceTest(o.Verts, PiRotX, errPrefix, t)
}

func vertexSliceTest(v1, v2 []*Vertex, errorPrefix string, t *testing.T) {
	for _, ver1 := range v1 {
		for _, ver2 := range v2 {
			if ver1.Name == ver2.Name {
				if !ver1.ApproxEqual(ver2) {
					t.Errorf("%s, should be %s, but is %s", errorPrefix, ver2, ver1)
				}
			}
		}
	}
}

func TestObjectTranslation(t *testing.T) {
	o, _ := NewObject("test", verts)
	errPrefix := "object not translated correctly"
	o.AbsTranslate(5, 0, 0)
	comp := []*Vertex{
		NewVertex(6, 0, 0, "x"),
		NewVertex(5, 1, 0, "y"),
		NewVertex(5, 0, 1, "z"),
	}
	vertexSliceTest(o.Verts, comp, errPrefix, t)
	o.RelTranslate(0, 5, 5)
	comp2 := []*Vertex{
		NewVertex(6, 5, 5, "x"),
		NewVertex(5, 6, 5, "y"),
		NewVertex(5, 5, 6, "z"),
	}
	vertexSliceTest(o.Verts, comp2, errPrefix, t)
}

func TestObjectSetPosition(t *testing.T) {
	o, _ := NewObject("test", verts)
	errPrefix := "object not translated correctly"
	tran := NewTransform(180, 0, 0, RotAbs, 1, 0, 0, TranAbs)
	p := NewPosition()
	tran.TransformPosition(p)
	o.SetPosition(p)
	comp := []*Vertex{
		NewVertex(2, 0, 0, "x"),
		NewVertex(1, -1, 0, "y"),
		NewVertex(1, 0, -1, "z"),
	}
	vertexSliceTest(o.Verts, comp, errPrefix, t)
	o.AbsRotate(5, 5, 5)
	o.AbsTranslate(4, 4, 4)
		o.SetPosition(p)
	comp2 := []*Vertex{
		NewVertex(2, 0, 0, "x"),
		NewVertex(1, -1, 0, "y"),
		NewVertex(1, 0, -1, "z"),
	}
	vertexSliceTest(o.Verts, comp2, errPrefix, t)
}
