package boneattachment

import (
	"log"

	"github.com/shadowapex/godot-go/gdnative"
	"github.com/shadowapex/godot-go/godot/class/spatial"
)

/*------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "class.go.tmpl" so they can be included in the generated
//   code.
//----------------------------------------------------------------------------*/

func NewBoneAttachmentFromPointer(ptr gdnative.Pointer) *BoneAttachment {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := BoneAttachment{}
	obj.SetOwner(owner)

	return &obj

}

/*
This node must be the child of a [Skeleton] node. You can then select a bone for this node to attach to. The BoneAttachment node will copy the transform of the selected bone.
*/
type BoneAttachment struct {
	spatial.Spatial
}

func (o *BoneAttachment) BaseClass() string {
	return "BoneAttachment"
}

/*
        Undocumented
	Args: [], Returns: String
*/

func (o *BoneAttachment) GetBoneName() gdnative.String {
	log.Println("Calling BoneAttachment.GetBoneName()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("BoneAttachment", "get_bone_name")

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
	Args: [{ false bone_name String}], Returns: void
*/

func (o *BoneAttachment) SetBoneName(boneName gdnative.String) {
	log.Println("Calling BoneAttachment.SetBoneName()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(boneName)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("BoneAttachment", "set_bone_name")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}
