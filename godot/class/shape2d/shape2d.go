package shape2d

import (
	"log"

	"github.com/shadowapex/godot-go/gdnative"
	"github.com/shadowapex/godot-go/godot/class/resource"
)

/*------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "class.go.tmpl" so they can be included in the generated
//   code.
//----------------------------------------------------------------------------*/

func NewShape2DFromPointer(ptr gdnative.Pointer) *Shape2D {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := Shape2D{}
	obj.SetOwner(owner)

	return &obj

}

/*
Base class for all 2D Shapes. All 2D shape types inherit from this.
*/
type Shape2D struct {
	resource.Resource
}

func (o *Shape2D) BaseClass() string {
	return "Shape2D"
}

/*
        Return whether this shape is colliding with another. This method needs the transformation matrix for this shape ([code]local_xform[/code]), the shape to check collisions with ([code]with_shape[/code]), and the transformation matrix of that shape ([code]shape_xform[/code]).
	Args: [{ false local_xform Transform2D} { false with_shape Shape2D} { false shape_xform Transform2D}], Returns: bool
*/

func (o *Shape2D) Collide(localXform gdnative.Transform2D, withShape Shape2D, shapeXform gdnative.Transform2D) gdnative.Bool {
	log.Println("Calling Shape2D.Collide()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 3, 3)
	ptrArguments[0] = gdnative.NewPointerFromTransform2D(localXform)
	ptrArguments[1] = gdnative.NewPointerFromObject(withShape.GetOwner())
	ptrArguments[2] = gdnative.NewPointerFromTransform2D(shapeXform)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Shape2D", "collide")

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
        Return a list of the points where this shape touches another. If there are no collisions, the list is empty. This method needs the transformation matrix for this shape ([code]local_xform[/code]), the shape to check collisions with ([code]with_shape[/code]), and the transformation matrix of that shape ([code]shape_xform[/code]).
	Args: [{ false local_xform Transform2D} { false with_shape Shape2D} { false shape_xform Transform2D}], Returns: Variant
*/

func (o *Shape2D) CollideAndGetContacts(localXform gdnative.Transform2D, withShape Shape2D, shapeXform gdnative.Transform2D) gdnative.Variant {
	log.Println("Calling Shape2D.CollideAndGetContacts()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 3, 3)
	ptrArguments[0] = gdnative.NewPointerFromTransform2D(localXform)
	ptrArguments[1] = gdnative.NewPointerFromObject(withShape.GetOwner())
	ptrArguments[2] = gdnative.NewPointerFromTransform2D(shapeXform)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Shape2D", "collide_and_get_contacts")

	// Call the parent method.
	// Variant
	retPtr := gdnative.NewEmptyVariant()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVariantFromPointer(retPtr)

	log.Println("  Got return value: ", ret)
	return ret
}

/*
        Return whether this shape would collide with another, if a given movement was applied. This method needs the transformation matrix for this shape ([code]local_xform[/code]), the movement to test on this shape ([code]local_motion[/code]), the shape to check collisions with ([code]with_shape[/code]), the transformation matrix of that shape ([code]shape_xform[/code]), and the movement to test onto the other object ([code]shape_motion[/code]).
	Args: [{ false local_xform Transform2D} { false local_motion Vector2} { false with_shape Shape2D} { false shape_xform Transform2D} { false shape_motion Vector2}], Returns: bool
*/

func (o *Shape2D) CollideWithMotion(localXform gdnative.Transform2D, localMotion gdnative.Vector2, withShape Shape2D, shapeXform gdnative.Transform2D, shapeMotion gdnative.Vector2) gdnative.Bool {
	log.Println("Calling Shape2D.CollideWithMotion()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 5, 5)
	ptrArguments[0] = gdnative.NewPointerFromTransform2D(localXform)
	ptrArguments[1] = gdnative.NewPointerFromVector2(localMotion)
	ptrArguments[2] = gdnative.NewPointerFromObject(withShape.GetOwner())
	ptrArguments[3] = gdnative.NewPointerFromTransform2D(shapeXform)
	ptrArguments[4] = gdnative.NewPointerFromVector2(shapeMotion)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Shape2D", "collide_with_motion")

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
        Return a list of the points where this shape would touch another, if a given movement was applied. If there are no collisions, the list is empty. This method needs the transformation matrix for this shape ([code]local_xform[/code]), the movement to test on this shape ([code]local_motion[/code]), the shape to check collisions with ([code]with_shape[/code]), the transformation matrix of that shape ([code]shape_xform[/code]), and the movement to test onto the other object ([code]shape_motion[/code]).
	Args: [{ false local_xform Transform2D} { false local_motion Vector2} { false with_shape Shape2D} { false shape_xform Transform2D} { false shape_motion Vector2}], Returns: Variant
*/

func (o *Shape2D) CollideWithMotionAndGetContacts(localXform gdnative.Transform2D, localMotion gdnative.Vector2, withShape Shape2D, shapeXform gdnative.Transform2D, shapeMotion gdnative.Vector2) gdnative.Variant {
	log.Println("Calling Shape2D.CollideWithMotionAndGetContacts()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 5, 5)
	ptrArguments[0] = gdnative.NewPointerFromTransform2D(localXform)
	ptrArguments[1] = gdnative.NewPointerFromVector2(localMotion)
	ptrArguments[2] = gdnative.NewPointerFromObject(withShape.GetOwner())
	ptrArguments[3] = gdnative.NewPointerFromTransform2D(shapeXform)
	ptrArguments[4] = gdnative.NewPointerFromVector2(shapeMotion)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Shape2D", "collide_with_motion_and_get_contacts")

	// Call the parent method.
	// Variant
	retPtr := gdnative.NewEmptyVariant()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVariantFromPointer(retPtr)

	log.Println("  Got return value: ", ret)
	return ret
}

/*
        Undocumented
	Args: [], Returns: float
*/

func (o *Shape2D) GetCustomSolverBias() gdnative.Float {
	log.Println("Calling Shape2D.GetCustomSolverBias()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Shape2D", "get_custom_solver_bias")

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
	Args: [{ false bias float}], Returns: void
*/

func (o *Shape2D) SetCustomSolverBias(bias gdnative.Float) {
	log.Println("Calling Shape2D.SetCustomSolverBias()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromFloat(bias)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Shape2D", "set_custom_solver_bias")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}
