package object

import (
	"fmt"
	"github.com/go-gl/mathgl/mgl32"
	"math"
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

	//rtvec will set the vertex back to the original
	//origin space for the translation after rotating
	//about the rotation point
	rtvec mgl32.Vec3

	//tvec is the home position of the vertex
	tvec mgl32.Vec3

	//cvec is the current vector that has been\
	//translated and rotated
	cvec mgl32.Vec3

	//oldvec is the last cvec. This gives us the
	//ability to reset the vector
	oldvec mgl32.Vec3

	//object is the object associated with the vertex
	obj *Object

	//Name is here for debugging and pretty error
	//printing.
	Name string

	//id is the internal ID given to make this Vertex
	//unique
	id int
}

var vertexid int

//NewVertex creates and initializes a Vertex.  A unique
//id is generated for each Vertex created by this method.
func NewVertex(x, y, z float32, name string) *Vertex {
	i := vertexid
	vertexid += 1
	return &Vertex{
		vec:   mgl32.Vec3{x, y, z},
		rvec:  mgl32.Vec3{x, y, z},
		cvec:  mgl32.Vec3{x, y, z},
		rtvec: mgl32.Vec3{0, 0, 0},
		tvec:  mgl32.Vec3{0, 0, 0},
		Name:  name,
		id:    i,
	}
}

//SetObject sets Object that this vertex is associated with
func (v *Vertex) SetObject(o *Object) {
	v.obj = o
}

func (v *Vertex) String() string {
	return fmt.Sprintf("Vertex: %s CurrentLocation x:%f y:%f z:%f", v.Name, v.cvec.X(), v.cvec.Y(), v.cvec.Z())
}

//SetRotationOriginVert sets the point the the vertex is rotated about based on the current
//position of newRot
func (v *Vertex) SetRotationOriginVert(newRot *Vertex) {
	v.rtvec = newRot.copy()
	v.rvec = v.rvec.Sub(newRot.vec)
}

//SetOffsetOrigin moves the home position of the vertex
//what this means is that when the vertex is translated
//absolutely it is from this point in space
func (v *Vertex) SetOffsetOrigin(x, y, z float32) {
	v.tvec = mgl32.Vec3{x, y, z}
	v.cvec = v.vec.Add(v.tvec)
}

//TransformToPosition moves the vertex to the position p.  All rotations
//are around the origin unless the Rotation Origin
//is set, in which case it is rotated about that.  It
//is translated absolutely to the origin 0, 0, 0 unless
//the point is offset by SetOffset.
func (v *Vertex) TransformToPosition(p *Position) {
	v.oldvec = v.cvec
	v.cvec = p.quat.Rotate(v.rvec)
	v.cvec = v.cvec.Add(v.rtvec)
	v.cvec = v.cvec.Add(v.tvec)
	v.cvec = v.cvec.Add(p.trans)
}

//Undo resets the vertex to the position previous the
//last call to TransformToPosition.  The vertex
//only stores 1 previous position, so multiple
//calls to this do nothing.
func (v *Vertex) Undo() {
	v.cvec = v.oldvec
}

//ApproxEqual calculates whether the current positions of vertexes v
//and v2 are approximately equal
func (v *Vertex) ApproxEqual(v2 *Vertex) bool {
	return v.cvec.ApproxEqualThreshold(v2.cvec, 1e-3)
}

//Less calculates whether the vertex v is less than
//vertex v2.  It does this by first comparing X
//then Y then Z coordinates.
func (v *Vertex) Less(v2 *Vertex) bool {
	if v.cvec.X() < v2.cvec.X() {
		return true
	} else if v.cvec.X() > v2.cvec.X() {
		return false
	} else {
		if v.cvec.Y() < v2.cvec.Y() {
			return true
		} else if v.cvec.Y() > v2.cvec.Y() {
			return false
		} else {
			if v.cvec.Z() < v.cvec.Z() {
				return true
			}
			return false
		}
	}
}

//DistanceToVertex computes the distance between the two vertexes
func (v *Vertex) DistanceToVertex(v2 *Vertex) float32 {
	return v.cvec.Sub(v2.cvec).Len()
}

//X returns the X coordinate of the current position
func (v *Vertex) X() float32 {
	return v.cvec.X()
}

//Y returns the Y coordinate of the current position
func (v *Vertex) Y() float32 {
	return v.cvec.Y()
}

//Z returns the Z coordinate of the current position
func (v *Vertex) Z() float32 {
	return v.cvec.Z()
}

//Id returns the unique id given to every new Vertex
func (v *Vertex) Id() int {
	return v.id
}

func (v *Vertex) copy() mgl32.Vec3 {
	return mgl32.Vec3{v.X(), v.Y(), v.Z()}
}

//Spherical returns the spherical coordinates of the current Position
func (v *Vertex) Spherical() (r, theta, phi float32) {
	r, theta, phi = mgl32.CartesianToSpherical(v.cvec)
	theta = mgl32.RadToDeg(theta)
	phi = mgl32.RadToDeg(phi)
	return
}

//Dot gets the dot product of the Vertexes current position
func (v *Vertex) Dot(v2 *Vertex) float32 {
	return v.cvec.Dot(v2.cvec)
}

//DotAngle computes the angle between the current position and the previous position
func (v *Vertex) DotAngle() float32 {
	return mgl32.RadToDeg(float32(math.Acos(float64(v.cvec.Dot(v.oldvec) / (v.cvec.Len() * v.oldvec.Len())))))
}

//DistChange computes the change in distance to the current position and the previous position
//of the vertex.  A positive change is where the current position length is greater than the
//previous position
func (v *Vertex) DistChange() float32 {
	return v.cvec.Len() - v.oldvec.Len()
}
