package object

import (
	"fmt"
	"github.com/go-gl/mathgl/mgl32"
)

//Vertex is the structure that holds the vector
//for a point in space.  This is most likely
//used in an object struct
//
//NewVertex should be used as many unexported
//fields need to be populated
type Vertex struct {
	//This vector is the original inputed vector
	//It will never changed unless directly
	vec mgl32.Vec3

	//This vector is the vector offset by the
	//new rotation origin.  This is separate from vec
	//so that new rotation points on the object
	//can be set
	rvec mgl32.Vec3

	//cvec is the current vector that has been\
	//translated and rotated
	cvec mgl32.Vec3

	//oldvec is the last cvec. This gives us the
	//ability to reset the vector
	oldvec mgl32.Vec3

	//Name is here for debugging and pretty error
	//printing.
	Name string

	//id is the internal ID given to make this Vertex
	//unique
	id int
}

var vertexid int = 0

func NewVertex(x, y, z float32, name string) *Vertex {
	i := vertexid
	vertexid += 1
	return &Vertex{
		vec:  mgl32.Vec3{x, y, z},
		rvec: mgl32.Vec3{x, y, z},
		cvec: mgl32.Vec3{x, y, z},
		Name: name,
		id:   i,
	}
}

func (v *Vertex) String() string {
	return fmt.Sprintf("Vertex: %s CurrentLocation x:%e y:%e z:%e", v.Name, v.cvec.X(), v.cvec.Y(), v.cvec.Z())
}

func (v *Vertex) setNewRotationOrigin(newRotOrig *Vertex) {
	v.rvec = v.rvec.Sub(newRotOrig.vec)
}

func (v *Vertex) transformToPosition(p *Position) {
	v.oldvec = v.cvec
	v.cvec = p.Quat.Rotate(v.rvec)
	v.cvec = v.cvec.Add(p.CurOffset)
}

func (v *Vertex) undo() {
	v.cvec = v.oldvec
}

func (v *Vertex) ApproxEqual(v2 *Vertex) bool {
	return v.cvec.ApproxEqualThreshold(v2.cvec, 1e-3)
}
