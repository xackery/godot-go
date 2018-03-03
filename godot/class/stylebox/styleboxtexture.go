package stylebox

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

func NewStyleBoxTextureFromPointer(ptr gdnative.Pointer) *StyleBoxTexture {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := StyleBoxTexture{}
	obj.SetOwner(owner)

	return &obj

}

/*
Texture Based 3x3 scale style. This stylebox performs a 3x3 scaling of a texture, where only the center cell is fully stretched. This allows for the easy creation of bordered styles.
*/
type StyleBoxTexture struct {
	StyleBox
}

func (o *StyleBoxTexture) BaseClass() string {
	return "StyleBoxTexture"
}

/*
        Undocumented
	Args: [{ false margin int}], Returns: float
*/

func (o *StyleBoxTexture) GetExpandMarginSize(margin gdnative.Int) gdnative.Float {
	log.Println("Calling StyleBoxTexture.GetExpandMarginSize()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(margin)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("StyleBoxTexture", "get_expand_margin_size")

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
	Args: [], Returns: enum.StyleBoxTexture::AxisStretchMode
*/

/*
        Undocumented
	Args: [{ false margin int}], Returns: float
*/

func (o *StyleBoxTexture) GetMarginSize(margin gdnative.Int) gdnative.Float {
	log.Println("Calling StyleBoxTexture.GetMarginSize()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(margin)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("StyleBoxTexture", "get_margin_size")

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
	Args: [], Returns: Color
*/

func (o *StyleBoxTexture) GetModulate() gdnative.Color {
	log.Println("Calling StyleBoxTexture.GetModulate()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("StyleBoxTexture", "get_modulate")

	// Call the parent method.
	// Color
	retPtr := gdnative.NewEmptyColor()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewColorFromPointer(retPtr)

	log.Println("  Got return value: ", ret)
	return ret
}

/*
        Undocumented
	Args: [], Returns: Resource
*/

func (o *StyleBoxTexture) GetNormalMap() resource.Resource {
	log.Println("Calling StyleBoxTexture.GetNormalMap()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("StyleBoxTexture", "get_normal_map")

	// Call the parent method.
	// Resource
	retPtr := resource.NewEmptyResource()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := resource.NewResourceFromPointer(retPtr)

	log.Println("  Got return value: ", ret)
	return ret
}

/*
        Undocumented
	Args: [], Returns: Rect2
*/

func (o *StyleBoxTexture) GetRegionRect() gdnative.Rect2 {
	log.Println("Calling StyleBoxTexture.GetRegionRect()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("StyleBoxTexture", "get_region_rect")

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
	Args: [], Returns: Resource
*/

func (o *StyleBoxTexture) GetTexture() resource.Resource {
	log.Println("Calling StyleBoxTexture.GetTexture()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("StyleBoxTexture", "get_texture")

	// Call the parent method.
	// Resource
	retPtr := resource.NewEmptyResource()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := resource.NewResourceFromPointer(retPtr)

	log.Println("  Got return value: ", ret)
	return ret
}

/*
        Undocumented
	Args: [], Returns: enum.StyleBoxTexture::AxisStretchMode
*/

/*
        Undocumented
	Args: [], Returns: bool
*/

func (o *StyleBoxTexture) IsDrawCenterEnabled() gdnative.Bool {
	log.Println("Calling StyleBoxTexture.IsDrawCenterEnabled()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("StyleBoxTexture", "is_draw_center_enabled")

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
	Args: [{ false enable bool}], Returns: void
*/

func (o *StyleBoxTexture) SetDrawCenter(enable gdnative.Bool) {
	log.Println("Calling StyleBoxTexture.SetDrawCenter()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(enable)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("StyleBoxTexture", "set_draw_center")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*

	Args: [{ false size float}], Returns: void
*/

func (o *StyleBoxTexture) SetExpandMarginAll(size gdnative.Float) {
	log.Println("Calling StyleBoxTexture.SetExpandMarginAll()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromFloat(size)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("StyleBoxTexture", "set_expand_margin_all")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*

	Args: [{ false size_left float} { false size_top float} { false size_right float} { false size_bottom float}], Returns: void
*/

func (o *StyleBoxTexture) SetExpandMarginIndividual(sizeLeft gdnative.Float, sizeTop gdnative.Float, sizeRight gdnative.Float, sizeBottom gdnative.Float) {
	log.Println("Calling StyleBoxTexture.SetExpandMarginIndividual()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 4, 4)
	ptrArguments[0] = gdnative.NewPointerFromFloat(sizeLeft)
	ptrArguments[1] = gdnative.NewPointerFromFloat(sizeTop)
	ptrArguments[2] = gdnative.NewPointerFromFloat(sizeRight)
	ptrArguments[3] = gdnative.NewPointerFromFloat(sizeBottom)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("StyleBoxTexture", "set_expand_margin_individual")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false margin int} { false size float}], Returns: void
*/

func (o *StyleBoxTexture) SetExpandMarginSize(margin gdnative.Int, size gdnative.Float) {
	log.Println("Calling StyleBoxTexture.SetExpandMarginSize()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(margin)
	ptrArguments[1] = gdnative.NewPointerFromFloat(size)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("StyleBoxTexture", "set_expand_margin_size")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false mode int}], Returns: void
*/

func (o *StyleBoxTexture) SetHAxisStretchMode(mode gdnative.Int) {
	log.Println("Calling StyleBoxTexture.SetHAxisStretchMode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(mode)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("StyleBoxTexture", "set_h_axis_stretch_mode")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false margin int} { false size float}], Returns: void
*/

func (o *StyleBoxTexture) SetMarginSize(margin gdnative.Int, size gdnative.Float) {
	log.Println("Calling StyleBoxTexture.SetMarginSize()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(margin)
	ptrArguments[1] = gdnative.NewPointerFromFloat(size)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("StyleBoxTexture", "set_margin_size")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false color Color}], Returns: void
*/

func (o *StyleBoxTexture) SetModulate(color gdnative.Color) {
	log.Println("Calling StyleBoxTexture.SetModulate()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromColor(color)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("StyleBoxTexture", "set_modulate")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false normal_map Resource}], Returns: void
*/

func (o *StyleBoxTexture) SetNormalMap(normalMap resource.Resource) {
	log.Println("Calling StyleBoxTexture.SetNormalMap()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(normalMap.GetOwner())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("StyleBoxTexture", "set_normal_map")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false region Rect2}], Returns: void
*/

func (o *StyleBoxTexture) SetRegionRect(region gdnative.Rect2) {
	log.Println("Calling StyleBoxTexture.SetRegionRect()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromRect2(region)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("StyleBoxTexture", "set_region_rect")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false texture Resource}], Returns: void
*/

func (o *StyleBoxTexture) SetTexture(texture resource.Resource) {
	log.Println("Calling StyleBoxTexture.SetTexture()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(texture.GetOwner())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("StyleBoxTexture", "set_texture")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false mode int}], Returns: void
*/

func (o *StyleBoxTexture) SetVAxisStretchMode(mode gdnative.Int) {
	log.Println("Calling StyleBoxTexture.SetVAxisStretchMode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(mode)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("StyleBoxTexture", "set_v_axis_stretch_mode")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}
