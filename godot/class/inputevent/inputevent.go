package inputevent

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

func NewInputEventFromPointer(ptr gdnative.Pointer) *InputEvent {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := InputEvent{}
	obj.SetOwner(owner)

	return &obj

}

/*
Base class of all sort of input event. See [method Node._input].
*/
type InputEvent struct {
	resource.Resource
}

func (o *InputEvent) BaseClass() string {
	return "InputEvent"
}

/*
        Returns [code]true[/code] if this event matches [code]event[/code].
	Args: [{ false event InputEvent}], Returns: bool
*/

func (o *InputEvent) ActionMatch(event InputEvent) gdnative.Bool {
	log.Println("Calling InputEvent.ActionMatch()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(event.GetOwner())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("InputEvent", "action_match")

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
        Returns a [String] representation of the event.
	Args: [], Returns: String
*/

func (o *InputEvent) AsText() gdnative.String {
	log.Println("Calling InputEvent.AsText()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("InputEvent", "as_text")

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
        Undocumented
	Args: [], Returns: int
*/

func (o *InputEvent) GetDevice() gdnative.Int {
	log.Println("Calling InputEvent.GetDevice()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("InputEvent", "get_device")

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
        Returns [code]true[/code] if this input event matches a pre-defined action of any type.
	Args: [{ false action String}], Returns: bool
*/

func (o *InputEvent) IsAction(action gdnative.String) gdnative.Bool {
	log.Println("Calling InputEvent.IsAction()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(action)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("InputEvent", "is_action")

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
        Returns [code]true[/code] if the given action is being pressed (and is not an echo event for KEY events). Not relevant for the event types [code]MOUSE_MOTION[/code], [code]SCREEN_DRAG[/code] or [code]NONE[/code].
	Args: [{ false action String}], Returns: bool
*/

func (o *InputEvent) IsActionPressed(action gdnative.String) gdnative.Bool {
	log.Println("Calling InputEvent.IsActionPressed()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(action)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("InputEvent", "is_action_pressed")

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
        Returns [code]true[/code] if the given action is released (i.e. not pressed). Not relevant for the event types [code]MOUSE_MOTION[/code], [code]SCREEN_DRAG[/code] or [code]NONE[/code].
	Args: [{ false action String}], Returns: bool
*/

func (o *InputEvent) IsActionReleased(action gdnative.String) gdnative.Bool {
	log.Println("Calling InputEvent.IsActionReleased()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(action)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("InputEvent", "is_action_released")

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
        Returns [code]true[/code] if this input event's type is one of the [code]InputEvent[/code] constants.
	Args: [], Returns: bool
*/

func (o *InputEvent) IsActionType() gdnative.Bool {
	log.Println("Calling InputEvent.IsActionType()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("InputEvent", "is_action_type")

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
        Returns [code]true[/code] if this input event is an echo event (only for events of type KEY).
	Args: [], Returns: bool
*/

func (o *InputEvent) IsEcho() gdnative.Bool {
	log.Println("Calling InputEvent.IsEcho()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("InputEvent", "is_echo")

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
        Returns [code]true[/code] if this input event is pressed. Not relevant for the event types [code]MOUSE_MOTION[/code], [code]SCREEN_DRAG[/code] or [code]NONE[/code].
	Args: [], Returns: bool
*/

func (o *InputEvent) IsPressed() gdnative.Bool {
	log.Println("Calling InputEvent.IsPressed()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("InputEvent", "is_pressed")

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
	Args: [{ false device int}], Returns: void
*/

func (o *InputEvent) SetDevice(device gdnative.Int) {
	log.Println("Calling InputEvent.SetDevice()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(device)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("InputEvent", "set_device")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*

	Args: [{ false event InputEvent}], Returns: bool
*/

func (o *InputEvent) ShortcutMatch(event InputEvent) gdnative.Bool {
	log.Println("Calling InputEvent.ShortcutMatch()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(event.GetOwner())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("InputEvent", "shortcut_match")

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

	Args: [{ false xform Transform2D} {(0, 0) true local_ofs Vector2}], Returns: InputEvent
*/

func (o *InputEvent) XformedBy(xform gdnative.Transform2D, localOfs gdnative.Vector2) InputEvent {
	log.Println("Calling InputEvent.XformedBy()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromTransform2D(xform)
	ptrArguments[1] = gdnative.NewPointerFromVector2(localOfs)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("InputEvent", "xformed_by")

	// Call the parent method.
	// InputEvent
	retPtr := NewEmptyInputEvent()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := NewInputEventFromPointer(retPtr)

	log.Println("  Got return value: ", ret)
	return ret
}
