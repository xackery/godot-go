package cylindermesh

import (
	"log"

	"github.com/shadowapex/godot-go/gdnative"
	"github.com/shadowapex/godot-go/godot/class/primitivemesh"
)

/*------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "class.go.tmpl" so they can be included in the generated
//   code.
//----------------------------------------------------------------------------*/

func NewCylinderMeshFromPointer(ptr gdnative.Pointer) *CylinderMesh {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := CylinderMesh{}
	obj.SetOwner(owner)

	return &obj

}

/*
Class representing a cylindrical [PrimitiveMesh].
*/
type CylinderMesh struct {
	primitivemesh.PrimitiveMesh
}

func (o *CylinderMesh) BaseClass() string {
	return "CylinderMesh"
}

/*
        Undocumented
	Args: [], Returns: float
*/

func (o *CylinderMesh) GetBottomRadius() gdnative.Float {
	log.Println("Calling CylinderMesh.GetBottomRadius()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CylinderMesh", "get_bottom_radius")

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

func (o *CylinderMesh) GetHeight() gdnative.Float {
	log.Println("Calling CylinderMesh.GetHeight()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CylinderMesh", "get_height")

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
	Args: [], Returns: int
*/

func (o *CylinderMesh) GetRadialSegments() gdnative.Int {
	log.Println("Calling CylinderMesh.GetRadialSegments()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CylinderMesh", "get_radial_segments")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)

	log.Println("  Got return value: ", ret)
	return ret
}

/*
        Undocumented
	Args: [], Returns: int
*/

func (o *CylinderMesh) GetRings() gdnative.Int {
	log.Println("Calling CylinderMesh.GetRings()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CylinderMesh", "get_rings")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)

	log.Println("  Got return value: ", ret)
	return ret
}

/*
        Undocumented
	Args: [], Returns: float
*/

func (o *CylinderMesh) GetTopRadius() gdnative.Float {
	log.Println("Calling CylinderMesh.GetTopRadius()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CylinderMesh", "get_top_radius")

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
	Args: [{ false radius float}], Returns: void
*/

func (o *CylinderMesh) SetBottomRadius(radius gdnative.Float) {
	log.Println("Calling CylinderMesh.SetBottomRadius()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromFloat(radius)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CylinderMesh", "set_bottom_radius")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false height float}], Returns: void
*/

func (o *CylinderMesh) SetHeight(height gdnative.Float) {
	log.Println("Calling CylinderMesh.SetHeight()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromFloat(height)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CylinderMesh", "set_height")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false segments int}], Returns: void
*/

func (o *CylinderMesh) SetRadialSegments(segments gdnative.Int) {
	log.Println("Calling CylinderMesh.SetRadialSegments()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(segments)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CylinderMesh", "set_radial_segments")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false rings int}], Returns: void
*/

func (o *CylinderMesh) SetRings(rings gdnative.Int) {
	log.Println("Calling CylinderMesh.SetRings()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(rings)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CylinderMesh", "set_rings")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false radius float}], Returns: void
*/

func (o *CylinderMesh) SetTopRadius(radius gdnative.Float) {
	log.Println("Calling CylinderMesh.SetTopRadius()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromFloat(radius)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CylinderMesh", "set_top_radius")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}
