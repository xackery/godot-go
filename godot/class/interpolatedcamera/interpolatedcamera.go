package interpolatedcamera

import (
	"log"

	"github.com/shadowapex/godot-go/gdnative"

	"github.com/shadowapex/godot-go/godot/class/camera"
	"github.com/shadowapex/godot-go/godot/class/object"
)

/*------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "class.go.tmpl" so they can be included in the generated
//   code.
//----------------------------------------------------------------------------*/

func NewInterpolatedCameraFromPointer(ptr gdnative.Pointer) *InterpolatedCamera {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := InterpolatedCamera{}
	obj.SetOwner(owner)

	return &obj

}

/*

 */
type InterpolatedCamera struct {
	camera.Camera
}

func (o *InterpolatedCamera) BaseClass() string {
	return "InterpolatedCamera"
}

/*
        Undocumented
	Args: [], Returns: float
*/

func (o *InterpolatedCamera) GetSpeed() gdnative.Float {
	log.Println("Calling InterpolatedCamera.GetSpeed()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("InterpolatedCamera", "get_speed")

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
	Args: [], Returns: NodePath
*/

func (o *InterpolatedCamera) GetTargetPath() gdnative.NodePath {
	log.Println("Calling InterpolatedCamera.GetTargetPath()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("InterpolatedCamera", "get_target_path")

	// Call the parent method.
	// NodePath
	retPtr := gdnative.NewEmptyNodePath()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewNodePathFromPointer(retPtr)

	log.Println("  Got return value: ", ret)
	return ret
}

/*
        Undocumented
	Args: [], Returns: bool
*/

func (o *InterpolatedCamera) IsInterpolationEnabled() gdnative.Bool {
	log.Println("Calling InterpolatedCamera.IsInterpolationEnabled()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("InterpolatedCamera", "is_interpolation_enabled")

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
	Args: [{ false target_path bool}], Returns: void
*/

func (o *InterpolatedCamera) SetInterpolationEnabled(targetPath gdnative.Bool) {
	log.Println("Calling InterpolatedCamera.SetInterpolationEnabled()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(targetPath)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("InterpolatedCamera", "set_interpolation_enabled")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false speed float}], Returns: void
*/

func (o *InterpolatedCamera) SetSpeed(speed gdnative.Float) {
	log.Println("Calling InterpolatedCamera.SetSpeed()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromFloat(speed)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("InterpolatedCamera", "set_speed")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*

	Args: [{ false target Object}], Returns: void
*/

func (o *InterpolatedCamera) SetTarget(target object.Object) {
	log.Println("Calling InterpolatedCamera.SetTarget()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(target.GetOwner())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("InterpolatedCamera", "set_target")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false target_path NodePath}], Returns: void
*/

func (o *InterpolatedCamera) SetTargetPath(targetPath gdnative.NodePath) {
	log.Println("Calling InterpolatedCamera.SetTargetPath()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromNodePath(targetPath)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("InterpolatedCamera", "set_target_path")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}
