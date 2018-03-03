package animatedsprite3d

import (
	"log"

	"github.com/shadowapex/godot-go/gdnative"

	"github.com/shadowapex/godot-go/godot/class/spritebase3d"
	"github.com/shadowapex/godot-go/godot/class/spriteframes"
)

/*------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "class.go.tmpl" so they can be included in the generated
//   code.
//----------------------------------------------------------------------------*/

func NewAnimatedSprite3DFromPointer(ptr gdnative.Pointer) *AnimatedSprite3D {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := AnimatedSprite3D{}
	obj.SetOwner(owner)

	return &obj

}

/*
Animations are created using a [SpriteFrames] resource, which can be configured in the editor via the SpriteFrames panel.
*/
type AnimatedSprite3D struct {
	spritebase3d.SpriteBase3D
}

func (o *AnimatedSprite3D) BaseClass() string {
	return "AnimatedSprite3D"
}

/*
        Undocumented
	Args: [], Returns: bool
*/

func (o *AnimatedSprite3D) X_IsPlaying() gdnative.Bool {
	log.Println("Calling AnimatedSprite3D.X_IsPlaying()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimatedSprite3D", "_is_playing")

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
	Args: [], Returns: void
*/

func (o *AnimatedSprite3D) X_ResChanged() {
	log.Println("Calling AnimatedSprite3D.X_ResChanged()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimatedSprite3D", "_res_changed")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false playing bool}], Returns: void
*/

func (o *AnimatedSprite3D) X_SetPlaying(playing gdnative.Bool) {
	log.Println("Calling AnimatedSprite3D.X_SetPlaying()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(playing)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimatedSprite3D", "_set_playing")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: String
*/

func (o *AnimatedSprite3D) GetAnimation() gdnative.String {
	log.Println("Calling AnimatedSprite3D.GetAnimation()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimatedSprite3D", "get_animation")

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

func (o *AnimatedSprite3D) GetFrame() gdnative.Int {
	log.Println("Calling AnimatedSprite3D.GetFrame()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimatedSprite3D", "get_frame")

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
	Args: [], Returns: SpriteFrames
*/

func (o *AnimatedSprite3D) GetSpriteFrames() spriteframes.SpriteFrames {
	log.Println("Calling AnimatedSprite3D.GetSpriteFrames()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimatedSprite3D", "get_sprite_frames")

	// Call the parent method.
	// SpriteFrames
	retPtr := spriteframes.NewEmptySpriteFrames()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := spriteframes.NewSpriteFramesFromPointer(retPtr)

	log.Println("  Got return value: ", ret)
	return ret
}

/*
        Return true if an animation if currently being played.
	Args: [], Returns: bool
*/

func (o *AnimatedSprite3D) IsPlaying() gdnative.Bool {
	log.Println("Calling AnimatedSprite3D.IsPlaying()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimatedSprite3D", "is_playing")

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
        Play the animation set in parameter. If no parameter is provided, the current animation is played.
	Args: [{ true anim String}], Returns: void
*/

func (o *AnimatedSprite3D) Play(anim gdnative.String) {
	log.Println("Calling AnimatedSprite3D.Play()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(anim)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimatedSprite3D", "play")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false animation String}], Returns: void
*/

func (o *AnimatedSprite3D) SetAnimation(animation gdnative.String) {
	log.Println("Calling AnimatedSprite3D.SetAnimation()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(animation)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimatedSprite3D", "set_animation")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false frame int}], Returns: void
*/

func (o *AnimatedSprite3D) SetFrame(frame gdnative.Int) {
	log.Println("Calling AnimatedSprite3D.SetFrame()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(frame)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimatedSprite3D", "set_frame")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false sprite_frames SpriteFrames}], Returns: void
*/

func (o *AnimatedSprite3D) SetSpriteFrames(spriteFrames spriteframes.SpriteFrames) {
	log.Println("Calling AnimatedSprite3D.SetSpriteFrames()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(spriteFrames.GetOwner())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimatedSprite3D", "set_sprite_frames")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*
        Stop the current animation (does not reset the frame counter).
	Args: [], Returns: void
*/

func (o *AnimatedSprite3D) Stop() {
	log.Println("Calling AnimatedSprite3D.Stop()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimatedSprite3D", "stop")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}
