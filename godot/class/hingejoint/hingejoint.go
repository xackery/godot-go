package hingejoint

import (
	"log"

	"github.com/shadowapex/godot-go/gdnative"
	"github.com/shadowapex/godot-go/godot/class/joint"
)

/*------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "class.go.tmpl" so they can be included in the generated
//   code.
//----------------------------------------------------------------------------*/

func NewHingeJointFromPointer(ptr gdnative.Pointer) *HingeJoint {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := HingeJoint{}
	obj.SetOwner(owner)

	return &obj

}

/*
Normally uses the z-axis of body A as the hinge axis, another axis can be specified when adding it manually though.
*/
type HingeJoint struct {
	joint.Joint
}

func (o *HingeJoint) BaseClass() string {
	return "HingeJoint"
}

/*
        Undocumented
	Args: [], Returns: float
*/

func (o *HingeJoint) X_GetLowerLimit() gdnative.Float {
	log.Println("Calling HingeJoint.X_GetLowerLimit()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("HingeJoint", "_get_lower_limit")

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
	Args: [], Returns: float
*/

func (o *HingeJoint) X_GetUpperLimit() gdnative.Float {
	log.Println("Calling HingeJoint.X_GetUpperLimit()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("HingeJoint", "_get_upper_limit")

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
	Args: [{ false lower_limit float}], Returns: void
*/

func (o *HingeJoint) X_SetLowerLimit(lowerLimit gdnative.Float) {
	log.Println("Calling HingeJoint.X_SetLowerLimit()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromFloat(lowerLimit)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("HingeJoint", "_set_lower_limit")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false upper_limit float}], Returns: void
*/

func (o *HingeJoint) X_SetUpperLimit(upperLimit gdnative.Float) {
	log.Println("Calling HingeJoint.X_SetUpperLimit()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromFloat(upperLimit)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("HingeJoint", "_set_upper_limit")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false flag int}], Returns: bool
*/

func (o *HingeJoint) GetFlag(flag gdnative.Int) gdnative.Bool {
	log.Println("Calling HingeJoint.GetFlag()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(flag)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("HingeJoint", "get_flag")

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
	Args: [{ false param int}], Returns: float
*/

func (o *HingeJoint) GetParam(param gdnative.Int) gdnative.Float {
	log.Println("Calling HingeJoint.GetParam()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(param)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("HingeJoint", "get_param")

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
	Args: [{ false flag int} { false enabled bool}], Returns: void
*/

func (o *HingeJoint) SetFlag(flag gdnative.Int, enabled gdnative.Bool) {
	log.Println("Calling HingeJoint.SetFlag()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(flag)
	ptrArguments[1] = gdnative.NewPointerFromBool(enabled)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("HingeJoint", "set_flag")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false param int} { false value float}], Returns: void
*/

func (o *HingeJoint) SetParam(param gdnative.Int, value gdnative.Float) {
	log.Println("Calling HingeJoint.SetParam()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(param)
	ptrArguments[1] = gdnative.NewPointerFromFloat(value)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("HingeJoint", "set_param")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}
