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

func NewVisualScriptConstructorFromPointer(ptr gdnative.Pointer) *VisualScriptConstructor {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := VisualScriptConstructor{}
	obj.SetOwner(owner)

	return &obj

}

/*
Undocumented
*/
type VisualScriptConstructor struct {
	VisualScriptNode
}

func (o *VisualScriptConstructor) BaseClass() string {
	return "VisualScriptConstructor"
}

/*
        Undocumented
	Args: [], Returns: Dictionary
*/

func (o *VisualScriptConstructor) GetConstructor() gdnative.Dictionary {
	log.Println("Calling VisualScriptConstructor.GetConstructor()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptConstructor", "get_constructor")

	// Call the parent method.
	// Dictionary
	retPtr := gdnative.NewEmptyDictionary()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewDictionaryFromPointer(retPtr)

	log.Println("  Got return value: ", ret)
	return ret
}

/*
        Undocumented
	Args: [], Returns: enum.Variant::Type
*/

/*
        Undocumented
	Args: [{ false constructor Dictionary}], Returns: void
*/

func (o *VisualScriptConstructor) SetConstructor(constructor gdnative.Dictionary) {
	log.Println("Calling VisualScriptConstructor.SetConstructor()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromDictionary(constructor)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptConstructor", "set_constructor")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false type int}], Returns: void
*/

func (o *VisualScriptConstructor) SetConstructorType(aType gdnative.Int) {
	log.Println("Calling VisualScriptConstructor.SetConstructorType()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(aType)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptConstructor", "set_constructor_type")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetOwner(), ptrArguments, retPtr)

}
