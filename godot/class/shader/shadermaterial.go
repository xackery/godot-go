package shader

import (
	"log"

	"github.com/shadowapex/godot-go/gdnative"
	"github.com/shadowapex/godot-go/godot/class/material"
)

/*------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "class.go.tmpl" so they can be included in the generated
//   code.
//----------------------------------------------------------------------------*/

func NewShaderMaterialFromPointer(ptr gdnative.Pointer) *ShaderMaterial {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := ShaderMaterial{}
	obj.SetOwner(owner)

	return &obj

}

/*

 */
type ShaderMaterial struct {
	material.Material
}

func (o *ShaderMaterial) BaseClass() string {
	return "ShaderMaterial"
}

/*
        Undocumented
	Args: [], Returns: Shader
*/

func (o *ShaderMaterial) GetShader() Shader {
	log.Println("Calling ShaderMaterial.GetShader()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ShaderMaterial", "get_shader")

	// Call the parent method.
	// Shader
	retPtr := NewEmptyShader()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := NewShaderFromPointer(retPtr)

	log.Println("  Got return value: ", ret)
	return ret
}

/*

	Args: [{ false param String}], Returns: Variant
*/

func (o *ShaderMaterial) GetShaderParam(param gdnative.String) gdnative.Variant {
	log.Println("Calling ShaderMaterial.GetShaderParam()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(param)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ShaderMaterial", "get_shader_param")

	// Call the parent method.
	// Variant
	retPtr := gdnative.NewEmptyVariant()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVariantFromPointer(retPtr)

	log.Println("  Got return value: ", ret)
	return ret
}

/*
        Undocumented
	Args: [{ false shader Shader}], Returns: void
*/

func (o *ShaderMaterial) SetShader(shader Shader) {
	log.Println("Calling ShaderMaterial.SetShader()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(shader.GetOwner())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ShaderMaterial", "set_shader")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*

	Args: [{ false param String} { false value Variant}], Returns: void
*/

func (o *ShaderMaterial) SetShaderParam(param gdnative.String, value gdnative.Variant) {
	log.Println("Calling ShaderMaterial.SetShaderParam()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromString(param)
	ptrArguments[1] = gdnative.NewPointerFromVariant(value)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ShaderMaterial", "set_shader_param")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}
