package streampeer

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

func NewStreamPeerTCPFromPointer(ptr gdnative.Pointer) *StreamPeerTCP {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := StreamPeerTCP{}
	obj.SetOwner(owner)

	return &obj

}

/*
TCP Stream peer. This object can be used to connect to TCP servers, or also is returned by a tcp server.
*/
type StreamPeerTCP struct {
	StreamPeer
}

func (o *StreamPeerTCP) BaseClass() string {
	return "StreamPeerTCP"
}

/*
        Connect to the specified host:port pair. A hostname will be resolved if valid. Returns [OK] on success or [FAILED] on failure.
	Args: [{ false host String} { false port int}], Returns: enum.Error
*/

/*
        Disconnect from host.
	Args: [], Returns: void
*/

func (o *StreamPeerTCP) DisconnectFromHost() {
	log.Println("Calling StreamPeerTCP.DisconnectFromHost()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("StreamPeerTCP", "disconnect_from_host")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*
        Return the IP of this peer.
	Args: [], Returns: String
*/

func (o *StreamPeerTCP) GetConnectedHost() gdnative.String {
	log.Println("Calling StreamPeerTCP.GetConnectedHost()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("StreamPeerTCP", "get_connected_host")

	// Call the parent method.
	// String
	retPtr := gdnative.NewEmptyString()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewStringFromPointer(retPtr)

	log.Println("  Got return value: ", ret)
	return ret
}

/*
        Return the port of this peer.
	Args: [], Returns: int
*/

func (o *StreamPeerTCP) GetConnectedPort() gdnative.Int {
	log.Println("Calling StreamPeerTCP.GetConnectedPort()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("StreamPeerTCP", "get_connected_port")

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
        Return the status of the connection, one of STATUS_* enum.
	Args: [], Returns: enum.StreamPeerTCP::Status
*/

/*

	Args: [], Returns: bool
*/

func (o *StreamPeerTCP) IsConnectedToHost() gdnative.Bool {
	log.Println("Calling StreamPeerTCP.IsConnectedToHost()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("StreamPeerTCP", "is_connected_to_host")

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
        Disable Nagle algorithm to improve latency for small packets. Note that for applications that send large packets, or need to transfer a lot of data, this can reduce total bandwidth.
	Args: [{ false enabled bool}], Returns: void
*/

func (o *StreamPeerTCP) SetNoDelay(enabled gdnative.Bool) {
	log.Println("Calling StreamPeerTCP.SetNoDelay()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(enabled)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("StreamPeerTCP", "set_no_delay")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}
