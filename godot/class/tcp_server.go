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

//func NewTCP_ServerFromPointer(ptr gdnative.Pointer) TCP_Server {
func NewTCP_ServerFromPointer(ptr gdnative.Pointer) TCP_Server {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := TCP_Server{}
	obj.SetBaseObject(owner)

	return obj
}

/*
TCP Server class. Listens to connections on a port and returns a [StreamPeerTCP] when got a connection.
*/
type TCP_Server struct {
	Reference
	owner gdnative.Object
}

func (o *TCP_Server) BaseClass() string {
	return "TCP_Server"
}

// SetBaseObject will internally set the Godot object inside the struct.
// This is used to call parent methods.
func (o *TCP_Server) SetBaseObject(object gdnative.Object) {
	o.owner = object
}

func (o *TCP_Server) GetBaseObject() gdnative.Object {
	return o.owner
}

/*
        Return true if a connection is available for taking.
	Args: [], Returns: bool
*/
func (o *TCP_Server) IsConnectionAvailable() gdnative.Bool {
	log.Println("Calling TCP_Server.IsConnectionAvailable()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TCP_Server", "is_connection_available")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	log.Println("  Got return value: ", ret)
	return ret
}

/*
        Listen on the "port" binding to "bind_address". If "bind_address" is set as "*" (default), the server will listen on all available addresses (both IPv4 and IPv6). If "bind_address" is set as "0.0.0.0" (for IPv4) or "::" (for IPv6), the server will listen on all available addresses matching that IP type. If "bind_address" is set to any valid address (e.g. "192.168.1.101", "::1", etc), the server will only listen on the interface with that addresses (or fail if no interface with the given address exists).
	Args: [{ false port int} {* true bind_address String}], Returns: enum.Error
*/

/*
        Stop listening.
	Args: [], Returns: void
*/
func (o *TCP_Server) Stop() {
	log.Println("Calling TCP_Server.Stop()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TCP_Server", "stop")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        If a connection is available, return a StreamPeerTCP with the connection/
	Args: [], Returns: StreamPeerTCP
*/
func (o *TCP_Server) TakeConnection() StreamPeerTCP {
	log.Println("Calling TCP_Server.TakeConnection()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TCP_Server", "take_connection")

	// Call the parent method.
	// StreamPeerTCP
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := NewStreamPeerTCPFromPointer(retPtr)
	log.Println("  Got return value: ", ret)
	return ret
}