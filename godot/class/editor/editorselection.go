package editor

import (
	"log"

	"github.com/shadowapex/godot-go/gdnative"

	"github.com/shadowapex/godot-go/godot/class/object"
)

/*------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "class.go.tmpl" so they can be included in the generated
//   code.
//----------------------------------------------------------------------------*/

func NewEditorSelectionFromPointer(ptr gdnative.Pointer) *EditorSelection {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := EditorSelection{}
	obj.SetOwner(owner)

	return &obj

}

/*
This object manages the SceneTree selection in the editor.
*/
type EditorSelection struct {
	object.Object
}

func (o *EditorSelection) BaseClass() string {
	return "EditorSelection"
}

/*
        Undocumented
	Args: [], Returns: void
*/

func (o *EditorSelection) X_EmitChange() {
	log.Println("Calling EditorSelection.X_EmitChange()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorSelection", "_emit_change")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false arg0 Object}], Returns: void
*/

func (o *EditorSelection) X_NodeRemoved(arg0 object.Object) {
	log.Println("Calling EditorSelection.X_NodeRemoved()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(arg0.GetOwner())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorSelection", "_node_removed")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*
        Add a node to the selection.
	Args: [{ false node Object}], Returns: void
*/

func (o *EditorSelection) AddNode(node object.Object) {
	log.Println("Calling EditorSelection.AddNode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(node.GetOwner())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorSelection", "add_node")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*
        Clear the selection.
	Args: [], Returns: void
*/

func (o *EditorSelection) Clear() {
	log.Println("Calling EditorSelection.Clear()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorSelection", "clear")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*
        Get the list of selected nodes.
	Args: [], Returns: Array
*/

func (o *EditorSelection) GetSelectedNodes() gdnative.Array {
	log.Println("Calling EditorSelection.GetSelectedNodes()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorSelection", "get_selected_nodes")

	// Call the parent method.
	// Array
	retPtr := gdnative.NewEmptyArray()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewArrayFromPointer(retPtr)

	log.Println("  Got return value: ", ret)
	return ret
}

/*
        Get the list of selected nodes, optimized for transform operations (ie, moving them, rotating, etc). This list avoids situations where a node is selected and also chid/grandchild.
	Args: [], Returns: Array
*/

func (o *EditorSelection) GetTransformableSelectedNodes() gdnative.Array {
	log.Println("Calling EditorSelection.GetTransformableSelectedNodes()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorSelection", "get_transformable_selected_nodes")

	// Call the parent method.
	// Array
	retPtr := gdnative.NewEmptyArray()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewArrayFromPointer(retPtr)

	log.Println("  Got return value: ", ret)
	return ret
}

/*
        Remove a node from the selection.
	Args: [{ false node Object}], Returns: void
*/

func (o *EditorSelection) RemoveNode(node object.Object) {
	log.Println("Calling EditorSelection.RemoveNode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(node.GetOwner())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorSelection", "remove_node")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}
