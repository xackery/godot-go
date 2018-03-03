package ninepatchrect

import (
	"log"

	"github.com/shadowapex/godot-go/gdnative"

	"github.com/shadowapex/godot-go/godot/class/control"
	"github.com/shadowapex/godot-go/godot/class/texture"
)

/*------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "class.go.tmpl" so they can be included in the generated
//   code.
//----------------------------------------------------------------------------*/

func NewNinePatchRectFromPointer(ptr gdnative.Pointer) *NinePatchRect {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := NinePatchRect{}
	obj.SetOwner(owner)

	return &obj

}

/*
Better known as 9-slice panels, NinePatchRect produces clean panels of any size, based on a small texture. To do so, it splits the texture in a 3 by 3 grid. When you scale the node, it tiles the texture's sides horizontally or vertically, the center on both axes but it doesn't scale or tile the corners.
*/
type NinePatchRect struct {
	control.Control
}

func (o *NinePatchRect) BaseClass() string {
	return "NinePatchRect"
}

/*
        Undocumented
	Args: [], Returns: enum.NinePatchRect::AxisStretchMode
*/

/*
        Undocumented
	Args: [{ false margin int}], Returns: int
*/

func (o *NinePatchRect) GetPatchMargin(margin gdnative.Int) gdnative.Int {
	log.Println("Calling NinePatchRect.GetPatchMargin()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(margin)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("NinePatchRect", "get_patch_margin")

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
	Args: [], Returns: Rect2
*/

func (o *NinePatchRect) GetRegionRect() gdnative.Rect2 {
	log.Println("Calling NinePatchRect.GetRegionRect()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("NinePatchRect", "get_region_rect")

	// Call the parent method.
	// Rect2
	retPtr := gdnative.NewEmptyRect2()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewRect2FromPointer(retPtr)

	log.Println("  Got return value: ", ret)
	return ret
}

/*
        Undocumented
	Args: [], Returns: Texture
*/

func (o *NinePatchRect) GetTexture() texture.Texture {
	log.Println("Calling NinePatchRect.GetTexture()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("NinePatchRect", "get_texture")

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
	Args: [], Returns: enum.NinePatchRect::AxisStretchMode
*/

/*
        Undocumented
	Args: [], Returns: bool
*/

func (o *NinePatchRect) IsDrawCenterEnabled() gdnative.Bool {
	log.Println("Calling NinePatchRect.IsDrawCenterEnabled()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("NinePatchRect", "is_draw_center_enabled")

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
	Args: [{ false draw_center bool}], Returns: void
*/

func (o *NinePatchRect) SetDrawCenter(drawCenter gdnative.Bool) {
	log.Println("Calling NinePatchRect.SetDrawCenter()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(drawCenter)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("NinePatchRect", "set_draw_center")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false mode int}], Returns: void
*/

func (o *NinePatchRect) SetHAxisStretchMode(mode gdnative.Int) {
	log.Println("Calling NinePatchRect.SetHAxisStretchMode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(mode)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("NinePatchRect", "set_h_axis_stretch_mode")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false margin int} { false value int}], Returns: void
*/

func (o *NinePatchRect) SetPatchMargin(margin gdnative.Int, value gdnative.Int) {
	log.Println("Calling NinePatchRect.SetPatchMargin()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(margin)
	ptrArguments[1] = gdnative.NewPointerFromInt(value)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("NinePatchRect", "set_patch_margin")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false rect Rect2}], Returns: void
*/

func (o *NinePatchRect) SetRegionRect(rect gdnative.Rect2) {
	log.Println("Calling NinePatchRect.SetRegionRect()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromRect2(rect)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("NinePatchRect", "set_region_rect")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false texture Texture}], Returns: void
*/

func (o *NinePatchRect) SetTexture(texture texture.Texture) {
	log.Println("Calling NinePatchRect.SetTexture()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(texture.GetOwner())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("NinePatchRect", "set_texture")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false mode int}], Returns: void
*/

func (o *NinePatchRect) SetVAxisStretchMode(mode gdnative.Int) {
	log.Println("Calling NinePatchRect.SetVAxisStretchMode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(mode)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("NinePatchRect", "set_v_axis_stretch_mode")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}
