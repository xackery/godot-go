package packedscene

import (
	"log"

	"github.com/shadowapex/godot-go/gdnative"

	"github.com/shadowapex/godot-go/godot/class/resource"
	"github.com/shadowapex/godot-go/godot/class/scenestate"

	"github.com/shadowapex/godot-go/godot/class/node"
)

/*------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "class.go.tmpl" so they can be included in the generated
//   code.
//----------------------------------------------------------------------------*/

func NewPackedSceneFromPointer(ptr gdnative.Pointer) *PackedScene {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := PackedScene{}
	obj.SetOwner(owner)

	return &obj

}

/*
A simplified interface to a scene file. Provides access to operations and checks that can be performed on the scene resource itself. TODO: explain ownership, and that node does not need to own itself
*/
type PackedScene struct {
	resource.Resource
}

func (o *PackedScene) BaseClass() string {
	return "PackedScene"
}

/*
        Undocumented
	Args: [], Returns: Dictionary
*/

func (o *PackedScene) X_GetBundledScene() gdnative.Dictionary {
	log.Println("Calling PackedScene.X_GetBundledScene()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PackedScene", "_get_bundled_scene")

	// Call the parent method.
	// Dictionary
	retPtr := gdnative.NewEmptyDictionary()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewDictionaryFromPointer(retPtr)

	log.Println("  Got return value: ", ret)
	return ret
}

/*
        Undocumented
	Args: [{ false arg0 Dictionary}], Returns: void
*/

func (o *PackedScene) X_SetBundledScene(arg0 gdnative.Dictionary) {
	log.Println("Calling PackedScene.X_SetBundledScene()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromDictionary(arg0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PackedScene", "_set_bundled_scene")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*
        Returns [code]true[/code] if the scene file has nodes.
	Args: [], Returns: bool
*/

func (o *PackedScene) CanInstance() gdnative.Bool {
	log.Println("Calling PackedScene.CanInstance()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PackedScene", "can_instance")

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
        Returns the [code]SceneState[/code] representing the scene file contents.
	Args: [], Returns: SceneState
*/

func (o *PackedScene) GetState() scenestate.SceneState {
	log.Println("Calling PackedScene.GetState()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PackedScene", "get_state")

	// Call the parent method.
	// SceneState
	retPtr := scenestate.NewEmptySceneState()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := scenestate.NewSceneStateFromPointer(retPtr)

	log.Println("  Got return value: ", ret)
	return ret
}

/*
        Instantiates the scene's node hierarchy. Triggers child scene instantiation(s). Triggers the [enum Object.NOTIFICATION_INSTANCED] notification on the root node.
	Args: [{0 true edit_state int}], Returns: Node
*/

func (o *PackedScene) Instance(editState gdnative.Int) node.Node {
	log.Println("Calling PackedScene.Instance()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(editState)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PackedScene", "instance")

	// Call the parent method.
	// Node
	retPtr := node.NewEmptyNode()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := node.NewNodeFromPointer(retPtr)

	log.Println("  Got return value: ", ret)
	return ret
}

/*
        Pack will ignore any sub-nodes not owned by given node. See [method Node.set_owner].
	Args: [{ false path Object}], Returns: enum.Error
*/
