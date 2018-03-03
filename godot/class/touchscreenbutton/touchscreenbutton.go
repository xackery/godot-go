package touchscreenbutton

import (
	"log"

	"github.com/shadowapex/godot-go/gdnative"

	"github.com/shadowapex/godot-go/godot/class/node2d"
	"github.com/shadowapex/godot-go/godot/class/texture"

	"github.com/shadowapex/godot-go/godot/class/inputevent"

	"github.com/shadowapex/godot-go/godot/class/bitmap"

	"github.com/shadowapex/godot-go/godot/class/shape2d"
)

/*------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "class.go.tmpl" so they can be included in the generated
//   code.
//----------------------------------------------------------------------------*/

func NewTouchScreenButtonFromPointer(ptr gdnative.Pointer) *TouchScreenButton {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := TouchScreenButton{}
	obj.SetOwner(owner)

	return &obj

}

/*
Button for touch screen devices. You can set it to be visible on all screens, or only on touch devices.
*/
type TouchScreenButton struct {
	node2d.Node2D
}

func (o *TouchScreenButton) BaseClass() string {
	return "TouchScreenButton"
}

/*
        Undocumented
	Args: [{ false arg0 InputEvent}], Returns: void
*/

func (o *TouchScreenButton) X_Input(arg0 inputevent.InputEvent) {
	log.Println("Calling TouchScreenButton.X_Input()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(arg0.GetOwner())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TouchScreenButton", "_input")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: String
*/

func (o *TouchScreenButton) GetAction() gdnative.String {
	log.Println("Calling TouchScreenButton.GetAction()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TouchScreenButton", "get_action")

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
	Args: [], Returns: BitMap
*/

func (o *TouchScreenButton) GetBitmask() bitmap.BitMap {
	log.Println("Calling TouchScreenButton.GetBitmask()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TouchScreenButton", "get_bitmask")

	// Call the parent method.
	// BitMap
	retPtr := bitmap.NewEmptyBitMap()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := bitmap.NewBitMapFromPointer(retPtr)

	log.Println("  Got return value: ", ret)
	return ret
}

/*
        Undocumented
	Args: [], Returns: Shape2D
*/

func (o *TouchScreenButton) GetShape() shape2d.Shape2D {
	log.Println("Calling TouchScreenButton.GetShape()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TouchScreenButton", "get_shape")

	// Call the parent method.
	// Shape2D
	retPtr := shape2d.NewEmptyShape2D()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := shape2d.NewShape2DFromPointer(retPtr)

	log.Println("  Got return value: ", ret)
	return ret
}

/*
        Undocumented
	Args: [], Returns: Texture
*/

func (o *TouchScreenButton) GetTexture() texture.Texture {
	log.Println("Calling TouchScreenButton.GetTexture()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TouchScreenButton", "get_texture")

	// Call the parent method.
	// Texture
	retPtr := texture.NewEmptyTexture()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := texture.NewTextureFromPointer(retPtr)

	log.Println("  Got return value: ", ret)
	return ret
}

/*
        Undocumented
	Args: [], Returns: Texture
*/

func (o *TouchScreenButton) GetTexturePressed() texture.Texture {
	log.Println("Calling TouchScreenButton.GetTexturePressed()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TouchScreenButton", "get_texture_pressed")

	// Call the parent method.
	// Texture
	retPtr := texture.NewEmptyTexture()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := texture.NewTextureFromPointer(retPtr)

	log.Println("  Got return value: ", ret)
	return ret
}

/*
        Undocumented
	Args: [], Returns: enum.TouchScreenButton::VisibilityMode
*/

/*
        Undocumented
	Args: [], Returns: bool
*/

func (o *TouchScreenButton) IsPassbyPressEnabled() gdnative.Bool {
	log.Println("Calling TouchScreenButton.IsPassbyPressEnabled()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TouchScreenButton", "is_passby_press_enabled")

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
        Returns [code]true[/code] if this button is currently pressed.
	Args: [], Returns: bool
*/

func (o *TouchScreenButton) IsPressed() gdnative.Bool {
	log.Println("Calling TouchScreenButton.IsPressed()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TouchScreenButton", "is_pressed")

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

func (o *TouchScreenButton) IsShapeCentered() gdnative.Bool {
	log.Println("Calling TouchScreenButton.IsShapeCentered()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TouchScreenButton", "is_shape_centered")

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

func (o *TouchScreenButton) IsShapeVisible() gdnative.Bool {
	log.Println("Calling TouchScreenButton.IsShapeVisible()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TouchScreenButton", "is_shape_visible")

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
	Args: [{ false action String}], Returns: void
*/

func (o *TouchScreenButton) SetAction(action gdnative.String) {
	log.Println("Calling TouchScreenButton.SetAction()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(action)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TouchScreenButton", "set_action")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false bitmask BitMap}], Returns: void
*/

func (o *TouchScreenButton) SetBitmask(bitmask bitmap.BitMap) {
	log.Println("Calling TouchScreenButton.SetBitmask()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(bitmask.GetOwner())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TouchScreenButton", "set_bitmask")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false enabled bool}], Returns: void
*/

func (o *TouchScreenButton) SetPassbyPress(enabled gdnative.Bool) {
	log.Println("Calling TouchScreenButton.SetPassbyPress()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(enabled)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TouchScreenButton", "set_passby_press")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false shape Shape2D}], Returns: void
*/

func (o *TouchScreenButton) SetShape(shape shape2d.Shape2D) {
	log.Println("Calling TouchScreenButton.SetShape()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(shape.GetOwner())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TouchScreenButton", "set_shape")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false bool bool}], Returns: void
*/

func (o *TouchScreenButton) SetShapeCentered(bool gdnative.Bool) {
	log.Println("Calling TouchScreenButton.SetShapeCentered()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(bool)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TouchScreenButton", "set_shape_centered")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false bool bool}], Returns: void
*/

func (o *TouchScreenButton) SetShapeVisible(bool gdnative.Bool) {
	log.Println("Calling TouchScreenButton.SetShapeVisible()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(bool)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TouchScreenButton", "set_shape_visible")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false texture Texture}], Returns: void
*/

func (o *TouchScreenButton) SetTexture(texture texture.Texture) {
	log.Println("Calling TouchScreenButton.SetTexture()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(texture.GetOwner())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TouchScreenButton", "set_texture")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false texture_pressed Texture}], Returns: void
*/

func (o *TouchScreenButton) SetTexturePressed(texturePressed texture.Texture) {
	log.Println("Calling TouchScreenButton.SetTexturePressed()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(texturePressed.GetOwner())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TouchScreenButton", "set_texture_pressed")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false mode int}], Returns: void
*/

func (o *TouchScreenButton) SetVisibilityMode(mode gdnative.Int) {
	log.Println("Calling TouchScreenButton.SetVisibilityMode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(mode)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TouchScreenButton", "set_visibility_mode")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}
