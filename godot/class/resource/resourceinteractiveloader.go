package resource

import (
	"log"

	"github.com/shadowapex/godot-go/gdnative"
	"github.com/shadowapex/godot-go/godot/class/reference"
)

/*------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "class.go.tmpl" so they can be included in the generated
//   code.
//----------------------------------------------------------------------------*/

func NewResourceInteractiveLoaderFromPointer(ptr gdnative.Pointer) *ResourceInteractiveLoader {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := ResourceInteractiveLoader{}
	obj.SetOwner(owner)

	return &obj

}

/*
Interactive Resource Loader. This object is returned by ResourceLoader when performing an interactive load. It allows to load with high granularity, so this is mainly useful for displaying load bars/percentages.
*/
type ResourceInteractiveLoader struct {
	reference.Reference
}

func (o *ResourceInteractiveLoader) BaseClass() string {
	return "ResourceInteractiveLoader"
}

/*
        Return the loaded resource (only if loaded). Otherwise, returns null.
	Args: [], Returns: Resource
*/

func (o *ResourceInteractiveLoader) GetResource() Resource {
	log.Println("Calling ResourceInteractiveLoader.GetResource()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ResourceInteractiveLoader", "get_resource")

	// Call the parent method.
	// Resource
	retPtr := NewEmptyResource()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := NewResourceFromPointer(retPtr)

	log.Println("  Got return value: ", ret)
	return ret
}

/*
        Return the load stage. The total amount of stages can be queried with [method get_stage_count]
	Args: [], Returns: int
*/

func (o *ResourceInteractiveLoader) GetStage() gdnative.Int {
	log.Println("Calling ResourceInteractiveLoader.GetStage()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ResourceInteractiveLoader", "get_stage")

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
        Return the total amount of stages (calls to [method poll]) needed to completely load this resource.
	Args: [], Returns: int
*/

func (o *ResourceInteractiveLoader) GetStageCount() gdnative.Int {
	log.Println("Calling ResourceInteractiveLoader.GetStageCount()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ResourceInteractiveLoader", "get_stage_count")

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
        Poll the load. If OK is returned, this means poll will have to be called again. If ERR_FILE_EOF is returned, them the load has finished and the resource can be obtained by calling [method get_resource].
	Args: [], Returns: enum.Error
*/

/*

	Args: [], Returns: enum.Error
*/
