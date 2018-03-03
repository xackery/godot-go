package audiostream

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

func NewAudioStreamFromPointer(ptr gdnative.Pointer) *AudioStream {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := AudioStream{}
	obj.SetOwner(owner)

	return &obj

}

/*
Base class for audio streams. Audio streams are used for music playback, or other types of streamed sounds that don't fit or require more flexibility than a [Sample].
*/
type AudioStream struct {
	resource.Resource
}

func (o *AudioStream) BaseClass() string {
	return "AudioStream"
}

/*

	Args: [], Returns: float
*/

func (o *AudioStream) GetLength() gdnative.Float {
	log.Println("Calling AudioStream.GetLength()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioStream", "get_length")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyFloat()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewFloatFromPointer(retPtr)

	log.Println("  Got return value: ", ret)
	return ret
}
