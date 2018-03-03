package capsuleshape2d

import (
	"log"

	"github.com/shadowapex/godot-go/gdnative"
	"github.com/shadowapex/godot-go/godot/class/shape2d"
)

/*------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "class.go.tmpl" so they can be included in the generated
//   code.
//----------------------------------------------------------------------------*/

func NewCapsuleShape2DFromPointer(ptr gdnative.Pointer) *CapsuleShape2D {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := CapsuleShape2D{}
	obj.SetOwner(owner)

	return &obj

}

/*
Capsule shape for 2D collisions.
*/
type CapsuleShape2D struct {
	shape2d.Shape2D
}

func (o *CapsuleShape2D) BaseClass() string {
	return "CapsuleShape2D"
}

/*
        Undocumented
	Args: [], Returns: float
*/

func (o *CapsuleShape2D) GetHeight() gdnative.Float {
	log.Println("Calling CapsuleShape2D.GetHeight()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CapsuleShape2D", "get_height")

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
	Args: [], Returns: float
*/

func (o *CapsuleShape2D) GetRadius() gdnative.Float {
	log.Println("Calling CapsuleShape2D.GetRadius()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CapsuleShape2D", "get_radius")

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
	Args: [{ false height float}], Returns: void
*/

func (o *CapsuleShape2D) SetHeight(height gdnative.Float) {
	log.Println("Calling CapsuleShape2D.SetHeight()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromFloat(height)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CapsuleShape2D", "set_height")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false radius float}], Returns: void
*/

func (o *CapsuleShape2D) SetRadius(radius gdnative.Float) {
	log.Println("Calling CapsuleShape2D.SetRadius()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromFloat(radius)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CapsuleShape2D", "set_radius")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}
