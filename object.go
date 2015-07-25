package object

import "fmt"

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

	o := &Object{Verts: v,
		Origin: 0,
		Name:   name,
		cp:     NewPosition(),
	}
	for i := range o.Verts {
		o.Verts[i].obj = o
	}

	return o, nil
}

func (o *Object) AddVertex(vert *Vertex, makeOrigin bool) error {
	o.Verts = append(o.Verts, vert)
	vert.obj = o
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
		o.Verts[i].SetRotationOriginVert(o.Verts[index])
	}
	return nil
}

func (o *Object) AbsRotate(x, y, z float32) {
	t := NewTransform(x, y, z, RotAbs, 0, 0, 0, TranNone)
	o.transformVerts(t)
}

//Rotate the object in the absolute reference frame
//A combination of two 90 degree rotates on the X axis
//should be equal to 1 180 degree rotate on the X axis
func (o *Object) RelRotateAbsRef(x, y, z float32) {
	t := NewTransform(x, y, z, RotAddAbs, 0, 0, 0, TranNone)
	o.transformVerts(t)
}

func (o *Object) AbsTranslate(x, y, z float32) {
	t := NewTransform(0, 0, 0, RotNone, x, y, z, TranAbs)
	o.transformVerts(t)
}

func (o *Object) RelTranslate(x, y, z float32) {
	t := NewTransform(0, 0, 0, RotNone, x, y, z, TranRel)
	o.transformVerts(t)
}

//This sets the translation offset of the object in world space
//The origin point of the object is set to these coordinates
//All absolute translations will happen from this point.
func (o *Object) SetObjectOffset(x, y, z float32) {
	for i := range o.Verts {
		o.Verts[i].SetOffsetOrigin(x, y, z)
	}
}

func (o *Object) transformVerts(t *Transform) {
	o.oldpos = o.cp.Copy()
	t.TransformPosition(o.cp)
	o.transVertsPos()
}

func (o *Object) transVertsPos() {
	for i := range o.Verts {
		o.Verts[i].TransformToPosition(o.cp)
	}
}

func (o *Object) CopyCurrentPosition() *Position {
	return o.cp.Copy()
}

func (o *Object) SetPosition(pos *Position) {
	o.oldpos = o.cp.Copy()
	o.cp.Overwrite(pos)
	o.transVertsPos()
}

func (o *Object) Undo() {
	o.cp.Overwrite(o.oldpos)
	for i := range o.Verts {
		o.Verts[i].Undo()
	}
}
