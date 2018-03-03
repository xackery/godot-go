package inputevent

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

func NewInputEventMagnifyGestureFromPointer(ptr gdnative.Pointer) *InputEventMagnifyGesture {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := InputEventMagnifyGesture{}
	obj.SetOwner(owner)

	return &obj

}

/*

 */
type InputEventMagnifyGesture struct {
	InputEventGesture
}

func (o *InputEventMagnifyGesture) BaseClass() string {
	return "InputEventMagnifyGesture"
}

/*
        Undocumented
	Args: [], Returns: float
*/

func (o *InputEventMagnifyGesture) GetFactor() gdnative.Float {
	log.Println("Calling InputEventMagnifyGesture.GetFactor()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("InputEventMagnifyGesture", "get_factor")

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
	Args: [{ false factor float}], Returns: void
*/

func (o *InputEventMagnifyGesture) SetFactor(factor gdnative.Float) {
	log.Println("Calling InputEventMagnifyGesture.SetFactor()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromFloat(factor)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("InputEventMagnifyGesture", "set_factor")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}
