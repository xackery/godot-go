package visualscript

import (
	"log"

	"github.com/shadowapex/godot-go/gdnative"
)

/*------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "class.go.tmpl" so they can be included in the generated
//   code.
//----------------------------------------------------------------------------*/

func NewVisualScriptLocalVarFromPointer(ptr gdnative.Pointer) *VisualScriptLocalVar {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := VisualScriptLocalVar{}
	obj.SetOwner(owner)

	return &obj

}

/*
Undocumented
*/
type VisualScriptLocalVar struct {
	VisualScriptNode
}

func (o *VisualScriptLocalVar) BaseClass() string {
	return "VisualScriptLocalVar"
}

/*
        Undocumented
	Args: [], Returns: String
*/

func (o *VisualScriptLocalVar) GetVarName() gdnative.String {
	log.Println("Calling VisualScriptLocalVar.GetVarName()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptLocalVar", "get_var_name")

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
	Args: [], Returns: enum.Variant::Type
*/

/*
        Undocumented
	Args: [{ false name String}], Returns: void
*/

func (o *VisualScriptLocalVar) SetVarName(name gdnative.String) {
	log.Println("Calling VisualScriptLocalVar.SetVarName()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(name)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptLocalVar", "set_var_name")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false type int}], Returns: void
*/

func (o *VisualScriptLocalVar) SetVarType(aType gdnative.Int) {
	log.Println("Calling VisualScriptLocalVar.SetVarType()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(aType)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptLocalVar", "set_var_type")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}
