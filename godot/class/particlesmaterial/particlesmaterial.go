package particlesmaterial

import (
	"log"

	"github.com/shadowapex/godot-go/gdnative"

	"github.com/shadowapex/godot-go/godot/class/material"
	"github.com/shadowapex/godot-go/godot/class/texture"

	"github.com/shadowapex/godot-go/godot/class/gradienttexture"

	"github.com/shadowapex/godot-go/godot/class/curvetexture"
)

/*------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "class.go.tmpl" so they can be included in the generated
//   code.
//----------------------------------------------------------------------------*/

func NewParticlesMaterialFromPointer(ptr gdnative.Pointer) *ParticlesMaterial {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := ParticlesMaterial{}
	obj.SetOwner(owner)

	return &obj

}

/*
ParticlesMaterial defines particle properties and behavior. It is used in the [code]process_material[/code] of [Particles] and [Particles2D] emitter nodes. Some of this material's properties are applied to each particle when emitted, while others can have a [CurveTexture] applied to vary values over the lifetime of the particle.
*/
type ParticlesMaterial struct {
	material.Material
}

func (o *ParticlesMaterial) BaseClass() string {
	return "ParticlesMaterial"
}

/*
        Undocumented
	Args: [], Returns: Color
*/

func (o *ParticlesMaterial) GetColor() gdnative.Color {
	log.Println("Calling ParticlesMaterial.GetColor()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ParticlesMaterial", "get_color")

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
	Args: [], Returns: Texture
*/

func (o *ParticlesMaterial) GetColorRamp() texture.Texture {
	log.Println("Calling ParticlesMaterial.GetColorRamp()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ParticlesMaterial", "get_color_ramp")

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
	Args: [], Returns: Vector3
*/

func (o *ParticlesMaterial) GetEmissionBoxExtents() gdnative.Vector3 {
	log.Println("Calling ParticlesMaterial.GetEmissionBoxExtents()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ParticlesMaterial", "get_emission_box_extents")

	// Call the parent method.
	// Vector3
	retPtr := gdnative.NewEmptyVector3()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVector3FromPointer(retPtr)

	log.Println("  Got return value: ", ret)
	return ret
}

/*
        Undocumented
	Args: [], Returns: Texture
*/

func (o *ParticlesMaterial) GetEmissionColorTexture() texture.Texture {
	log.Println("Calling ParticlesMaterial.GetEmissionColorTexture()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ParticlesMaterial", "get_emission_color_texture")

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

func (o *ParticlesMaterial) GetEmissionNormalTexture() texture.Texture {
	log.Println("Calling ParticlesMaterial.GetEmissionNormalTexture()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ParticlesMaterial", "get_emission_normal_texture")

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
	Args: [], Returns: int
*/

func (o *ParticlesMaterial) GetEmissionPointCount() gdnative.Int {
	log.Println("Calling ParticlesMaterial.GetEmissionPointCount()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ParticlesMaterial", "get_emission_point_count")

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
	Args: [], Returns: Texture
*/

func (o *ParticlesMaterial) GetEmissionPointTexture() texture.Texture {
	log.Println("Calling ParticlesMaterial.GetEmissionPointTexture()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ParticlesMaterial", "get_emission_point_texture")

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
	Args: [], Returns: enum.ParticlesMaterial::EmissionShape
*/

/*
        Undocumented
	Args: [], Returns: float
*/

func (o *ParticlesMaterial) GetEmissionSphereRadius() gdnative.Float {
	log.Println("Calling ParticlesMaterial.GetEmissionSphereRadius()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ParticlesMaterial", "get_emission_sphere_radius")

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
	Args: [{ false flag int}], Returns: bool
*/

func (o *ParticlesMaterial) GetFlag(flag gdnative.Int) gdnative.Bool {
	log.Println("Calling ParticlesMaterial.GetFlag()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(flag)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ParticlesMaterial", "get_flag")

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
	Args: [], Returns: float
*/

func (o *ParticlesMaterial) GetFlatness() gdnative.Float {
	log.Println("Calling ParticlesMaterial.GetFlatness()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ParticlesMaterial", "get_flatness")

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
	Args: [], Returns: Vector3
*/

func (o *ParticlesMaterial) GetGravity() gdnative.Vector3 {
	log.Println("Calling ParticlesMaterial.GetGravity()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ParticlesMaterial", "get_gravity")

	// Call the parent method.
	// Vector3
	retPtr := gdnative.NewEmptyVector3()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVector3FromPointer(retPtr)

	log.Println("  Got return value: ", ret)
	return ret
}

/*
        Undocumented
	Args: [{ false param int}], Returns: float
*/

func (o *ParticlesMaterial) GetParam(param gdnative.Int) gdnative.Float {
	log.Println("Calling ParticlesMaterial.GetParam()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(param)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ParticlesMaterial", "get_param")

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
	Args: [{ false param int}], Returns: float
*/

func (o *ParticlesMaterial) GetParamRandomness(param gdnative.Int) gdnative.Float {
	log.Println("Calling ParticlesMaterial.GetParamRandomness()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(param)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ParticlesMaterial", "get_param_randomness")

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
	Args: [{ false param int}], Returns: Texture
*/

func (o *ParticlesMaterial) GetParamTexture(param gdnative.Int) texture.Texture {
	log.Println("Calling ParticlesMaterial.GetParamTexture()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(param)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ParticlesMaterial", "get_param_texture")

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
	Args: [], Returns: float
*/

func (o *ParticlesMaterial) GetSpread() gdnative.Float {
	log.Println("Calling ParticlesMaterial.GetSpread()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ParticlesMaterial", "get_spread")

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
	Args: [], Returns: GradientTexture
*/

func (o *ParticlesMaterial) GetTrailColorModifier() gradienttexture.GradientTexture {
	log.Println("Calling ParticlesMaterial.GetTrailColorModifier()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ParticlesMaterial", "get_trail_color_modifier")

	// Call the parent method.
	// GradientTexture
	retPtr := gradienttexture.NewEmptyGradientTexture()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gradienttexture.NewGradientTextureFromPointer(retPtr)

	log.Println("  Got return value: ", ret)
	return ret
}

/*
        Undocumented
	Args: [], Returns: int
*/

func (o *ParticlesMaterial) GetTrailDivisor() gdnative.Int {
	log.Println("Calling ParticlesMaterial.GetTrailDivisor()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ParticlesMaterial", "get_trail_divisor")

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
	Args: [], Returns: CurveTexture
*/

func (o *ParticlesMaterial) GetTrailSizeModifier() curvetexture.CurveTexture {
	log.Println("Calling ParticlesMaterial.GetTrailSizeModifier()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ParticlesMaterial", "get_trail_size_modifier")

	// Call the parent method.
	// CurveTexture
	retPtr := curvetexture.NewEmptyCurveTexture()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := curvetexture.NewCurveTextureFromPointer(retPtr)

	log.Println("  Got return value: ", ret)
	return ret
}

/*
        Undocumented
	Args: [{ false color Color}], Returns: void
*/

func (o *ParticlesMaterial) SetColor(color gdnative.Color) {
	log.Println("Calling ParticlesMaterial.SetColor()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromColor(color)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ParticlesMaterial", "set_color")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false ramp Texture}], Returns: void
*/

func (o *ParticlesMaterial) SetColorRamp(ramp texture.Texture) {
	log.Println("Calling ParticlesMaterial.SetColorRamp()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(ramp.GetOwner())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ParticlesMaterial", "set_color_ramp")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false extents Vector3}], Returns: void
*/

func (o *ParticlesMaterial) SetEmissionBoxExtents(extents gdnative.Vector3) {
	log.Println("Calling ParticlesMaterial.SetEmissionBoxExtents()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVector3(extents)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ParticlesMaterial", "set_emission_box_extents")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false texture Texture}], Returns: void
*/

func (o *ParticlesMaterial) SetEmissionColorTexture(texture texture.Texture) {
	log.Println("Calling ParticlesMaterial.SetEmissionColorTexture()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(texture.GetOwner())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ParticlesMaterial", "set_emission_color_texture")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false texture Texture}], Returns: void
*/

func (o *ParticlesMaterial) SetEmissionNormalTexture(texture texture.Texture) {
	log.Println("Calling ParticlesMaterial.SetEmissionNormalTexture()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(texture.GetOwner())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ParticlesMaterial", "set_emission_normal_texture")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false point_count int}], Returns: void
*/

func (o *ParticlesMaterial) SetEmissionPointCount(pointCount gdnative.Int) {
	log.Println("Calling ParticlesMaterial.SetEmissionPointCount()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(pointCount)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ParticlesMaterial", "set_emission_point_count")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false texture Texture}], Returns: void
*/

func (o *ParticlesMaterial) SetEmissionPointTexture(texture texture.Texture) {
	log.Println("Calling ParticlesMaterial.SetEmissionPointTexture()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(texture.GetOwner())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ParticlesMaterial", "set_emission_point_texture")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false shape int}], Returns: void
*/

func (o *ParticlesMaterial) SetEmissionShape(shape gdnative.Int) {
	log.Println("Calling ParticlesMaterial.SetEmissionShape()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(shape)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ParticlesMaterial", "set_emission_shape")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false radius float}], Returns: void
*/

func (o *ParticlesMaterial) SetEmissionSphereRadius(radius gdnative.Float) {
	log.Println("Calling ParticlesMaterial.SetEmissionSphereRadius()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromFloat(radius)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ParticlesMaterial", "set_emission_sphere_radius")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false flag int} { false enable bool}], Returns: void
*/

func (o *ParticlesMaterial) SetFlag(flag gdnative.Int, enable gdnative.Bool) {
	log.Println("Calling ParticlesMaterial.SetFlag()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(flag)
	ptrArguments[1] = gdnative.NewPointerFromBool(enable)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ParticlesMaterial", "set_flag")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false amount float}], Returns: void
*/

func (o *ParticlesMaterial) SetFlatness(amount gdnative.Float) {
	log.Println("Calling ParticlesMaterial.SetFlatness()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromFloat(amount)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ParticlesMaterial", "set_flatness")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false accel_vec Vector3}], Returns: void
*/

func (o *ParticlesMaterial) SetGravity(accelVec gdnative.Vector3) {
	log.Println("Calling ParticlesMaterial.SetGravity()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVector3(accelVec)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ParticlesMaterial", "set_gravity")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false param int} { false value float}], Returns: void
*/

func (o *ParticlesMaterial) SetParam(param gdnative.Int, value gdnative.Float) {
	log.Println("Calling ParticlesMaterial.SetParam()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(param)
	ptrArguments[1] = gdnative.NewPointerFromFloat(value)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ParticlesMaterial", "set_param")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false param int} { false randomness float}], Returns: void
*/

func (o *ParticlesMaterial) SetParamRandomness(param gdnative.Int, randomness gdnative.Float) {
	log.Println("Calling ParticlesMaterial.SetParamRandomness()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(param)
	ptrArguments[1] = gdnative.NewPointerFromFloat(randomness)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ParticlesMaterial", "set_param_randomness")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false param int} { false texture Texture}], Returns: void
*/

func (o *ParticlesMaterial) SetParamTexture(param gdnative.Int, texture texture.Texture) {
	log.Println("Calling ParticlesMaterial.SetParamTexture()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(param)
	ptrArguments[1] = gdnative.NewPointerFromObject(texture.GetOwner())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ParticlesMaterial", "set_param_texture")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false degrees float}], Returns: void
*/

func (o *ParticlesMaterial) SetSpread(degrees gdnative.Float) {
	log.Println("Calling ParticlesMaterial.SetSpread()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromFloat(degrees)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ParticlesMaterial", "set_spread")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false texture GradientTexture}], Returns: void
*/

func (o *ParticlesMaterial) SetTrailColorModifier(texture gradienttexture.GradientTexture) {
	log.Println("Calling ParticlesMaterial.SetTrailColorModifier()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(texture.GetOwner())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ParticlesMaterial", "set_trail_color_modifier")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false divisor int}], Returns: void
*/

func (o *ParticlesMaterial) SetTrailDivisor(divisor gdnative.Int) {
	log.Println("Calling ParticlesMaterial.SetTrailDivisor()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(divisor)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ParticlesMaterial", "set_trail_divisor")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false texture CurveTexture}], Returns: void
*/

func (o *ParticlesMaterial) SetTrailSizeModifier(texture curvetexture.CurveTexture) {
	log.Println("Calling ParticlesMaterial.SetTrailSizeModifier()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(texture.GetOwner())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ParticlesMaterial", "set_trail_size_modifier")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}
