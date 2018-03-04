package class

import (
	"log"

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

//func NewMultiMeshInstanceFromPointer(ptr gdnative.Pointer) MultiMeshInstance {
func NewMultiMeshInstanceFromPointer(ptr gdnative.Pointer) MultiMeshInstance {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := MultiMeshInstance{}
	obj.SetBaseObject(owner)

	return obj
}

/*
[code]MultiMeshInstance[/code] is a specialized node to instance [GeometryInstance]s based on a [MultiMesh] resource. This is useful to optimize the rendering of a high amount of instances of a given mesh (for example tree in a forest or grass strands).
*/
type MultiMeshInstance struct {
	GeometryInstance
	owner gdnative.Object
}

func (o *MultiMeshInstance) BaseClass() string {
	return "MultiMeshInstance"
}

// SetBaseObject will internally set the Godot object inside the struct.
// This is used to call parent methods.
func (o *MultiMeshInstance) SetBaseObject(object gdnative.Object) {
	o.owner = object
}

func (o *MultiMeshInstance) GetBaseObject() gdnative.Object {
	return o.owner
}

/*
        Undocumented
	Args: [], Returns: MultiMesh
*/
func (o *MultiMeshInstance) GetMultimesh() MultiMesh {
	log.Println("Calling MultiMeshInstance.GetMultimesh()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("MultiMeshInstance", "get_multimesh")

	// Call the parent method.
	// MultiMesh
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := NewMultiMeshFromPointer(retPtr)
	log.Println("  Got return value: ", ret)
	return ret
}

/*
        Undocumented
	Args: [{ false multimesh MultiMesh}], Returns: void
*/
func (o *MultiMeshInstance) SetMultimesh(multimesh MultiMesh) {
	log.Println("Calling MultiMeshInstance.SetMultimesh()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(multimesh.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("MultiMeshInstance", "set_multimesh")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}