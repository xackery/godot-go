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

func NewInputEventKeyFromPointer(ptr gdnative.Pointer) *InputEventKey {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := InputEventKey{}
	obj.SetOwner(owner)

	return &obj

}

/*
Stores key presses on the keyboard. Supports key presses, key releases and [member echo] events.
*/
type InputEventKey struct {
	InputEventWithModifiers
}

func (o *InputEventKey) BaseClass() string {
	return "InputEventKey"
}

/*
        Undocumented
	Args: [], Returns: int
*/

func (o *InputEventKey) GetScancode() gdnative.Int {
	log.Println("Calling InputEventKey.GetScancode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("InputEventKey", "get_scancode")

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

func (o *InputEventKey) GetScancodeWithModifiers() gdnative.Int {
	log.Println("Calling InputEventKey.GetScancodeWithModifiers()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("InputEventKey", "get_scancode_with_modifiers")

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
        Undocumented
	Args: [], Returns: int
*/

func (o *InputEventKey) GetUnicode() gdnative.Int {
	log.Println("Calling InputEventKey.GetUnicode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("InputEventKey", "get_unicode")

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
        Undocumented
	Args: [{ false echo bool}], Returns: void
*/

func (o *InputEventKey) SetEcho(echo gdnative.Bool) {
	log.Println("Calling InputEventKey.SetEcho()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(echo)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("InputEventKey", "set_echo")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false pressed bool}], Returns: void
*/

func (o *InputEventKey) SetPressed(pressed gdnative.Bool) {
	log.Println("Calling InputEventKey.SetPressed()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(pressed)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("InputEventKey", "set_pressed")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false scancode int}], Returns: void
*/

func (o *InputEventKey) SetScancode(scancode gdnative.Int) {
	log.Println("Calling InputEventKey.SetScancode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(scancode)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("InputEventKey", "set_scancode")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false unicode int}], Returns: void
*/

func (o *InputEventKey) SetUnicode(unicode gdnative.Int) {
	log.Println("Calling InputEventKey.SetUnicode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(unicode)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("InputEventKey", "set_unicode")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}
