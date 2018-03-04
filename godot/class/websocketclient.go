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

//func NewWebSocketClientFromPointer(ptr gdnative.Pointer) WebSocketClient {
func NewWebSocketClientFromPointer(ptr gdnative.Pointer) WebSocketClient {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := WebSocketClient{}
	obj.SetBaseObject(owner)

	return obj
}

/*

 */
type WebSocketClient struct {
	WebSocketMultiplayerPeer
	owner gdnative.Object
}

func (o *WebSocketClient) BaseClass() string {
	return "WebSocketClient"
}

// SetBaseObject will internally set the Godot object inside the struct.
// This is used to call parent methods.
func (o *WebSocketClient) SetBaseObject(object gdnative.Object) {
	o.owner = object
}

func (o *WebSocketClient) GetBaseObject() gdnative.Object {
	return o.owner
}

/*

	Args: [{ false url String} {[] true protocols PoolStringArray} {False true gd_mp_api bool}], Returns: enum.Error
*/

/*

	Args: [], Returns: void
*/
func (o *WebSocketClient) DisconnectFromHost() {
	log.Println("Calling WebSocketClient.DisconnectFromHost()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("WebSocketClient", "disconnect_from_host")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}