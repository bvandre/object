package object

import (
	"math"
	"testing"
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

var PiRotY []*Vertex = []*Vertex{
	XPiRotY,
	YPiRotY,
	ZPiRotY,
}

var (
	XPiRotX *Vertex = NewVertex(1, 0, 0, "x")
	XPiRotY *Vertex = NewVertex(-1, 0, 0, "x")

	YPiRotX *Vertex = NewVertex(0, -1, 0, "y")
	YPiRotY *Vertex = NewVertex(0, 1, 0, "y")
	ZPiRotX *Vertex = NewVertex(0, 0, -1, "z")
	ZPiRotY *Vertex = NewVertex(0, 0, -1, "z")
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

func TestObjectRotation(t *testing.T) {
	o, _ := NewObject("test", verts)
	o.AbsRotate(math.Pi, 0, 0)
	for _, v1 := range o.Verts {
		for _, v2 := range PiRotX {
			if v1.Name == v2.Name {
				if !v1.ApproxEqual(v2) {
					t.Errorf("vertex not correctly rotated, should be %s, but is %s", v2, v1)
				}
			}
		}
	}
	o.AbsRotate(0, math.Pi, 0)
	for _, v1 := range o.Verts {
		for _, v2 := range PiRotY {
			if v1.Name == v2.Name {
				if !v1.ApproxEqual(v2) {
					t.Errorf("vertex not correctly rotated, should be %s, but is %s", v2, v1)
				}
			}
		}
	}

}

func TestObjectTranslation(t *testing.T) {
	o, _ := NewObject("test", verts)
	o.AbsTranslate(5, 0, 0)
	comp := []*Vertex{
		NewVertex(6, 0, 0, "x"),
		NewVertex(5, 1, 0, "y"),
		NewVertex(5, 0, 1, "z"),
	}
	for _, v1 := range o.Verts {
		for _, v2 := range comp {
			if v2.Name == v1.Name {
				if !v1.ApproxEqual(v2) {
					t.Errorf("vertex not translated correctly, should be %s, but is %s", v2, v1)
				}
			}
		}
	}
}
