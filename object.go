package object

import (
	"fmt"
)

type Object struct {
	//Verts are the vertices of the object
	//in its own object coordinates with the
	//origin vertex 0,0,0
	Verts  []*Vertex
	Origin int

	//Name for the object, helpful if you want to debug
	Name string

	//Current Position of the object
	cp *Position

	//old position is the last position
	//so we can snap back
	oldpos *Position
}

//Creates new object and automatically adds a 0,0,0 Vertex
//vertices mus have at least one vectex in it
func NewObject(name string, vertices []*Vertex) (*Object, error) {
	if len(vertices) == 0 {
		return nil, fmt.Errorf("object needs at least one vertex")
	}
	v := make([]*Vertex, 0, 1+len(vertices))
	v = append(v, NewVertex(0, 0, 0, "origin"))
	v = append(v, vertices...)

	return &Object{Verts: v,
		Origin: 0,
		Name:   name,
		cp:     NewPosition(),
	}, nil
}

func (o *Object) AddVertex(vert *Vertex, makeOrigin bool) error {
	o.Verts = append(o.Verts, vert)
	if makeOrigin {
		return o.SetNewOrigin(len(o.Verts) - 1)
	}
	return nil
}

//Sets the vertex index as the new origin
//of the object.  The origin of the object
//is where the object rotates around
func (o *Object) SetNewOrigin(index int) error {
	if index >= len(o.Verts) {
		return fmt.Errorf("index not accessable")
	}
	o.Origin = index
	for i := range o.Verts {
		o.Verts[i].setNewRotationOrigin(o.Verts[index])
	}
	return nil
}

func (o *Object) AbsRotate(x, y, z float32) {
	t := NewTransform(x, y, z, RotAbs, 0, 0, 0, TranNone)
	o.oldpos = o.cp.Copy()
	t.TransformPosition(o.cp)
	o.transformVerts()
}

func (o *Object) AbsTranslate(x, y, z float32) {
	t := NewTransform(0, 0, 0, RotNone, x, y, z, TranAbs)
	o.oldpos = o.cp.Copy()
	t.TransformPosition(o.cp)
	o.transformVerts()
}

func (o *Object) transformVerts() {
	for i := range o.Verts {
		o.Verts[i].transformToPosition(o.cp)
	}
}

func (o *Object) CopyCurrentPosition() *Position {
	return o.cp.Copy()
}

func (o *Object) SetPosition(pos *Position) {
	o.oldpos = o.cp.Copy()
	o.cp.Overwrite(pos)
	o.transformVerts()
}

func (o *Object) Undo() {
	o.cp.Overwrite(o.oldpos)
	for i := range o.Verts {
		o.Verts[i].undo()
	}
}
