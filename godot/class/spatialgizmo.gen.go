package class

import (
	"github.com/shadowapex/godot-go/gdnative"
)

/*------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "class.go.tmpl" so they can be included in the generated
//   code.
//----------------------------------------------------------------------------*/

//func NewSpatialGizmoFromPointer(ptr gdnative.Pointer) SpatialGizmo {
func NewSpatialGizmoFromPointer(ptr gdnative.Pointer) SpatialGizmo {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := SpatialGizmo{}
	obj.SetBaseObject(owner)

	return obj
}

/*

 */
type SpatialGizmo struct {
	Reference
	owner gdnative.Object
}

func (o *SpatialGizmo) BaseClass() string {
	return "SpatialGizmo"
}

// SetBaseObject will internally set the Godot object inside the struct.
// This is used to call parent methods.
func (o *SpatialGizmo) SetBaseObject(object gdnative.Object) {
	o.owner = object
}

func (o *SpatialGizmo) GetBaseObject() gdnative.Object {
	return o.owner
}