package physics

import (
	"log"

	"github.com/shadowapex/godot-go/gdnative"

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

func NewPhysics2DDirectBodyStateFromPointer(ptr gdnative.Pointer) *Physics2DDirectBodyState {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := Physics2DDirectBodyState{}
	obj.SetOwner(owner)

	return &obj

}

/*
Direct access object to a physics body in the [Physics2DServer]. This object is passed via the direct state callback of rigid/character bodies, and is intended for changing the direct state of that body.
*/
type Physics2DDirectBodyState struct {
	object.Object
}

func (o *Physics2DDirectBodyState) BaseClass() string {
	return "Physics2DDirectBodyState"
}

/*
        Undocumented
	Args: [], Returns: float
*/

func (o *Physics2DDirectBodyState) GetAngularVelocity() gdnative.Float {
	log.Println("Calling Physics2DDirectBodyState.GetAngularVelocity()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Physics2DDirectBodyState", "get_angular_velocity")

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
        Return the [RID] of the collider.
	Args: [{ false contact_idx int}], Returns: RID
*/

func (o *Physics2DDirectBodyState) GetContactCollider(contactIdx gdnative.Int) gdnative.RID {
	log.Println("Calling Physics2DDirectBodyState.GetContactCollider()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(contactIdx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Physics2DDirectBodyState", "get_contact_collider")

	// Call the parent method.
	// RID
	retPtr := gdnative.NewEmptyRid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewRidFromPointer(retPtr)

	log.Println("  Got return value: ", ret)
	return ret
}

/*
        Return the object id of the collider.
	Args: [{ false contact_idx int}], Returns: int
*/

func (o *Physics2DDirectBodyState) GetContactColliderId(contactIdx gdnative.Int) gdnative.Int {
	log.Println("Calling Physics2DDirectBodyState.GetContactColliderId()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(contactIdx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Physics2DDirectBodyState", "get_contact_collider_id")

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
        Return the collider object, this depends on how it was created (will return a scene node if such was used to create it).
	Args: [{ false contact_idx int}], Returns: Object
*/

func (o *Physics2DDirectBodyState) GetContactColliderObject(contactIdx gdnative.Int) object.Object {
	log.Println("Calling Physics2DDirectBodyState.GetContactColliderObject()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(contactIdx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Physics2DDirectBodyState", "get_contact_collider_object")

	// Call the parent method.
	// Object
	retPtr := object.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := object.NewObjectFromPointer(retPtr)

	log.Println("  Got return value: ", ret)
	return ret
}

/*
        Return the contact position in the collider.
	Args: [{ false contact_idx int}], Returns: Vector2
*/

func (o *Physics2DDirectBodyState) GetContactColliderPosition(contactIdx gdnative.Int) gdnative.Vector2 {
	log.Println("Calling Physics2DDirectBodyState.GetContactColliderPosition()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(contactIdx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Physics2DDirectBodyState", "get_contact_collider_position")

	// Call the parent method.
	// Vector2
	retPtr := gdnative.NewEmptyVector2()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVector2FromPointer(retPtr)

	log.Println("  Got return value: ", ret)
	return ret
}

/*
        Return the collider shape index.
	Args: [{ false contact_idx int}], Returns: int
*/

func (o *Physics2DDirectBodyState) GetContactColliderShape(contactIdx gdnative.Int) gdnative.Int {
	log.Println("Calling Physics2DDirectBodyState.GetContactColliderShape()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(contactIdx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Physics2DDirectBodyState", "get_contact_collider_shape")

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
        Return the metadata of the collided shape. This metadata is different from [method Object.get_meta], and is set with [method Physics2DServer.shape_set_data].
	Args: [{ false contact_idx int}], Returns: Variant
*/

func (o *Physics2DDirectBodyState) GetContactColliderShapeMetadata(contactIdx gdnative.Int) gdnative.Variant {
	log.Println("Calling Physics2DDirectBodyState.GetContactColliderShapeMetadata()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(contactIdx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Physics2DDirectBodyState", "get_contact_collider_shape_metadata")

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
        Return the linear velocity vector at contact point of the collider.
	Args: [{ false contact_idx int}], Returns: Vector2
*/

func (o *Physics2DDirectBodyState) GetContactColliderVelocityAtPosition(contactIdx gdnative.Int) gdnative.Vector2 {
	log.Println("Calling Physics2DDirectBodyState.GetContactColliderVelocityAtPosition()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(contactIdx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Physics2DDirectBodyState", "get_contact_collider_velocity_at_position")

	// Call the parent method.
	// Vector2
	retPtr := gdnative.NewEmptyVector2()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVector2FromPointer(retPtr)

	log.Println("  Got return value: ", ret)
	return ret
}

/*
        Return the amount of contacts this body has with other bodies. Note that by default this returns 0 unless bodies are configured to log contacts.
	Args: [], Returns: int
*/

func (o *Physics2DDirectBodyState) GetContactCount() gdnative.Int {
	log.Println("Calling Physics2DDirectBodyState.GetContactCount()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Physics2DDirectBodyState", "get_contact_count")

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
        Return the local normal (of this body) of the contact point.
	Args: [{ false contact_idx int}], Returns: Vector2
*/

func (o *Physics2DDirectBodyState) GetContactLocalNormal(contactIdx gdnative.Int) gdnative.Vector2 {
	log.Println("Calling Physics2DDirectBodyState.GetContactLocalNormal()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(contactIdx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Physics2DDirectBodyState", "get_contact_local_normal")

	// Call the parent method.
	// Vector2
	retPtr := gdnative.NewEmptyVector2()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVector2FromPointer(retPtr)

	log.Println("  Got return value: ", ret)
	return ret
}

/*
        Return the local position (of this body) of the contact point.
	Args: [{ false contact_idx int}], Returns: Vector2
*/

func (o *Physics2DDirectBodyState) GetContactLocalPosition(contactIdx gdnative.Int) gdnative.Vector2 {
	log.Println("Calling Physics2DDirectBodyState.GetContactLocalPosition()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(contactIdx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Physics2DDirectBodyState", "get_contact_local_position")

	// Call the parent method.
	// Vector2
	retPtr := gdnative.NewEmptyVector2()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVector2FromPointer(retPtr)

	log.Println("  Got return value: ", ret)
	return ret
}

/*
        Return the local shape index of the collision.
	Args: [{ false contact_idx int}], Returns: int
*/

func (o *Physics2DDirectBodyState) GetContactLocalShape(contactIdx gdnative.Int) gdnative.Int {
	log.Println("Calling Physics2DDirectBodyState.GetContactLocalShape()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(contactIdx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Physics2DDirectBodyState", "get_contact_local_shape")

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
	Args: [], Returns: float
*/

func (o *Physics2DDirectBodyState) GetInverseInertia() gdnative.Float {
	log.Println("Calling Physics2DDirectBodyState.GetInverseInertia()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Physics2DDirectBodyState", "get_inverse_inertia")

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

func (o *Physics2DDirectBodyState) GetInverseMass() gdnative.Float {
	log.Println("Calling Physics2DDirectBodyState.GetInverseMass()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Physics2DDirectBodyState", "get_inverse_mass")

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
	Args: [], Returns: Vector2
*/

func (o *Physics2DDirectBodyState) GetLinearVelocity() gdnative.Vector2 {
	log.Println("Calling Physics2DDirectBodyState.GetLinearVelocity()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Physics2DDirectBodyState", "get_linear_velocity")

	// Call the parent method.
	// Vector2
	retPtr := gdnative.NewEmptyVector2()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVector2FromPointer(retPtr)

	log.Println("  Got return value: ", ret)
	return ret
}

/*
        Return the current state of space, useful for queries.
	Args: [], Returns: Physics2DDirectSpaceState
*/

func (o *Physics2DDirectBodyState) GetSpaceState() Physics2DDirectSpaceState {
	log.Println("Calling Physics2DDirectBodyState.GetSpaceState()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Physics2DDirectBodyState", "get_space_state")

	// Call the parent method.
	// Physics2DDirectSpaceState
	retPtr := NewEmptyPhysics2DDirectSpaceState()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := NewPhysics2DDirectSpaceStateFromPointer(retPtr)

	log.Println("  Got return value: ", ret)
	return ret
}

/*
        Undocumented
	Args: [], Returns: float
*/

func (o *Physics2DDirectBodyState) GetStep() gdnative.Float {
	log.Println("Calling Physics2DDirectBodyState.GetStep()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Physics2DDirectBodyState", "get_step")

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

func (o *Physics2DDirectBodyState) GetTotalAngularDamp() gdnative.Float {
	log.Println("Calling Physics2DDirectBodyState.GetTotalAngularDamp()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Physics2DDirectBodyState", "get_total_angular_damp")

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
	Args: [], Returns: Vector2
*/

func (o *Physics2DDirectBodyState) GetTotalGravity() gdnative.Vector2 {
	log.Println("Calling Physics2DDirectBodyState.GetTotalGravity()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Physics2DDirectBodyState", "get_total_gravity")

	// Call the parent method.
	// Vector2
	retPtr := gdnative.NewEmptyVector2()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVector2FromPointer(retPtr)

	log.Println("  Got return value: ", ret)
	return ret
}

/*
        Undocumented
	Args: [], Returns: float
*/

func (o *Physics2DDirectBodyState) GetTotalLinearDamp() gdnative.Float {
	log.Println("Calling Physics2DDirectBodyState.GetTotalLinearDamp()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Physics2DDirectBodyState", "get_total_linear_damp")

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
	Args: [], Returns: Transform2D
*/

func (o *Physics2DDirectBodyState) GetTransform() gdnative.Transform2D {
	log.Println("Calling Physics2DDirectBodyState.GetTransform()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Physics2DDirectBodyState", "get_transform")

	// Call the parent method.
	// Transform2D
	retPtr := gdnative.NewEmptyTransform2D()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewTransform2DFromPointer(retPtr)

	log.Println("  Got return value: ", ret)
	return ret
}

/*
        Call the built-in force integration code.
	Args: [], Returns: void
*/

func (o *Physics2DDirectBodyState) IntegrateForces() {
	log.Println("Calling Physics2DDirectBodyState.IntegrateForces()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Physics2DDirectBodyState", "integrate_forces")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: bool
*/

func (o *Physics2DDirectBodyState) IsSleeping() gdnative.Bool {
	log.Println("Calling Physics2DDirectBodyState.IsSleeping()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Physics2DDirectBodyState", "is_sleeping")

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
	Args: [{ false velocity float}], Returns: void
*/

func (o *Physics2DDirectBodyState) SetAngularVelocity(velocity gdnative.Float) {
	log.Println("Calling Physics2DDirectBodyState.SetAngularVelocity()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromFloat(velocity)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Physics2DDirectBodyState", "set_angular_velocity")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false velocity Vector2}], Returns: void
*/

func (o *Physics2DDirectBodyState) SetLinearVelocity(velocity gdnative.Vector2) {
	log.Println("Calling Physics2DDirectBodyState.SetLinearVelocity()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVector2(velocity)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Physics2DDirectBodyState", "set_linear_velocity")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false enabled bool}], Returns: void
*/

func (o *Physics2DDirectBodyState) SetSleepState(enabled gdnative.Bool) {
	log.Println("Calling Physics2DDirectBodyState.SetSleepState()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(enabled)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Physics2DDirectBodyState", "set_sleep_state")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false transform Transform2D}], Returns: void
*/

func (o *Physics2DDirectBodyState) SetTransform(transform gdnative.Transform2D) {
	log.Println("Calling Physics2DDirectBodyState.SetTransform()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromTransform2D(transform)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Physics2DDirectBodyState", "set_transform")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}
