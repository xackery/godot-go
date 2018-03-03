package collisionpolygon

import (
	"log"

	"github.com/shadowapex/godot-go/gdnative"
	"github.com/shadowapex/godot-go/godot/class/spatial"
)

/*------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "class.go.tmpl" so they can be included in the generated
//   code.
//----------------------------------------------------------------------------*/

func NewCollisionPolygonFromPointer(ptr gdnative.Pointer) *CollisionPolygon {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := CollisionPolygon{}
	obj.SetOwner(owner)

	return &obj

}

/*
Allows editing a collision polygon's vertices on a selected plane. Can also set a depth perpendicular to that plane. This class is only available in the editor. It will not appear in the scene tree at runtime. Creates a [Shape] for gameplay. Properties modified during gameplay will have no effect.
*/
type CollisionPolygon struct {
	spatial.Spatial
}

func (o *CollisionPolygon) BaseClass() string {
	return "CollisionPolygon"
}

/*
        Undocumented
	Args: [], Returns: float
*/

func (o *CollisionPolygon) GetDepth() gdnative.Float {
	log.Println("Calling CollisionPolygon.GetDepth()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CollisionPolygon", "get_depth")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyFloat()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewFloatFromPointer(retPtr)

	log.Println("  Got return value: ", ret)
	return ret
}

/*
        Undocumented
	Args: [], Returns: PoolVector2Array
*/

func (o *CollisionPolygon) GetPolygon() gdnative.PoolVector2Array {
	log.Println("Calling CollisionPolygon.GetPolygon()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CollisionPolygon", "get_polygon")

	// Call the parent method.
	// PoolVector2Array
	retPtr := gdnative.NewEmptyPoolVector2Array()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewPoolVector2ArrayFromPointer(retPtr)

	log.Println("  Got return value: ", ret)
	return ret
}

/*
        Undocumented
	Args: [], Returns: bool
*/

func (o *CollisionPolygon) IsDisabled() gdnative.Bool {
	log.Println("Calling CollisionPolygon.IsDisabled()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CollisionPolygon", "is_disabled")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)

	log.Println("  Got return value: ", ret)
	return ret
}

/*
        Undocumented
	Args: [{ false depth float}], Returns: void
*/

func (o *CollisionPolygon) SetDepth(depth gdnative.Float) {
	log.Println("Calling CollisionPolygon.SetDepth()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromFloat(depth)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CollisionPolygon", "set_depth")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false disabled bool}], Returns: void
*/

func (o *CollisionPolygon) SetDisabled(disabled gdnative.Bool) {
	log.Println("Calling CollisionPolygon.SetDisabled()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(disabled)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CollisionPolygon", "set_disabled")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false polygon PoolVector2Array}], Returns: void
*/

func (o *CollisionPolygon) SetPolygon(polygon gdnative.PoolVector2Array) {
	log.Println("Calling CollisionPolygon.SetPolygon()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromPoolVector2Array(polygon)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CollisionPolygon", "set_polygon")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}
