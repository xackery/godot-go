package gradienttexture

import (
	"log"

	"github.com/shadowapex/godot-go/gdnative"

	"github.com/shadowapex/godot-go/godot/class/gradient"
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

func NewGradientTextureFromPointer(ptr gdnative.Pointer) *GradientTexture {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := GradientTexture{}
	obj.SetOwner(owner)

	return &obj

}

/*
Uses a [Gradient] to fill the texture data, the gradient will be filled from left to right using colors obtained from the gradient, this means that the texture does not necessarily represent an exact copy of the gradient, but instead an interpolation of samples obtained from the gradient at fixed steps (see [method set_width]).
*/
type GradientTexture struct {
	texture.Texture
}

func (o *GradientTexture) BaseClass() string {
	return "GradientTexture"
}

/*
        Undocumented
	Args: [], Returns: void
*/

func (o *GradientTexture) X_Update() {
	log.Println("Calling GradientTexture.X_Update()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GradientTexture", "_update")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: Gradient
*/

func (o *GradientTexture) GetGradient() gradient.Gradient {
	log.Println("Calling GradientTexture.GetGradient()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GradientTexture", "get_gradient")

	// Call the parent method.
	// Gradient
	retPtr := gradient.NewEmptyGradient()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gradient.NewGradientFromPointer(retPtr)

	log.Println("  Got return value: ", ret)
	return ret
}

/*
        Undocumented
	Args: [{ false gradient Gradient}], Returns: void
*/

func (o *GradientTexture) SetGradient(gradient gradient.Gradient) {
	log.Println("Calling GradientTexture.SetGradient()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(gradient.GetOwner())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GradientTexture", "set_gradient")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false width int}], Returns: void
*/

func (o *GradientTexture) SetWidth(width gdnative.Int) {
	log.Println("Calling GradientTexture.SetWidth()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(width)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GradientTexture", "set_width")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}
