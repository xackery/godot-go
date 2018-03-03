package scenetree

import (
	"log"

	"github.com/shadowapex/godot-go/gdnative"

	"github.com/shadowapex/godot-go/godot/class/mainloop"
	"github.com/shadowapex/godot-go/godot/class/networkedmultiplayerpeer"

	"github.com/shadowapex/godot-go/godot/class/viewport"

	"github.com/shadowapex/godot-go/godot/class/object"

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

func NewSceneTreeFromPointer(ptr gdnative.Pointer) *SceneTree {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := SceneTree{}
	obj.SetOwner(owner)

	return &obj

}

/*

 */
type SceneTree struct {
	mainloop.MainLoop
}

func (o *SceneTree) BaseClass() string {
	return "SceneTree"
}

/*
        Undocumented
	Args: [{ false arg0 Object}], Returns: void
*/

func (o *SceneTree) X_ChangeScene(arg0 object.Object) {
	log.Println("Calling SceneTree.X_ChangeScene()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(arg0.GetOwner())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SceneTree", "_change_scene")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: void
*/

func (o *SceneTree) X_ConnectedToServer() {
	log.Println("Calling SceneTree.X_ConnectedToServer()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SceneTree", "_connected_to_server")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: void
*/

func (o *SceneTree) X_ConnectionFailed() {
	log.Println("Calling SceneTree.X_ConnectionFailed()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SceneTree", "_connection_failed")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false arg0 int}], Returns: void
*/

func (o *SceneTree) X_NetworkPeerConnected(arg0 gdnative.Int) {
	log.Println("Calling SceneTree.X_NetworkPeerConnected()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(arg0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SceneTree", "_network_peer_connected")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false arg0 int}], Returns: void
*/

func (o *SceneTree) X_NetworkPeerDisconnected(arg0 gdnative.Int) {
	log.Println("Calling SceneTree.X_NetworkPeerDisconnected()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(arg0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SceneTree", "_network_peer_disconnected")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: void
*/

func (o *SceneTree) X_ServerDisconnected() {
	log.Println("Calling SceneTree.X_ServerDisconnected()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SceneTree", "_server_disconnected")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*

	Args: [{ false group String} { false method String}], Returns: Variant
*/

func (o *SceneTree) CallGroup(group gdnative.String, method gdnative.String) gdnative.Variant {
	log.Println("Calling SceneTree.CallGroup()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromString(group)
	ptrArguments[1] = gdnative.NewPointerFromString(method)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SceneTree", "call_group")

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

	Args: [{ false flags int} { false group String} { false method String}], Returns: Variant
*/

func (o *SceneTree) CallGroupFlags(flags gdnative.Int, group gdnative.String, method gdnative.String) gdnative.Variant {
	log.Println("Calling SceneTree.CallGroupFlags()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 3, 3)
	ptrArguments[0] = gdnative.NewPointerFromInt(flags)
	ptrArguments[1] = gdnative.NewPointerFromString(group)
	ptrArguments[2] = gdnative.NewPointerFromString(method)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SceneTree", "call_group_flags")

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

	Args: [{ false path String}], Returns: enum.Error
*/

/*

	Args: [{ false packed_scene PackedScene}], Returns: enum.Error
*/

/*

	Args: [{ false time_sec float} {True true pause_mode_process bool}], Returns: SceneTreeTimer
*/

func (o *SceneTree) CreateTimer(timeSec gdnative.Float, pauseModeProcess gdnative.Bool) SceneTreeTimer {
	log.Println("Calling SceneTree.CreateTimer()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromFloat(timeSec)
	ptrArguments[1] = gdnative.NewPointerFromBool(pauseModeProcess)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SceneTree", "create_timer")

	// Call the parent method.
	// SceneTreeTimer
	retPtr := NewEmptySceneTreeTimer()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := NewSceneTreeTimerFromPointer(retPtr)

	log.Println("  Got return value: ", ret)
	return ret
}

/*
        Undocumented
	Args: [], Returns: Node
*/

func (o *SceneTree) GetCurrentScene() node.Node {
	log.Println("Calling SceneTree.GetCurrentScene()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SceneTree", "get_current_scene")

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
        Undocumented
	Args: [], Returns: Node
*/

func (o *SceneTree) GetEditedSceneRoot() node.Node {
	log.Println("Calling SceneTree.GetEditedSceneRoot()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SceneTree", "get_edited_scene_root")

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

	Args: [], Returns: int
*/

func (o *SceneTree) GetFrame() gdnative.Int {
	log.Println("Calling SceneTree.GetFrame()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SceneTree", "get_frame")

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

	Args: [], Returns: PoolIntArray
*/

func (o *SceneTree) GetNetworkConnectedPeers() gdnative.PoolIntArray {
	log.Println("Calling SceneTree.GetNetworkConnectedPeers()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SceneTree", "get_network_connected_peers")

	// Call the parent method.
	// PoolIntArray
	retPtr := gdnative.NewEmptyPoolIntArray()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewPoolIntArrayFromPointer(retPtr)

	log.Println("  Got return value: ", ret)
	return ret
}

/*
        Undocumented
	Args: [], Returns: NetworkedMultiplayerPeer
*/

func (o *SceneTree) GetNetworkPeer() networkedmultiplayerpeer.NetworkedMultiplayerPeer {
	log.Println("Calling SceneTree.GetNetworkPeer()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SceneTree", "get_network_peer")

	// Call the parent method.
	// NetworkedMultiplayerPeer
	retPtr := networkedmultiplayerpeer.NewEmptyNetworkedMultiplayerPeer()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := networkedmultiplayerpeer.NewNetworkedMultiplayerPeerFromPointer(retPtr)

	log.Println("  Got return value: ", ret)
	return ret
}

/*

	Args: [], Returns: int
*/

func (o *SceneTree) GetNetworkUniqueId() gdnative.Int {
	log.Println("Calling SceneTree.GetNetworkUniqueId()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SceneTree", "get_network_unique_id")

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

	Args: [], Returns: int
*/

func (o *SceneTree) GetNodeCount() gdnative.Int {
	log.Println("Calling SceneTree.GetNodeCount()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SceneTree", "get_node_count")

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

	Args: [{ false group String}], Returns: Array
*/

func (o *SceneTree) GetNodesInGroup(group gdnative.String) gdnative.Array {
	log.Println("Calling SceneTree.GetNodesInGroup()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(group)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SceneTree", "get_nodes_in_group")

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
        Undocumented
	Args: [], Returns: Viewport
*/

func (o *SceneTree) GetRoot() viewport.Viewport {
	log.Println("Calling SceneTree.GetRoot()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SceneTree", "get_root")

	// Call the parent method.
	// Viewport
	retPtr := viewport.NewEmptyViewport()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := viewport.NewViewportFromPointer(retPtr)

	log.Println("  Got return value: ", ret)
	return ret
}

/*

	Args: [], Returns: int
*/

func (o *SceneTree) GetRpcSenderId() gdnative.Int {
	log.Println("Calling SceneTree.GetRpcSenderId()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SceneTree", "get_rpc_sender_id")

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

	Args: [{ false name String}], Returns: bool
*/

func (o *SceneTree) HasGroup(name gdnative.String) gdnative.Bool {
	log.Println("Calling SceneTree.HasGroup()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(name)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SceneTree", "has_group")

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
        Returns true if there is a [NetworkedMultiplayerPeer] set (with [method SceneTree.set_network_peer]).
	Args: [], Returns: bool
*/

func (o *SceneTree) HasNetworkPeer() gdnative.Bool {
	log.Println("Calling SceneTree.HasNetworkPeer()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SceneTree", "has_network_peer")

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
	Args: [], Returns: bool
*/

func (o *SceneTree) IsDebuggingCollisionsHint() gdnative.Bool {
	log.Println("Calling SceneTree.IsDebuggingCollisionsHint()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SceneTree", "is_debugging_collisions_hint")

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
	Args: [], Returns: bool
*/

func (o *SceneTree) IsDebuggingNavigationHint() gdnative.Bool {
	log.Println("Calling SceneTree.IsDebuggingNavigationHint()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SceneTree", "is_debugging_navigation_hint")

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

	Args: [], Returns: bool
*/

func (o *SceneTree) IsInputHandled() gdnative.Bool {
	log.Println("Calling SceneTree.IsInputHandled()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SceneTree", "is_input_handled")

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
        Returns true if this SceneTree's [NetworkedMultiplayerPeer] is in server mode (listening for connections).
	Args: [], Returns: bool
*/

func (o *SceneTree) IsNetworkServer() gdnative.Bool {
	log.Println("Calling SceneTree.IsNetworkServer()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SceneTree", "is_network_server")

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
	Args: [], Returns: bool
*/

func (o *SceneTree) IsPaused() gdnative.Bool {
	log.Println("Calling SceneTree.IsPaused()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SceneTree", "is_paused")

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
	Args: [], Returns: bool
*/

func (o *SceneTree) IsRefusingNewNetworkConnections() gdnative.Bool {
	log.Println("Calling SceneTree.IsRefusingNewNetworkConnections()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SceneTree", "is_refusing_new_network_connections")

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
	Args: [], Returns: bool
*/

func (o *SceneTree) IsUsingFontOversampling() gdnative.Bool {
	log.Println("Calling SceneTree.IsUsingFontOversampling()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SceneTree", "is_using_font_oversampling")

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

	Args: [{ false group String} { false notification int}], Returns: void
*/

func (o *SceneTree) NotifyGroup(group gdnative.String, notification gdnative.Int) {
	log.Println("Calling SceneTree.NotifyGroup()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromString(group)
	ptrArguments[1] = gdnative.NewPointerFromInt(notification)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SceneTree", "notify_group")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*

	Args: [{ false call_flags int} { false group String} { false notification int}], Returns: void
*/

func (o *SceneTree) NotifyGroupFlags(callFlags gdnative.Int, group gdnative.String, notification gdnative.Int) {
	log.Println("Calling SceneTree.NotifyGroupFlags()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 3, 3)
	ptrArguments[0] = gdnative.NewPointerFromInt(callFlags)
	ptrArguments[1] = gdnative.NewPointerFromString(group)
	ptrArguments[2] = gdnative.NewPointerFromInt(notification)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SceneTree", "notify_group_flags")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*

	Args: [{ false obj Object}], Returns: void
*/

func (o *SceneTree) QueueDelete(obj object.Object) {
	log.Println("Calling SceneTree.QueueDelete()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(obj.GetOwner())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SceneTree", "queue_delete")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*

	Args: [], Returns: void
*/

func (o *SceneTree) Quit() {
	log.Println("Calling SceneTree.Quit()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SceneTree", "quit")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*

	Args: [], Returns: enum.Error
*/

/*

	Args: [{ false enabled bool}], Returns: void
*/

func (o *SceneTree) SetAutoAcceptQuit(enabled gdnative.Bool) {
	log.Println("Calling SceneTree.SetAutoAcceptQuit()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(enabled)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SceneTree", "set_auto_accept_quit")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false child_node Object}], Returns: void
*/

func (o *SceneTree) SetCurrentScene(childNode object.Object) {
	log.Println("Calling SceneTree.SetCurrentScene()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(childNode.GetOwner())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SceneTree", "set_current_scene")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false enable bool}], Returns: void
*/

func (o *SceneTree) SetDebugCollisionsHint(enable gdnative.Bool) {
	log.Println("Calling SceneTree.SetDebugCollisionsHint()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(enable)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SceneTree", "set_debug_collisions_hint")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false enable bool}], Returns: void
*/

func (o *SceneTree) SetDebugNavigationHint(enable gdnative.Bool) {
	log.Println("Calling SceneTree.SetDebugNavigationHint()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(enable)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SceneTree", "set_debug_navigation_hint")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false scene Object}], Returns: void
*/

func (o *SceneTree) SetEditedSceneRoot(scene object.Object) {
	log.Println("Calling SceneTree.SetEditedSceneRoot()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(scene.GetOwner())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SceneTree", "set_edited_scene_root")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*

	Args: [{ false group String} { false property String} { false value Variant}], Returns: void
*/

func (o *SceneTree) SetGroup(group gdnative.String, property gdnative.String, value gdnative.Variant) {
	log.Println("Calling SceneTree.SetGroup()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 3, 3)
	ptrArguments[0] = gdnative.NewPointerFromString(group)
	ptrArguments[1] = gdnative.NewPointerFromString(property)
	ptrArguments[2] = gdnative.NewPointerFromVariant(value)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SceneTree", "set_group")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*

	Args: [{ false call_flags int} { false group String} { false property String} { false value Variant}], Returns: void
*/

func (o *SceneTree) SetGroupFlags(callFlags gdnative.Int, group gdnative.String, property gdnative.String, value gdnative.Variant) {
	log.Println("Calling SceneTree.SetGroupFlags()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 4, 4)
	ptrArguments[0] = gdnative.NewPointerFromInt(callFlags)
	ptrArguments[1] = gdnative.NewPointerFromString(group)
	ptrArguments[2] = gdnative.NewPointerFromString(property)
	ptrArguments[3] = gdnative.NewPointerFromVariant(value)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SceneTree", "set_group_flags")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*

	Args: [], Returns: void
*/

func (o *SceneTree) SetInputAsHandled() {
	log.Println("Calling SceneTree.SetInputAsHandled()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SceneTree", "set_input_as_handled")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false peer NetworkedMultiplayerPeer}], Returns: void
*/

func (o *SceneTree) SetNetworkPeer(peer networkedmultiplayerpeer.NetworkedMultiplayerPeer) {
	log.Println("Calling SceneTree.SetNetworkPeer()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(peer.GetOwner())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SceneTree", "set_network_peer")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false enable bool}], Returns: void
*/

func (o *SceneTree) SetPause(enable gdnative.Bool) {
	log.Println("Calling SceneTree.SetPause()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(enable)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SceneTree", "set_pause")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*

	Args: [{ false enabled bool}], Returns: void
*/

func (o *SceneTree) SetQuitOnGoBack(enabled gdnative.Bool) {
	log.Println("Calling SceneTree.SetQuitOnGoBack()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(enabled)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SceneTree", "set_quit_on_go_back")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false refuse bool}], Returns: void
*/

func (o *SceneTree) SetRefuseNewNetworkConnections(refuse gdnative.Bool) {
	log.Println("Calling SceneTree.SetRefuseNewNetworkConnections()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(refuse)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SceneTree", "set_refuse_new_network_connections")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*

	Args: [{ false mode int} { false aspect int} { false minsize Vector2} {1 true shrink float}], Returns: void
*/

func (o *SceneTree) SetScreenStretch(mode gdnative.Int, aspect gdnative.Int, minsize gdnative.Vector2, shrink gdnative.Float) {
	log.Println("Calling SceneTree.SetScreenStretch()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 4, 4)
	ptrArguments[0] = gdnative.NewPointerFromInt(mode)
	ptrArguments[1] = gdnative.NewPointerFromInt(aspect)
	ptrArguments[2] = gdnative.NewPointerFromVector2(minsize)
	ptrArguments[3] = gdnative.NewPointerFromFloat(shrink)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SceneTree", "set_screen_stretch")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false enable bool}], Returns: void
*/

func (o *SceneTree) SetUseFontOversampling(enable gdnative.Bool) {
	log.Println("Calling SceneTree.SetUseFontOversampling()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(enable)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SceneTree", "set_use_font_oversampling")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}
